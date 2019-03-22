package controllers

import (
	"fmt"
	"gmweb/common"
	"gmweb/models"
	"strings"
)

type AuthController struct {
	BaseController
}

func (this *AuthController) Index() {
	this.Data["pageTitle"] = "权限因子"
	this.display()
}

func (this *AuthController) List() {
	this.Data["zTree"] = true //引入ztreecss
	this.Data["pageTitle"] = "权限因子"
	this.display()
}

//获取全部节点
func (this *AuthController) GetNodes() {
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, count := models.AuthGetList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["pId"] = v.Pid
		row["name"] = v.Name
		row["open"] = true
		list[k] = row
	}

	this.ajaxList("成功", common.MSG_OK, count, list)
}

//获取一个节点
func (this *AuthController) GetNode() {
	id, _ := this.GetInt("id")
	result := models.AuthGetById(id)

	row := make(map[string]interface{})
	row["id"] = result.Id
	row["pid"] = result.Pid
	row["auth_name"] = result.Name
	row["auth_url"] = result.AuthUrl
	row["sort"] = result.Sort
	row["is_show"] = result.SiteBar
	row["icon"] = result.Icon
	fmt.Println(row)
	this.ajaxList("成功", common.MSG_OK, 0, row)
}

//新增或修改
func (this *AuthController) AjaxSave() {
	auth := new(models.Auth)
	auth.Pid, _ = this.GetInt("pid")
	auth.Name = strings.TrimSpace(this.GetString("auth_name"))
	auth.AuthUrl = strings.TrimSpace(this.GetString("auth_url"))
	auth.Sort, _ = this.GetInt("sort")
	auth.SiteBar, _ = this.GetInt("is_show")
	auth.Icon = strings.TrimSpace(this.GetString("icon"))
	auth.Status = 1
	id, _ := this.GetInt("id")
	if id == 0 {
		//新增
		if _, err := models.AuthAdd(auth); err != nil {
			this.ajaxMsg(err.Error(), common.MSG_ERR)
		}
	} else {
		auth.Id = id
		if err := auth.Update(); err != nil {
			this.ajaxMsg(err.Error(), common.MSG_ERR)
		}
	}

	this.ajaxMsg("", common.MSG_OK)
}

//删除
func (this *AuthController) AjaxDel() {
	id, _ := this.GetInt("id")
	auth := models.AuthGetById(id)
	auth.Id = id
	auth.Status = 0
	if err := auth.Update(); err != nil {
		this.ajaxMsg(err.Error(), common.MSG_ERR)
	}
	this.ajaxMsg("", common.MSG_OK)
}
