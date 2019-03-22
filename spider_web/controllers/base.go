/**********************************************
** @Des:
** @Author: wangliangde
** @Date:   2018/4/3 16:20
***********************************************/

package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/wliangde/goleond/spider_web/common"
	"github.com/wliangde/goleond/spider_web/models"
	"strings"
)

type BaseController struct {
	beego.Controller
	strControllerName string              //controller 类名字
	strActionName     string              //controller 函数名字
	ptCurUser         *models.BackendUser //session缓存的登录玩家
	allowUrl          string              //允许方位的url
	strCurUrl         string              //当前访问的url
	pageSize          int                 //表每页默认大小
}

func (this *BaseController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

func (this *BaseController) Prepare() {
	this.ptCurUser = models.NewBackEndUser() //先初始化一个，防止宕机

	//分页情况默认每页大小
	this.pageSize = 50
	controllerName, actionName := this.GetControllerAndAction()
	//去掉Controller
	this.strControllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.strActionName = strings.ToLower(actionName)
	this.strCurUrl = fmt.Sprintf("/%s/%s;", this.strControllerName, this.strActionName)
	this.Data["version"] = beego.AppConfig.String("version")
	this.Data["siteName"] = beego.AppConfig.String("site.name")
	this.Data["curRoute"] = this.strControllerName + "." + this.strActionName
	this.Data["curController"] = this.strControllerName
	this.Data["curAction"] = this.strActionName
	//从session里获取登录user
	this.getUserFromSession()
	//check session
	this.checkLogin()
	//对登录玩家进行认证，并返回侧边栏信息
	this.auth()

	this.Data["loginUserId"] = this.ptCurUser.Id
	this.Data["loginUserName"] = this.ptCurUser.Name

	//ptPlatCfg, err := models.GetPlatById(this.ptCurUser.LastPlatId)
	//if err == nil {
	//	this.Data["PlatName"] = ptPlatCfg.Name
	//}
	this.Data["PlatList"] = this.getPlatList()

}

//获取平台渠道列表
func (this *BaseController) getPlatList() []map[string]interface{} {
	list := make([]map[string]interface{}, 0)

	//for _, ptPlatCfg := range models.GetPlatList() {
	//	row := make(map[string]interface{}, 0)
	//	row["id"] = ptPlatCfg.Id
	//	row["name"] = ptPlatCfg.Name
	//	list = append(list, row)
	//}
	return list
}

func (this *BaseController) setUser2Session(ptUser *models.BackendUser) {
	//拉去权限
	this.ptCurUser = ptUser
	this.SetSession(common.SESSION_USER, ptUser)
}

func (this *BaseController) getUserFromSession() {
	a := this.GetSession(common.SESSION_USER)
	if a != nil {
		this.ptCurUser = a.(*models.BackendUser)
		this.Data["backenduser"] = a
	} else {
		return
	}

	logs.Debug("从 session获得玩家:%s", this.ptCurUser)
}

//获取用户IP地址
func (this *BaseController) getClientIp() string {
	s := this.Ctx.Request.RemoteAddr
	l := strings.LastIndex(s, ":")
	return s[0:l]
}

func (this *BaseController) auth() {
	//不需要认证的接口
	isNoAuth := strings.Contains(common.NO_AUTH, this.strActionName)
	if isNoAuth {
		return
	}

	// 左侧导航栏
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	//不是超级管理员，去获取权限，可缓存
	var slcAutch []*models.Auth
	if !this.ptCurUser.IsSuper && this.ptCurUser.Group.Id != common.GROUP_SUPERGM {
		//普通管理员
		slcAutch = models.GetAuthList(this.ptCurUser.Group.Id)
	} else {
		slcAutch = models.GetAllAuth()
	}
	list := make([]map[string]interface{}, len(slcAutch))
	list2 := make([]map[string]interface{}, len(slcAutch))
	allow_url := ""
	i, j := 0, 0
	for _, v := range slcAutch {
		if v.AuthUrl != " " {
			allow_url += v.AuthUrl + ";"
		}

		if v.SiteBar == 0 {
			continue
		}
		row := make(map[string]interface{})
		if v.Pid == 1 { //一级菜单
			row["Id"] = int(v.Id)
			row["Sort"] = v.Sort
			row["AuthName"] = v.Name
			row["AuthUrl"] = v.AuthUrl
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			list[i] = row
			i++
		}
		if v.Pid != 1 { //二级菜单
			row["Id"] = int(v.Id)
			row["Sort"] = v.Sort
			row["AuthName"] = v.Name
			row["AuthUrl"] = v.AuthUrl
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			list2[j] = row
			j++
		}
	}

	this.Data["SideMenu1"] = list[:i]  //一级菜单
	this.Data["SideMenu2"] = list2[:j] //二级菜单
	this.allowUrl = allow_url + "/home/index;/home/chooseplatid;"
	isHasAuth := strings.Contains(this.allowUrl, this.strCurUrl)
	if isHasAuth == false {
		logs.Warn("[访问] 抱歉您没有权限操作 %s url:%s allowurl:%s", this.ptCurUser.Name, this.strCurUrl, this.allowUrl)
		this.ajaxMsg("抱歉您没有权限操作", common.MSG_ERR)
		return
	}
}

//如果没有登录，则跳到登录页面
func (this *BaseController) checkLogin() {
	noAuth := "/loginin/dologin/loginout/"
	bNoCheck := strings.Contains(noAuth, this.strActionName)
	if !bNoCheck && this.ptCurUser.Id == 0 {
		if this.strCurUrl == "/home/index;" {
			//弹出登出界面重新登录
			this.LoginOut()
			return
		}

		logs.Warn("[访问] 请重新登录 %s url:%s", this.ptCurUser.Name, this.strCurUrl)

		//session认证失败或者已经过期请点击重新登录
		this.DelSession(common.SESSION_USER)
		this.TplName = "public/sessionout.html"
		this.Render()
		this.StopRun()
	}
}

func (this *BaseController) CheckIsRightLogin(ldwId uint64) {
	this.checkLogin()

	if this.ptCurUser.Id != ldwId {
		this.DelSession(common.SESSION_USER)
		this.TplName = "public/sessionout.html"
		this.Render()
		this.StopRun()
	}
}

//加载模板
func (this *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = this.strControllerName + "/" + this.strActionName + ".html"
	}
	this.Layout = "public/layout.html"
	this.TplName = tplname
}

//用layout2做模板，网页的弹出层用这个函数。
func (this *BaseController) display2(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = this.strControllerName + "/" + this.strActionName + ".html"
	}
	this.Layout = "public/layout2.html"
	this.TplName = tplname
}

//ajax返回
func (this *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

//ajax返回 列表
func (this *BaseController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

// 是否POST提交
func (self *BaseController) isPost() bool {
	return self.Ctx.Request.Method == "POST"
}

func (this *BaseController) getCheckIp(str string) string {
	strIp := strings.TrimSpace(this.GetString(str))
	if len(strIp) == 0 {
		this.ajaxMsg("ip为空", common.MSG_ERR)
	}
	return strIp
}

func (this *BaseController) getCheckPort(str string) uint32 {
	port, err := this.GetUint32(str)
	if port == 0 || err != nil {
		this.ajaxMsg("端口不正确", common.MSG_ERR)
	}
	return port
}

func (this *BaseController) LoginOut() {
	logs.Info("[登出] 用户:%s", this.ptCurUser)
	this.DelSession(common.SESSION_USER)
	this.redirect(beego.URLFor("LoginController.Login"))
}
