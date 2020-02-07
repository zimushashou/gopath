package routers

import (
	"classOne/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	//beego.Router("/index", &controllers.IndexController{})
	beego.Router("/simple/:id([0-9]+)", &controllers.SimpleController{}, "Get:GetFunc;Post:PostFunc")

    beego.Router("/mysql", &controllers.MysqlController{}, "Get:OperatMysql")
    beego.Router("/orm", &controllers.ORMController{}, "Get:ShowORM")
	beego.Router("/register", &controllers.RegisterController{}, "Get:ShowRegister;Post:HandleRegister")
	beego.Router("/login", &controllers.LoginController{}, "Get:ShowLogin;Post:HandleLogin")
	beego.Router("/", &controllers.IndexsController{}, "Get:ShowIndex;Post:HandleIndex")
	beego.Router("/AddArtical", &controllers.AddArticalController{}, "Get:ShowArtical;Post:AddArtical")
	beego.Router("/ArticalList", &controllers.ArticalListController{}, "Get:ShowArticalList;Post:HandleArticalList")
	beego.Router("/AddArticalType", &controllers.AddArticalTypeController{}, "Get:ShowArticalType;Post:AddArticalType")
	beego.Router("/ArticalDetail", &controllers.ArticalDetailController{}, "Get:ShowArticalDetail;Post:HandleArticalDetail")
	beego.Router("/ArticalDelete", &controllers.ArticalDeleteController{}, "Get:ShowArticalDelete;Post:HandleArticalDelete")
	beego.Router("/ArticalEdit", &controllers.ArticalEditController{}, "Get:ShowArticalEdit;Post:HandleArticalEdit")
}
