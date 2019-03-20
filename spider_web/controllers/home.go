/**
@user:          liangde.wld
@createtime:    2019/3/19 9:12 PM
@desc:
**/
package controllers

type HomeController struct {
	BaseController
}

func (this *HomeController) Index() {
	this.TplName = "index.html"
	this.Data["d"] = "wld"
}

func (this *HomeController) Start() {
	this.TplName = "/home/start.html"
}
