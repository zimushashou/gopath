package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-grpc"
	GETAREA "ihome/GetArea/proto/example"

	_ "github.com/micro/go-micro/client"
	"ihome/ihomeweb/models"
	"ihome/ihomeweb/utils"
	"net/http"
)

//获取地区
func GetArea(w http.ResponseWriter, r *http.Request,_ httprouter.Params) {
	beego.Info("获取地区请求客户端 url:api/v1.0/areas")

	//创建新的grpc返回句柄
	server :=grpc.NewService()
	//服务出初始化
	server.Init()




	//创建获取地区的服务并且返回句柄
	exampleClient := GETAREA.NewExampleService("go.micro.srv.GetArea",
		server.Client())

	//调用函数并且获得返回数据
	rsp, err := exampleClient.GetArea(context.TODO(), &GETAREA.Request{})
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
		}
	//创建返回类型的切片
	area_list := []models.Area{}
	//循环读取服务返回的数据
	for _,value := range rsp.Data{

		tmp :=models.Area{Id:int(value.Aid),Name:value.Aname,Houses:nil}

		area_list = append(area_list,tmp)
		 }
	//创建返回数据map
	response := map[string]interface{}{
		"errno": rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data" : area_list,
	}
	//注意的点
	w.Header().Set("Content-Type", "application/json")

	 // 将返回数据map发送给前端
	 if err := json.NewEncoder(w).Encode(response); err != nil {

		http.Error(w, err.Error(), 503)
		return

	}

}


//获取session
func GetSession(w http.ResponseWriter, r *http.Request,_ httprouter.Params) {

	beego.Info("获取Session url:api/v1.0/session")

	//创建返回数据map
	response := map[string]interface{}{
		"errno": utils.RECODE_SESSIONERR,
		"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),

		}
	w.Header().Set("Content-Type", "application/json")

	// 将返回数据map发送给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
}

//获取首页轮播
func GetIndex(w http.ResponseWriter, r *http.Request,_ httprouter.Params) {

	beego.Info("获取首页轮播 url:api/v1.0/houses/index")

	//创建返回数据map
	response := map[string]interface{}{
		"errno": utils.RECODE_OK,
		"errmsg": utils.RecodeText(utils.RECODE_OK),

		}
	w.Header().Set("Content-Type", "application/json")

	// 将返回数据map发送给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		 http.Error(w, err.Error(), 503)
		 return

	}

}
