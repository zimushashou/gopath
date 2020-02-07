package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)
import _ "github.com/go-sql-driver/mysql"
type User struct {
	Id int
	Name string
	Passwd string
	Articals[] *Artical		`orm:"reverse(many)"`
}
type Artical struct {
	Id int			`orm:"auto;pk"`
	Title string	`orm:"size(20)"`
	Type string		`orm:"size(15)"`
	Content string	`orm:"size(500)"`
	Image string	`orm:"size(50)"`
	AddTime time.Time  `orm:"auto_now"`
	Count	int `orm:"default(0);null"`
	ArticalType *ArticalType 	`orm:"rel(fk)"`
	Users[] *User		`orm:"rel(m2m)"`
}
type ArticalType struct {
	Id int
	TypeName string			`orm:size(20)`
	Articals[] *Artical		`orm:"reverse(many)"`
}
func init()  {
	// 1.连接数据库
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/BeegoDB")
	// 2.注册表
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Artical), new(ArticalType))
	// 3.生成表
	orm.RunSyncdb("default", false, true)
}