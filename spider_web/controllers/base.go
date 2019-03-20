/**
@user:          liangde.wld
@createtime:    2019/3/19 9:09 PM
@desc:
**/
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {

	this.Data["user"] = "wld"
	logs.Info("prepare")
}
