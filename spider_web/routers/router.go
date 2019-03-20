package routers

import (
	"github.com/astaxie/beego"
	"github.com/wliangde/goleond/spider_web/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.HomeController{}, "*:Index")
	beego.Router("/home", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")

}
