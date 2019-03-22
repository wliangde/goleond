/**********************************************
** @Author: liangde.wld
** @Desc:

***********************************************/

package common

import (
	"protocol/in/in_base"
)

//单服gm返回
type TSingleResult struct {
	Code      int         `json:"code"`       //返回码
	Msg       string      `json:"msg"`        //提示
	TotalSize uint32      `json:"total_size"` //如果data是list total_size表示总长度
	Data      interface{} `json:"data"`       //结果data
}

//服务器列表返回
type TServerData struct {
	Id   uint64
	Name string
	Type uint32
}
type TResultServerList struct {
	TSingleResult
	Data []*TServerData `json:"data"`
}

//服务器详细列表返回
type TResultServerDetail struct {
	TSingleResult
	Data []map[string]interface{} `json:"data"`
}

//服务器操作 返回
type TResultServerOper struct {
	TSingleResult
	Data map[string]interface{} `json:"data"`
}

//launcher 列表返回
type TResultLauncherList struct {
	TSingleResult
	Data []map[string]interface{} `json:"data"` //覆盖父类Data
}

//launcher配置返回
type TResultLauncherOper struct {
	TSingleResult
	Data map[string]interface{} `json:"data"` //覆盖父类Data
}

//一键养成返回
type TResultOneKeyCultivate struct {
	TSingleResult
	Data string `json:"data"` //成功返回养成成功的uuid
}

//查询玩家信息返回
type TUserInfo struct {
	Base           *in_base.SBase
	Sex            string
	CreateTime     string
	OnlineTime     string
	OfflineTime    string
	LastOnlineTime string
	State          string
	ServerName     string
	VipLevel       uint32
	//json
	Knights           string
	Equips            string //装备json
	DungeonBranchTask string
	ArenaRecords      string
	Artifacts         string
	Gems              string
}
type TResultGetUserInfo struct {
	TSingleResult
	Data TUserInfo `json:"data"`
}

//查询玩家信息返回
type TForbidUserInfo struct {
	Id       int
	Uuid     string
	Time     uint32
	FreeTime uint32
	Reason   string
}
type TResultForbidUserList struct {
	TSingleResult
	Data []*TForbidUserInfo `json:"data"`
}

type TResult struct {
	Code         int         `json:"code"`
	Msg          string      `json:"msg"`
	AllFailed    bool        `json:"all_failed"` //是否全部失败
	SuccessData  interface{} `json:"success_data"`
	FailDetail   interface{} `json:"fail_detail"`
	SuccessCount uint32      `json:"success_count"`
}
