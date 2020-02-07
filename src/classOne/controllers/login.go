package controllers

import (
	"classOne/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController)ShowLogin()  {
	c.TplName = "login.html"
}
func (c *LoginController)HandleLogin()  {
	var user models.User
	newOrm := orm.NewOrm()
	name := c.GetString("userName")
	password := c.GetString("password")

	if err := newOrm.Read(&user, name); err != nil{
		c.TplName = "login.html"
		return
	}
	if password != user.Passwd {
		c.Ctx.WriteString("登录失败")
	}
	c.Redirect("/", 302)
}