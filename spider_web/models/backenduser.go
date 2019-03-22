package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/wliangde/goleond/spider_web/common"
	"github.com/wliangde/goleond/spider_web/utils"
)

func GetUserByName(strName string) (*BackendUser, error) {
	ptUser := new(BackendUser)
	err := orm.NewOrm().QueryTable(TableName(common.TBNAME_BACKEND_USER)).Filter("name", strName).One(ptUser)
	if err != nil {
		return nil, err
	}
	return ptUser, nil
}

func GetUserById(ldwId uint64) (*BackendUser, error) {
	ptUser := new(BackendUser)
	err := orm.NewOrm().QueryTable(TableName(common.TBNAME_BACKEND_USER)).Filter("id", ldwId).One(ptUser)
	if err != nil {
		return nil, err
	}
	return ptUser, nil
}

//获取所有后台账号
func GetUserList(nPage, nPageSize int, iFilters ...interface{}) ([]*BackendUser, int64) {
	if nPage < 1 {
		nPage = 1
	}
	nOffset := (nPage - 1) * nPageSize
	qs := orm.NewOrm().QueryTable(TableName(common.TBNAME_BACKEND_USER)).RelatedSel()

	nFiltersLen := len(iFilters)
	if nFiltersLen%2 != 0 {
		return nil, 0
	}
	for i := 0; i < nFiltersLen; i += 2 {
		qs = qs.Filter(iFilters[i].(string), iFilters[i+1])
	}
	lnTotal, _ := qs.Count()
	slcUser := make([]*BackendUser, 0, lnTotal)
	qs.OrderBy("id").Limit(nPageSize, nOffset).All(&slcUser)
	return slcUser, lnTotal
}

func AddUser(ptUser *BackendUser) error {
	ldwId, err := orm.NewOrm().Insert(ptUser)
	if err != nil {
		return err
	}
	ptUser.Id = uint64(ldwId)
	return nil
}

func GetBackendUserByName(strName string) (*BackendUser, error) {
	ptUser := &BackendUser{Name: strName}
	err := orm.NewOrm().Read(ptUser)
	if err != nil {
		return nil, err
	}
	return ptUser, nil
}

type BackendUser struct {
	Id         uint64 //唯一id
	Name       string //用户名
	RealName   string
	Pwd        string //密码
	Group      *Group `orm:"rel(one)"` //权限组
	IsSuper    bool   //是否是超级管理员
	Salt       string //密码盐
	LastLogin  int64  //上次登录时间
	LastIp     string //上次登录ip
	Status     int    //-1 禁用
	CreateTime int64  //创建时间
	UpdateTime int64  //更新时间
	LastPlatId uint32 `orm:"-"` //权限组
}

func NewBackEndUser() *BackendUser {
	return &BackendUser{
		Group: new(Group),
	}
}

func (this BackendUser) String() string {
	dwGroupId := 0
	if this.Group != nil {
		dwGroupId = this.Group.Id
	}
	return fmt.Sprintf("uid[%d] name[%s] 组[%d] LastPlatId:%d", this.Id, this.Name, dwGroupId, this.LastPlatId)
}

func (this *BackendUser) TableName() string {
	return TableName(common.TBNAME_BACKEND_USER)
}

func (this *BackendUser) CheckPwd(strPwd string) bool {
	strMy := utils.String2md5(strPwd, this.Salt)
	if this.Pwd != strMy {
		logs.Debug("[登录] %s %s", this, strMy)
		return false
	}
	return true
}

//更新
func (this *BackendUser) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		logs.Error("[登录] 更新玩家失败 %s %s", this, err)
		return err
	}
	logs.Error("[登录] 更新玩家成功 %s last_login:%d last_ip:%s", this, this.LastLogin, this.LastIp)
	return nil
}

func (this *BackendUser) ToWebData() map[string]interface{} {
	data := make(map[string]interface{}, 0)
	data["id"] = this.Id
	data["name"] = this.Name
	data["real_name"] = this.RealName
	data["group_id"] = this.Group.Id
	if ptGroup := GetGroupById(this.Group.Id); ptGroup != nil {
		this.Group = ptGroup
	}
	data["group_name"] = this.Group.Name
	return data
}
