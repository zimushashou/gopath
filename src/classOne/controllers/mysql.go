package controllers

import (
	"classOne/models"
	"database/sql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)
import _ "github.com/go-sql-driver/mysql"

type MysqlController struct {
	beego.Controller
}

func (c* MysqlController)OperatMysql()  {
	// 1.打开数据库
	// 2.创建表
	// 3.关闭
	db, e := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/BeegoDB")
	if e!= nil {
		return
	}
	defer db.Close()

	//result, e := db.Exec("create table student(id int not null AUTO_INCREMENT, name varchar(20), age int, PRIMARY KEY(id))")
	//if e!=nil {
	//	beego.Info("err: ", e)
	//	return
	//}
	//beego.Info("create table student sucess ", result)
	//result, e := db.Exec("insert into student VALUES (?,?,?)", 1, "lina", 15)
	//if e!=nil {
	//	beego.Info("err: ", e)
	//	return
	//}
	//beego.Info("insert data sucess ", result)
	data, e := db.Query("select name from student")
	if e!=nil {
		beego.Info("err: ", e)
		return
	}

	var userName string
	for data.Next(){
		data.Scan(&userName)
		beego.Info(userName)
	}
	c.Ctx.WriteString("insert data sucess!")
}

type ORMController struct {
	beego.Controller
}

func (c *ORMController)ShowORM()  {

	newOrm := orm.NewOrm()
	var user models.User
	user.Name = "itcast"
	user.Passwd = "heima"
	id, e := newOrm.Insert(&user)
	if e==nil{
		c.Data["testData"] = id
	}
	beego.Info("插入成功")
	c.TplName = "test.html"
}