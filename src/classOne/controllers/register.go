package controllers

import (
	"classOne/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController)ShowRegister()  {

	c.TplName = "register.html"
}

func (c *RegisterController)HandleRegister()  {
	// 写入数据库时判断
	userName := c.GetString("userName")
	passWord := c.GetString("password")
	if userName == "" || passWord == "" {
		beego.Info("数据不完整，请继续输入")
		c.TplName = "register.html"
		return
	}
	// 写入数据库
	newOrm := orm.NewOrm()
	var user models.User
	user.Name = userName
	user.Passwd = passWord
	beego.Info(user.Name, user.Passwd)
	newOrm.Read(&user, "name")
	beego.Info(user.Name, user.Passwd)
	if _, err := newOrm.Insert(&user); err != nil{
		beego.Info("用户名已存在，请修改用户名后重新注册")
		c.TplName = "register.html"
		return
	}
	//c.Ctx.WriteString("注册成功")
	c.Redirect("/login", 302)
}