/**********************************************
** @Des:
** @Author: wangliangde
** @Date:   2018/4/17 14:57
***********************************************/

package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/wliangde/goleond/spider_web/common"
	"time"
)

func GetLogList(page, pageSize int, filters ...interface{}) ([]*Log, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Log, 0)
	query := orm.NewOrm().QueryTable(TableName(common.TBNAME_LOG))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

func AddLog(ldwUid uint64, strName string, action common.ELogAction, dwPlatId uint32, ldwServerId uint64, ldwTargetId uint64, strCmd string, strParam, strUrl string, strErr string) {
	ptLog := NewLog(ldwUid, strName, action)

	ptLog.PlatId = dwPlatId
	ptLog.ServerId = ldwServerId
	ptLog.TargetId = ldwTargetId
	ptLog.Cmd = strCmd
	ptLog.Param = strParam
	ptLog.Url = strUrl
	ptLog.Err = strErr

	ptLog.Add()
}

type Log struct {
	Id       uint64
	Uid      uint64 //操作者uid
	Name     string //操作者名字
	Action   common.ELogAction
	PlatId   uint32
	ServerId uint64
	TargetId uint64
	Cmd      string
	Param    string
	Url      string
	Time     int64
	Err      string
}

func (this *Log) TableName() string {
	return TableName(common.TBNAME_LOG)
}

func NewLog(ldwUid uint64, strName string, action common.ELogAction) *Log {
	ptLog := &Log{Uid: ldwUid,
		Name:   strName,
		Action: action}
	ptLog.Time = time.Now().Unix()
	return ptLog
}
func (this *Log) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}
