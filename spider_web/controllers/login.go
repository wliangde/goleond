/**********************************************
** @Des:
** @Author: wangliangde
** @Date:   2018/4/3 16:25
***********************************************/

package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/wliangde/goleond/spider_web/models"
	"strings"
	"time"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Login() {
	//this.GetSession("sesuser")
	if this.ptCurUser.Id != 0 {
		//登录过直接进入主页
		this.redirect(beego.URLFor("HomeController.Index"))
	}
	//登录
	this.TplName = "login/login.html"
}

func (this *LoginController) DoLogin() {
	strUserName := strings.TrimSpace(this.GetString("username"))
	strPwd := strings.TrimSpace(this.GetString("password"))
	ptFlash := beego.NewFlash()
	strErrMsg := ""
	for {
		if len(strUserName) == 0 || len(strPwd) == 0 {
			strErrMsg = "用户名密码不能为空"
			break
		}
		ptUser, err := models.GetUserByName(strUserName)
		if err != nil {
			strErrMsg = "找不到对应用户"
			break
		}
		if ptUser.CheckPwd(strPwd) == false {
			strErrMsg = "密码错误"
			break
		}
		if ptUser.Status == -1 {
			strErrMsg = "账号被禁用"
			break
		}

		ptUser.LastPlatId = 1 //TODO 默认是平台1
		ptUser.LastLogin = time.Now().Unix()
		ptUser.LastIp = this.getClientIp()
		ptUser.Update()
		//authkey := libs.Md5([]byte(self.getClientIp() + "|" + user.Password + user.Salt))
		this.setUser2Session(ptUser)
		logs.Info("玩家登录成功user%s", ptUser)
		this.redirect(beego.URLFor("HomeController.Index"))
		break
	}

	if len(strErrMsg) > 0 {
		logs.Error("[登录] 失败username:%s msg:%s", strUserName, strErrMsg)
		ptFlash.Error(strErrMsg)
		ptFlash.Store(&this.Controller)
	}
	this.TplName = "login/login.html"
}

func (this *LoginController) LoginOut() {
	this.BaseController.LoginOut()
}
