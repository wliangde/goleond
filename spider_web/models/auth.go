/**********************************************
** @Des:
** @Author: wangliangde
** @Date:   2018/4/8 11:35
***********************************************/

package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/wliangde/goleond/spider_web/common"
)

func AuthGetList(page, pageSize int, filters ...interface{}) ([]*Auth, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Auth, 0)
	query := orm.NewOrm().QueryTable(TableName(common.TBNAME_AUTH))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("pid", "sort").Limit(pageSize, offset).All(&list)

	return list, total
}

func AuthGetById(nId int) *Auth {
	ptAuth := &Auth{Id: nId}
	if err := orm.NewOrm().Read(ptAuth); err != nil {
		return nil
	}
	return ptAuth
}

type Auth struct {
	Id        int
	Pid       int
	Name      string
	AuthUrl   string
	Sort      int
	Icon      string
	Status    int
	SiteBar   int          //是否在侧边栏显示
	GroupAuth []*GroupAuth `orm:"reverse(many)"` // 设置一对多的反向关系
}

func (this Auth) String() string {
	return fmt.Sprintf("id:%d pid:%d 名字:%s AuthUrl:%s 排序:%d status:%d", this.Id, this.Pid, this.Name, this.AuthUrl, this.Sort, this.Status)
}

func (this *Auth) TableName() string {
	return TableName(common.TBNAME_AUTH)
}

func (this *Auth) Update() error {
	_, err := orm.NewOrm().Update(this)
	return err
}

func (this *Auth) Del() error {
	_, err := orm.NewOrm().Delete(this)
	return err
}

func AuthAdd(ptAuth *Auth) (int64, error) {
	return orm.NewOrm().Insert(ptAuth)
}
