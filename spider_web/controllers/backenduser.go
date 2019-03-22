/**********************************************
** @Des:
** @Author: wangliangde
** @Date:   2018/4/8 15:27
***********************************************/

package controllers

import (
	"github.com/astaxie/beego/logs"
	"gmweb/common"
	"gmweb/models"
	"gmweb/utils"
	"strconv"
	"strings"
	"time"
)

type BackEndUserController struct {
	BaseController
}

//个人中心-密码修改
func (this *BackEndUserController) Pwd() {
	this.Data["pageTitle"] = "资料修改"
	row := make(map[string]interface{})
	row["id"] = this.ptCurUser.Id
	row["login_name"] = this.ptCurUser.Name
	row["last_login"] = utils.Unix2TimeStr(this.ptCurUser.LastLogin)
	row["last_ip"] = this.ptCurUser.LastIp
	this.Data["admin"] = row
	this.display()
}

//个人中心-密码修改-提交
func (this *BackEndUserController) PwdSave() {
	id, _ := this.GetUint64("id")
	if this.ptCurUser.Id != id {
		return
	}
	//修改
	ptUser := this.ptCurUser

	resetPwd := this.GetString("reset_pwd")
	if resetPwd != "1" {
		this.StopRun()
	}

	pwdOld := strings.TrimSpace(this.GetString("password_old"))
	pwdOldMd5 := utils.String2md5(pwdOld, ptUser.Salt)
	if ptUser.Pwd != pwdOldMd5 {
		this.ajaxMsg("旧密码错误", common.MSG_ERR)
	}

	pwdNew1 := strings.TrimSpace(this.GetString("password_new1"))
	pwdNew2 := strings.TrimSpace(this.GetString("password_new2"))

	if pwdNew1 != pwdNew2 {
		this.ajaxMsg("两次密码不一致", common.MSG_ERR)
	}
	pwd, salt := utils.Password(4, pwdNew1)
	ptUser.Pwd = pwd
	ptUser.Salt = salt

	if err := ptUser.Update(); err != nil {
		this.ajaxMsg(err.Error(), common.MSG_ERR)
	}

	this.setUser2Session(this.ptCurUser)
	this.ajaxMsg("", common.MSG_OK)
}

//权限管理-管理员列表
func (this *BackEndUserController) List() {
	this.Data["pageTitle"] = "用户管理"
	this.display()
}

//权限管理-用户管理-列表
func (this *BackEndUserController) Table() {
	//列表
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 30
	}

	name := strings.TrimSpace(this.GetString("realName"))

	StatusText := make(map[int]string)
	StatusText[0] = "<font color='red'>禁用</font>"
	StatusText[1] = "正常"

	this.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	//
	if name != "" {
		filters = append(filters, "name", name)
	}
	result, count := models.GetUserList(page, this.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["login_name"] = v.Name
		row["real_name"] = v.RealName
		row["group_name"] = v.Group.Name
		row["role_ids"] = v.Group.Id
		row["status"] = v.Status
		row["status_text"] = StatusText[v.Status]
		list[k] = row
	}
	this.ajaxList("成功", common.MSG_OK, count, list)
}

//权限管理-用户管理-新增
func (this *BackEndUserController) Add() {
	this.Data["pageTitle"] = "新增管理员"
	//
	slcGroupData := models.GetGroupListWebData()
	if slcGroupData != nil {
		this.Data["Groups"] = slcGroupData
	}
	this.display()
}

//权限管理-用户管理-新增
func (this *BackEndUserController) AjaxAddNewUser() {
	ptUser := models.NewBackEndUser()
	ptUser.Name = strings.TrimSpace(this.GetString("login_name"))
	ptUser.RealName = strings.TrimSpace(this.GetString("real_name"))
	ptUser.Group.Id, _ = this.GetInt("group")
	ptUser.CreateTime = time.Now().Unix()
	ptUser.UpdateTime = ptUser.CreateTime
	ptUser.Status = 1
	// 检查登录名是否已经存在
	_, err := models.GetBackendUserByName(ptUser.Name)
	if err == nil {
		this.ajaxMsg("登录名已经存在", common.MSG_ERR)
	}
	//新增
	pwd, salt := utils.Password(4, "")
	ptUser.Pwd = pwd
	ptUser.Salt = salt
	ptUser.IsSuper = false
	if ptUser.Group.Id == 1 { //1 默认是超级管理员
		ptUser.IsSuper = true
	}
	if err := models.AddUser(ptUser); err != nil {
		this.ajaxMsg(err.Error(), common.MSG_ERR)
	}
	this.ajaxMsg("", common.MSG_OK)
}

//权限管理-用户管理-新增
func (this *BackEndUserController) AjaxEditUser() {
	id, _ := this.GetUint64("id")

	ptUser, _ := models.GetUserById(id)
	//修改
	ptUser.UpdateTime = time.Now().Unix()
	ptUser.Group.Id, _ = this.GetInt("group")
	ptUser.RealName = strings.TrimSpace(this.GetString("real_name"))

	resetPwd, _ := this.GetInt("reset_pwd")
	if resetPwd == 1 {
		pwd, salt := utils.Password(4, "")
		ptUser.Pwd = pwd
		ptUser.Salt = salt
	}
	if err := ptUser.Update(); err != nil {
		this.ajaxMsg(err.Error(), common.MSG_ERR)
	}
	this.ajaxMsg(strconv.Itoa(resetPwd), common.MSG_OK)
}

//权限管理-用户管理-编辑
func (this *BackEndUserController) Edit() {
	this.Data["pageTitle"] = "编辑管理员"
	ldwId, _ := this.GetUint64("id", 0)

	ptUser, err := models.GetUserById(ldwId)
	if err != nil {
		logs.Error(err)
		return
	}
	this.Data["backenduser"] = ptUser.ToWebData()

	slcGroupData := models.GetGroupListWebData()
	if slcGroupData != nil {
		this.Data["Groups"] = slcGroupData
	}
	this.display()
}

//封禁账号
func (this *BackEndUserController) AjaxDel() {
	ldwUserId, _ := this.GetUint64("id")
	status := strings.TrimSpace(this.GetString("status"))
	ptUser_status := 0
	if status == "enable" {
		ptUser_status = 1
	}
	ptUser, _ := models.GetUserById(ldwUserId)
	if ptUser == nil {
		this.ajaxMsg("找不到对应后台账号", common.MSG_ERR)
	}

	if ptUser.IsSuper {
		this.ajaxMsg("超级管理员不允许操作", common.MSG_ERR)
	}

	ptUser.Status = ptUser_status
	if err := ptUser.Update(); err != nil {
		this.ajaxMsg(err.Error(), common.MSG_ERR)
	}
	this.ajaxMsg("操作成功", common.MSG_OK)
}
