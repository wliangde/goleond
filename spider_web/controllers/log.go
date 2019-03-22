/**********************************************
** @Des:
** @Author: wangliangde
** @Date:   2018/4/17 15:31
***********************************************/

package controllers

import (
	"gmweb/common"
	"gmweb/models"
	"gmweb/utils"
	"strings"
)

type LogController struct {
	BaseController
}

func (this *LogController) List() {
	this.Data["pageTitle"] = "操作记录"
	this.display()
}

func (this *LogController) Table() {
	//列表
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 50
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

	result, count := models.GetLogList(page, this.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["time"] = utils.Unix2TimeStr(v.Time)
		row["id"] = v.Id
		row["uid"] = v.Uid
		row["name"] = v.Name
		row["action"] = v.Action
		row["plat_id"] = v.PlatId
		row["server_id"] = v.ServerId
		row["target_id"] = v.TargetId
		row["cmd"] = v.Cmd
		row["param"] = v.Param
		row["url"] = v.Url
		row["err"] = v.Err
		list[k] = row
	}
	this.ajaxList("成功", common.MSG_OK, count, list)
}
