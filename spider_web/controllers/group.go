/**********************************************
** @Des:
** @Author: wangliangde
** @Date:   2018/4/8 15:27
***********************************************/

package controllers

import (
	"fmt"
	"gmweb/common"
	"gmweb/models"
	"strconv"
	"strings"
)

type GroupController struct {
	BaseController
}

func (this *GroupController) List() {
	this.Data["pageTitle"] = "用户组管理"
	this.display()
}

func (this *GroupController) Add() {
	this.Data["zTree"] = true //引入ztreecss
	this.Data["pageTitle"] = "新增用户组"
	this.display()
}

func (this *GroupController) Edit() {
	this.Data["zTree"] = true //引入ztreecss
	this.Data["pageTitle"] = "编辑用户组"

	id, _ := this.GetInt("id", 0)
	role := models.GetGroupById(id)
	row := make(map[string]interface{})
	row["id"] = role.Id
	row["role_name"] = role.Name
	row["detail"] = role.Desc
	this.Data["role"] = row

	//获取选择的树节点
	roleAuth := models.GetAuthList(id)
	authId := make([]int, 0)
	for _, v := range roleAuth {
		authId = append(authId, v.Id)
	}
	this.Data["auth"] = authId
	fmt.Println(authId)
	this.display()
}

func (this *GroupController) AjaxSave() {
	role := new(models.Group)
	role.Name = strings.TrimSpace(this.GetString("role_name"))
	role.Desc = strings.TrimSpace(this.GetString("desc"))
	role.Status = true
	auths := strings.TrimSpace(this.GetString("nodes_data"))
	role_id, _ := this.GetInt("id")
	if role_id == 0 {
		//新增
		if id, err := models.GroupAdd(role); err != nil {
			this.ajaxMsg(err.Error(), common.MSG_ERR)
		} else {
			authsSlice := strings.Split(auths, ",")
			for _, v := range authsSlice {
				ra := models.NewGroupAuth()
				aid, _ := strconv.Atoi(v)
				ra.Group.Id = id
				ra.Auth.Id = aid
				models.GroupAuthAdd(ra)
			}
		}
		this.ajaxMsg("", common.MSG_OK)
	}
	//修改
	role.Id = role_id
	if err := role.Update(); err != nil {
		this.ajaxMsg(err.Error(), common.MSG_ERR)
	} else {
		// 删除该用户组权限
		models.GroupAuthDelete(role_id)
		authsSlice := strings.Split(auths, ",")
		for _, v := range authsSlice {
			if v == "" {
				continue
			}
			ra := models.NewGroupAuth()
			aid, _ := strconv.Atoi(v)
			ra.Group.Id = role_id
			ra.Auth.Id = aid
			models.GroupAuthAdd(ra)
		}

	}
	this.ajaxMsg("", common.MSG_OK)
}

func (this *GroupController) AjaxDel() {
	role_id, _ := this.GetInt("id")
	role := models.GetGroupById(role_id)
	//role.Status = false
	//role.Id = role_id
	//if err := role.Update(); err != nil {
	//	this.ajaxMsg(err.Error(), _const.MSG_ERR)
	//}
	if role == nil {
		this.ajaxMsg("找不到用户组", common.MSG_ERR)
	}
	if err := role.GroupDel(); err != nil {
		this.ajaxMsg(err.Error(), common.MSG_ERR)
	}
	this.ajaxMsg("", common.MSG_OK)
}

func (this *GroupController) Table() {
	//列表
	list := models.GetGroupListWebData()
	this.ajaxList("成功", common.MSG_OK, int64(len(list)), list)
}
