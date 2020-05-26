package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	"ihome/ihomeweb/models"
	"ihome/ihomeweb/utils"
	"time"

	_ "github.com/astaxie/beego/cache"

	_ "github.com/astaxie/beego/cache/redis"

	_ "github.com/gomodule/redigo/redis"

	_ "github.com/garyburd/redigo/redis"

	example "ihome/GetArea/proto/example"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetArea(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("Get Area api/v1.0/areas")
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	// ready redis
	redis_config_map := map[string]string{
		"key":   utils.G_server_name,
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum": utils.G_redis_dbnum,
		"auth": "123456",
	}
	beego.Info(redis_config_map)
	redis_config, _ := json.Marshal(redis_config_map)
	bm, err := cache.NewCache("redis", string(redis_config))
	if err != nil {
		beego.Info("new cache failed ", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	// 1.获取缓存数据
	areas_info_value := bm.Get("areas_info")
	if areas_info_value != nil {
		// send info to webclient
		ares_info := []map[string]interface{}{}
		err = json.Unmarshal(areas_info_value.([]byte), &ares_info)
		for key, value := range ares_info {
			beego.Info(key, value)
			area := example.Response_Address{Aid :int32(value["aid"].(float64)),
				Aname :value["aname"].(string)}
			rsp.Data = append(rsp.Data, &area)
		}
		return nil
	}
	// 2.如果没有缓存从mysql里进行查询
	o := orm.NewOrm()
	qs := o.QueryTable("area")
	var areas []models.Area
	num, err := qs.All(&areas)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	if num == 0 {
		rsp.Errno = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	// 3.获取数据写入缓存
	beego.Info("input to redis")
	areas_info_list, _ := json.Marshal(areas)
	err = bm.Put("area_info", areas_info_list, time.Second*3600)
	if err != nil {
		beego.Info("put error to redis", err)
		rsp.Errno = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	// response area info
	for key, value := range areas {
		beego.Info(key, value)
		area := example.Response_Address{Aid: int32(value.Id), Aname: string(value.Name)}
		rsp.Data = append(rsp.Data, &area)
	}
	return nil
}
