/**********************************************
** @Des:
** @Author: wangliangde
** @Date:   2018/4/8 12:02
***********************************************/

package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/wliangde/goleond/spider_web/common"
)

type GroupAuth struct {
	Id    uint32
	Group *Group `orm:"rel(fk)"`  //外键
	Auth  *Auth  `orm:"rel(fk)" ` // 外键
}

func NewGroupAuth() *GroupAuth {
	return &GroupAuth{
		Group: new(Group),
		Auth:  new(Auth),
	}
}

func (this *GroupAuth) TableName() string {
	return TableName(common.TBNAME_GROUP_AUTH)
}

//根据group id 查询权限
func GetAuthList(dwGroupId int) []*Auth {
	slcAuthList := make([]*GroupAuth, 0, 2)
	//auth__sort 对auth表的sort排序
	dwNum, err := orm.NewOrm().QueryTable(TableName(common.TBNAME_GROUP_AUTH)).Filter("Group", dwGroupId).RelatedSel("Auth").OrderBy("auth__sort").All(&slcAuthList)
	if err != nil {
		logs.Error("查询所有权限错误 %s", err)
		return nil
	}

	slcResult := make([]*Auth, 0, dwNum)
	for _, ptGroupAutch := range slcAuthList {
		if ptGroupAutch == nil || ptGroupAutch.Auth == nil {
			continue
		}
		logs.Debug("获得权限，组%s 权限%s", ptGroupAutch.Group, ptGroupAutch.Auth)
		slcResult = append(slcResult, ptGroupAutch.Auth)
	}
	return slcResult
}

func GetAllAuth() []*Auth {
	slcAuthList := make([]*Auth, 0, 2)
	_, err := orm.NewOrm().QueryTable(TableName(common.TBNAME_AUTH)).OrderBy("sort").All(&slcAuthList)
	if err != nil {
		logs.Error("查询所有权限错误 %s", err)
		return nil
	}
	return slcAuthList
}

func GroupAuthAdd(ptGroupAuth *GroupAuth) {
	_, err := orm.NewOrm().Insert(ptGroupAuth)
	if err != nil {
		logs.Error(err)
	}
}

func GroupAuthDelete(id int) (int64, error) {
	query := orm.NewOrm().QueryTable(TableName(common.TBNAME_GROUP_AUTH))
	return query.Filter("group_id", id).Delete()
}
