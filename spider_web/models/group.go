/**********************************************
** @Des:
** @Author: wangliangde
** @Date:   2018/4/8 12:01
***********************************************/

package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/wliangde/goleond/spider_web/common"
)

func GetGroupList() []*Group {
	slcGroup := make([]*Group, 0)
	_, err := orm.NewOrm().QueryTable(TableName(common.TBNAME_GROUP)).Filter("status", 1).All(&slcGroup)
	if err != nil {
		logs.Error(err)
		return slcGroup
	}
	return slcGroup
}

func GetGroupListWebData() []map[string]interface{} {
	slcGroup := GetGroupList()
	if len(slcGroup) == 0 {
		return nil
	}
	list := make([]map[string]interface{}, 0, len(slcGroup))

	for _, ptGroup := range slcGroup {
		if ptGroup == nil {
			continue
		}
		row := make(map[string]interface{}, 0)
		row["id"] = ptGroup.Id
		row["name"] = ptGroup.Name
		row["desc"] = ptGroup.Desc
		list = append(list, row)
	}
	return list
}

func GetGroupById(nId int) *Group {
	ptGroup := &Group{Id: nId}
	if err := orm.NewOrm().Read(ptGroup); err != nil {
		logs.Error(err)
		return nil
	}
	return ptGroup
}

func GroupAdd(ptGroup *Group) (int, error) {
	id, err := orm.NewOrm().Insert(ptGroup)
	return int(id), err
}

type Group struct {
	Id          int
	Name        string
	Desc        string
	Status      bool
	BackendUser *BackendUser `orm:"reverse(one)"`  //设置反向关系
	GroupAuth   []*GroupAuth `orm:"reverse(many)"` // 设置一对多的反向关系
}

func (this Group) String() string {
	return fmt.Sprintf("id%d 名字:%s status:%d", this.Id, this.Name, this.Status)
}

func (this *Group) TableName() string {
	return TableName(common.TBNAME_GROUP)
}

func (this *Group) Update() error {
	_, err := orm.NewOrm().Update(this)
	return err
}

func (this *Group) GroupDel() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}

	ptGroupAuth := NewGroupAuth()
	ptGroupAuth.Group.Id = this.Id
	_, err := orm.NewOrm().Delete(ptGroupAuth)
	return err
}
