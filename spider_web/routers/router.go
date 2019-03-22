package routers

import (
	"github.com/astaxie/beego"
	"github.com/wliangde/goleond/spider_web/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//home
	beego.Router("/", &controllers.HomeController{}, "*:Index")
	beego.Router("/home", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")
	beego.Router("/home/chooseplatid", &controllers.HomeController{}, "*:ChoosePlatId")

	//登录
	beego.Router("/login", &controllers.LoginController{}, "*:Login")
	beego.Router("/dologin", &controllers.LoginController{}, "Post:DoLogin")
	beego.Router("/loginout", &controllers.LoginController{}, "*:LoginOut")

	//权限管理
	//个人中心-密码修改
	beego.AutoRouter(&controllers.BackEndUserController{})

	//权限管理-用户组管理
	beego.AutoRouter(&controllers.GroupController{})
	//权限管理-权限
	beego.AutoRouter(&controllers.AuthController{})
	//日志
	beego.AutoRouter(&controllers.LogController{})

}
