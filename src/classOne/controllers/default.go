package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["testData"] = "区块链 get"
	c.TplName = "test.html"
}

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Post() {
	c.Data["testData"] = "区块链 post"
	c.TplName = "test.html"
}

type SimpleController struct {
	beego.Controller
}

func (c *SimpleController) PostFunc() {
	c.Data["postData"] = "Post function"
	c.TplName = "test.html"
}

func (c *SimpleController) GetFunc() {
	path := c.GetString(":id")
	beego.Info("path = ", path)
	c.Data["getData"] = "Get function"
	c.TplName = "test.html"
}