package controllers

import (
	"classOne/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"path"
	"strconv"
	"time"
)

type IndexsController struct {
	beego.Controller
}

func (c *IndexsController)ShowIndex()  {
	c.TplName = "index.html"
}
func (c *IndexsController)HandleIndex()  {
	c.TplName = "index.html"
}

type AddArticalController struct {
	beego.Controller
}
func (c *AddArticalController)ShowArtical()  {
	c.TplName = "add.html"
}
func (c *AddArticalController)AddArtical()  {
	//1.获取数据
	title := c.GetString("articleName")
	//arcType := c.GetString("select")
	content := c.GetString("content")

	//2.判断合法性
	if title == "" || content == ""{
		c.Data["errmsg"] = "数据不完整，请补充后再提交"
		c.TplName = "add.html"
		return
	}
	beego.Info("title: ", title)
	file, header, err := c.GetFile("uploadname")
	if err!=nil{		// 必须在此处判断，否则如果没有上传图片访问file的操作报错
		c.Data["errmsg"] = "数据不完整，请补充后再提交"
		c.TplName = "add.html"
		return
	}
	defer file.Close()

	beego.Info("111: ", title)
	//判断文件类型
	ext := path.Ext(header.Filename)
	if ext != ".jpg" && ext != ".png"{
		beego.Info("上传的文件类型不正确")
		c.Data["errmsg"] = "数据不完整，请补充后再提交"
		c.TplName = "add.html"
		return
	}
	//判断文件大小
	if header.Size > 500000{
		beego.Info("上传的文件太大")
		c.Data["errmsg"] = "数据不完整，请补充后再提交"
		c.TplName = "add.html"
		return
	}
	//生成不重名文件
	fileName := time.Now().Format("2006-01-02 15:04:05")

	c.SaveToFile("uploadname", "./static/img/"+fileName+ext)
	//3.插入数据
	newOrm := orm.NewOrm()
	artical := models.Artical{}
	artical.Title = title
	artical.Content = content
	artical.Image = "./static/img/"+header.Filename
	//artical.Type = arcType
	_, err = newOrm.Insert(&artical)
	if err != nil{
		c.Data["errmsg"] = "添加数据失败"
		c.TplName = "add.html"
		beego.Info("err:", err)
		return
	}

	//4.返回视图
	c.Redirect("/ArticalList",302)
}
type ArticalListController struct {
	beego.Controller
}
func (c *ArticalListController)ShowArticalList()  {

	newOrm := orm.NewOrm()
	var articals[] models.Artical
	qs := newOrm.QueryTable("artical")
	count,err := qs.Count()
	if err != nil{
		beego.Info("获取数据数量失败")
		return
	}
	pageSize := 2
	pageIndexTmp := c.GetString("pageIndex")
	pageIndex,_ := strconv.Atoi(pageIndexTmp)
	if pageIndex == 0{
		pageIndex = 1
	}
	startIndex := pageSize*(pageIndex-1)
	qs.Limit(pageSize, startIndex).All(&articals)

	pageCount := math.Ceil(float64(count)/float64(pageSize))		// math.floor向下取整
	c.Data["pageCount"] = int(pageCount)
	c.Data["count"] = count
	//c.Data["pageIndex"] = pageIndex
	c.Data["articals"] = articals
	c.TplName = "index.html"
}
func (c *ArticalListController)HandleArticalList()  {
	c.Redirect("/ArticalList",302)
}

type AddArticalTypeController struct {
	beego.Controller
}
func (c *AddArticalTypeController)ShowArticalType()  {
	c.TplName = "addType.html"
}
func (c *AddArticalTypeController)AddArticalType()  {

	c.Redirect("/AddArticalType",302)
}

type ArticalDetailController struct {
	beego.Controller
}
func (c *ArticalDetailController)ShowArticalDetail()  {
	id, er := c.GetInt("articalId")
	//数据校验
	if er != nil{
		beego.Info("传递的链接错误")
	}
	beego.Info("id = ", id)
	//操作数据
	o := orm.NewOrm()
	var artical models.Artical
	artical.Id = id
	o.Read(&artical)
	artical.Count += 1		// 必须在读取后修改才有效
	o.Update(&artical)
	c.Data["artical"] = artical
	c.TplName = "content.html"
}
func (c *ArticalDetailController)HandleArticalDetail()  {
	c.Redirect("/ArticalList", 302)
}
type ArticalDeleteController struct {
	beego.Controller
}
func (c *ArticalDeleteController)ShowArticalDelete()  {
	id, er := c.GetInt("articalId")
	//数据校验
	if er != nil{
		beego.Info("传递的链接错误")
	}
	beego.Info("id = ", id)
	newOrm := orm.NewOrm()
	// 删除数据
	var artical models.Artical
	artical.Id = id
	newOrm.Delete(&artical)
	// 查询数据
	var articals[] models.Artical
	qs := newOrm.QueryTable("artical")
	qs.All(&articals)
	c.Data["articals"] = articals
	c.TplName = "index.html"
}
func (c *ArticalDeleteController)HandleArticalDelete()  {
	c.Redirect("/ArticalList", 302)
}

type ArticalEditController struct {
	beego.Controller
}

func (c *ArticalEditController)ShowArticalEdit()  {
	id, er := c.GetInt("articalId")
	//数据校验
	if er != nil{
		beego.Info("传递的链接错误")
	}
	beego.Info("id = ", id)
	// 查询编辑文章内容
	o := orm.NewOrm()
	var artical models.Artical
	artical.Id = id
	err := o.Read(&artical)
	if err != nil{
		c.Data["errmsg"] = "获取数据失败"
		c.TplName = "update.html"
		beego.Info("err:", err)
		return
	}
	c.Data["artical"] = artical
	c.TplName = "update.html"
}

func (c *ArticalEditController)HandleArticalEdit()  {
	id, er := c.GetInt("articalId")
	//数据校验
	if er != nil{
		beego.Info("传递的链接错误")
	}
	title := c.GetString("articleName")
	//arcType := c.GetString("select")
	content := c.GetString("content")

	//2.判断合法性
	if title == "" || content == ""{
		c.Data["errmsg"] = "数据不完整，请补充后再提交"
		c.TplName = "add.html"
		return
	}
	beego.Info("title: ", title)
	o := orm.NewOrm()
	var artical1 models.Artical
	artical1.Id = id
	err := o.Read(&artical1)
	if err != nil{
		c.Data["errmsg"] = "获取数据失败"
		c.TplName = "update.html"
		beego.Info("err:", err)
		return
	}

	file, header, err := c.GetFile("uploadname")
	if artical1.Image == "" {

		if err!=nil{		// 必须在此处判断，否则如果没有上传图片访问file的操作报错
			beego.Info(err)
			//c.Data["errmsg"] = "数据不完整，请补充后再提交"
			//c.TplName = "add.html"
			//return
		}
		defer file.Close()

		beego.Info("111: ", title)
		//判断文件类型
		ext := path.Ext(header.Filename)
		if ext != ".jpg" && ext != ".png"{
			beego.Info("上传的文件类型不正确")
			c.Data["errmsg"] = "数据不完整，请补充后再提交"
			c.TplName = "add.html"
			return
		}
		//判断文件大小
		if header.Size > 500000{
			beego.Info("上传的文件太大")
			c.Data["errmsg"] = "数据不完整，请补充后再提交"
			c.TplName = "add.html"
			return
		}
		//生成不重名文件
		fileName := time.Now().Format("2006-01-02 15:04:05")

		c.SaveToFile("uploadname", "./static/img/"+fileName+ext)
	}

	//操作数据
	var artical models.Artical
	artical.Id = id
	o.Read(&artical)
	artical.Title = title
	artical.Content = content
	if artical.Image == "" {
		artical.Image = "./static/img/"+header.Filename
	}

	o.Update(&artical)
	c.Redirect("/ArticalList", 302)
}