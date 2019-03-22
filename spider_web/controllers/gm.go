/**********************************************
** @Des:
** @Author: wangliangde
** @Date:   2018/4/8 17:48
***********************************************/

package controllers

//
//import (
//	basecommon "base/common"
//	"bytes"
//	"config"
//	"encoding/json"
//	"encoding/xml"
//	"fmt"
//	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/httplib"
//	"github.com/astaxie/beego/logs"
//	"gmweb/common"
//	"gmweb/config"
//	"gmweb/models"
//	"gmweb/utils"
//	"io/ioutil"
//	"net/url"
//	"os/exec"
//	"strconv"
//	"strings"
//	"sync"
//	"time"
//)
//
//type GmController struct {
//	BaseController
//
//	ldwServerId uint64
//	ldwUserId   uint64
//}
//
////导航-网页-接口玩家查询===POST接口查询玩家
//func (this *GmController) SearchUser() {
//	this.Data["pageTitle"] = "玩家查询"
//	ptProxyClient := NewProxyClient(common.EProxyClientType_Default, this.ptCurUser.LastPlatId)
//	if ptProxyClient == nil {
//		this.ajaxMsg("平台错误", common.MSG_ERR)
//	}
//
//	list, _ := ptProxyClient.GetServerList(common.EServerType_Game)
//	this.Data["ServerList"] = list
//	if this.isPost() {
//		this.doSearch()
//	}
//
//	this.Data["ServerId"] = this.ldwServerId
//	//资源列表
//	this.Data["ResourceList"] = this.getResourceList()
//	this.Data["GMList"] = gmconfig.G_ptCommonXmlConfig.PtGMCfg.SlcGM
//
//	this.display()
//}
//
////网页-查询玩家简单信息
//func (this *GmController) SearchUserBrief() {
//	dwPlatId := this.ptCurUser.LastPlatId
//	ldwServerId, _ := this.GetUint64("server_id")
//	ldwUserId, _ := this.GetUint64("user_id")
//	strName := this.GetString("user_name")
//
//	count := int64(0)
//	list := make([]map[string]interface{}, 0, 1)
//	for {
//		this.ldwServerId = ldwServerId
//		this.ldwUserId = ldwUserId
//		if ldwServerId == 0 {
//			break
//		}
//		//if ldwUserId == 0 {
//		//	break
//		//}
//		ptPlatCfg, err := models.GetPlatById(dwPlatId)
//		if err != nil {
//			this.ajaxMsg("平台错误", common.MSG_ERR)
//		}
//
//		ptProxyClient := NewProxyClient(common.EProxyClientType_Default, dwPlatId)
//		if ptProxyClient == nil {
//			this.ajaxMsg("平台错误", common.MSG_ERR)
//		}
//
//		ptResult, _ := ptProxyClient.GetUserInfo(ldwServerId, ldwUserId, strName)
//		if ptResult == nil {
//			break
//		}
//		row := make(map[string]interface{}, 0)
//
//		row["id"] = 1
//		row["plat_id"] = dwPlatId
//		row["plat_name"] = ptPlatCfg.Name
//		row["svr_id"] = ldwServerId
//		row["svr_name"] = ptResult.Data.ServerName
//		row["user_id"] = ptResult.Data.Base.GetUserId()
//		row["user_name"] = ptResult.Data.Base.GetName()
//		row["level"] = ptResult.Data.Base.GetLevel()
//		row["vip_level"] = gmconfig.G_ptCommonXmlConfig.PtVipLevelCfg.GetVipExp(ptResult.Data.Base.GetVipExp())
//
//		list = append(list, row)
//		break
//	}
//
//	this.ajaxList("成功", common.MSG_OK, count, list)
//}
//
////导航-网页-更新配置
//func (this *GmController) UpdateCfg() {
//	this.Data["pageTitle"] = "更新配置"
//	//ptProxyClient := NewProxyClient(common.EProxyClientType_Default, this.ptCurUser.LastPlatId)
//	//if ptProxyClient == nil {
//	//	this.ajaxMsg("平台错误", common.MSG_ERR)
//	//}
//	//
//	//list, _ := ptProxyClient.GetServerList(common.EServerType_Game)
//	//this.Data["ServerList"] = list
//	//this.Data["ServerId"] = this.ldwServerId
//	this.display()
//}
//
////AJAX-查找玩家
//func (this *GmController) doSearch() {
//	ldwServerId, _ := this.GetUint64("server_id")
//	ldwUserId, _ := this.GetUint64("user_id")
//	strName := this.GetString("user_name")
//	dwPlatId := this.ptCurUser.LastPlatId
//	errorMsg := ""
//	for {
//		this.ldwServerId = ldwServerId
//		this.ldwUserId = ldwUserId
//		if ldwServerId == 0 {
//			errorMsg = "请选择服务器id"
//			break
//		}
//		if ldwUserId == 0 && strName == "" {
//			errorMsg = "请输入角色id或者角色名"
//			break
//		}
//		ptProxyClient := NewProxyClient(common.EProxyClientType_Default, dwPlatId)
//		if ptProxyClient == nil {
//			this.ajaxMsg("平台错误", common.MSG_ERR)
//		}
//
//		ptResult, errMsg := ptProxyClient.GetUserInfo(ldwServerId, ldwUserId, strName)
//		if ptResult == nil {
//			errorMsg = errMsg
//			break
//		}
//		this.Data["User"] = ptResult.Data.Base
//		this.Data["Sex"] = ptResult.Data.Sex
//		this.Data["CreateTime"] = ptResult.Data.CreateTime
//		this.Data["OnlineTime"] = ptResult.Data.OnlineTime
//		this.Data["OfflineTime"] = ptResult.Data.OfflineTime
//		this.Data["LastOnlineTime"] = ptResult.Data.LastOnlineTime
//		this.Data["State"] = ptResult.Data.State
//		this.Data["VipLevel"] = ptResult.Data.VipLevel
//		this.Data["Knights"] = FormatJson(ptResult.Data.Knights)
//		this.Data["Equips"] = FormatJson(ptResult.Data.Equips)
//		this.Data["DungeonBranchTask"] = FormatJson(ptResult.Data.DungeonBranchTask)
//		this.Data["ArenaRecords"] = FormatJson(ptResult.Data.ArenaRecords)
//		this.Data["Artifacts"] = FormatJson(ptResult.Data.Artifacts)
//		this.Data["Gems"] = FormatJson(ptResult.Data.Gems)
//
//		break
//	}
//
//	if len(errorMsg) != 0 {
//		flash := beego.NewFlash()
//		flash.Error(errorMsg)
//		flash.Store(&this.Controller)
//		//this.redirect(beego.URLFor("GmController.SearchUser"))
//		return
//	}
//}
//
////AJAX-任意GM
//func (this *GmController) DoGM() {
//	strServerId := this.GetString("server_id")
//	strUserId := this.GetString("user_id")
//	if len(strUserId) == 0 {
//		this.ajaxMsg("玩家id为空", common.MSG_ERR)
//	}
//	strCmd := this.GetString("gmcmd")
//	strParam := this.GetString("param")
//	this.doGM(strServerId, strUserId, strCmd, strParam)
//}
//
////AJAX-从策划表获取-资源列表
//func (this *GmController) getResourceList() []map[string]interface{} {
//	list := make([]map[string]interface{}, 0)
//	if gmconfig.G_ptCommonXmlConfig == nil || gmconfig.G_ptCommonXmlConfig.PtResourceCfg == nil {
//		return list
//	}
//	gmconfig.G_ptCommonXmlConfig.PtResourceCfg.SortedForeach(func(dwKey uint32, ptData *config.TResourceItem) {
//		if ptData == nil {
//			return
//		}
//		row := make(map[string]interface{}, 0)
//		row["id"] = dwKey
//		row["name"] = ptData.PtXml.Name
//		list = append(list, row)
//	})
//	return list
//}
//
////http://127.0.0.1:8080/gm/getresourceidlist?pid=110000
////AJAX
//func (this *GmController) GetResourceIdList() {
//	dwPid, _ := this.GetUint32("pid")
//	ptResourceCfg := gmconfig.G_ptCommonXmlConfig.PtResourceCfg.GetResource(dwPid)
//	if ptResourceCfg == nil {
//		return
//	}
//	list := make([]map[string]interface{}, 0)
//	for _, tId2Name := range ptResourceCfg.SlcId2Name {
//		row := make(map[string]interface{}, 0)
//		row["id"] = tId2Name.GetId()
//		row["name"] = tId2Name.GetName()
//		list = append(list, row)
//	}
//	this.ajaxMsg(list, common.MSG_OK)
//}
//
////http://127.0.0.1:8080/gm/addresource?server_id=5&user_id=1000&res_key=101000&res_id=1011&res_num=1
////AJAX-加道具
//func (this *GmController) AddResource() {
//	strServerId := this.GetString("server_id")
//	strUserId := this.GetString("user_id")
//	if len(strUserId) == 0 {
//		this.ajaxMsg("玩家id为空", common.MSG_ERR)
//	}
//	dwResKey, _ := this.GetUint32("res_key")
//	dwResValue, _ := this.GetUint32("res_id")
//	dwResNum, _ := this.GetUint32("res_num")
//	if dwResNum == 0 {
//		this.ajaxMsg("资源数量为空", common.MSG_ERR)
//	}
//	ptResourceCfg := gmconfig.G_ptCommonXmlConfig.PtResourceCfg.GetResource(dwResKey)
//	if ptResourceCfg == nil {
//		this.ajaxMsg("资源类型错误", common.MSG_ERR)
//	}
//	dwResType := dwResKey / 1000
//	if ptResourceCfg.PtXml.Value > 0 {
//		dwResValue = ptResourceCfg.PtXml.Value
//	}
//	strParam := fmt.Sprintf(`{"restype1":%d,"resvalue1":%d,"ressize1":%d}`, dwResType, dwResValue, dwResNum)
//	this.doGM(strServerId, strUserId, "add_resource", strParam)
//}
//
////内部-执行gm
//func (this *GmController) doGM(strServerId, strUserId string, strCmd, strParam string) {
//	dwPlatId := this.ptCurUser.LastPlatId
//	ptProxyClient := NewProxyClient(common.EProxyClientType_Default, dwPlatId)
//	if ptProxyClient == nil {
//		this.ajaxMsg("平台错误", common.MSG_ERR)
//	}
//	errMsg, strUrl, _ := ptProxyClient.ExecUserGmcmd(strServerId, strUserId, strCmd, strParam)
//	ldwServerId, _ := strconv.ParseUint(strServerId, 10, 64)
//	ldwUserId, _ := strconv.ParseUint(strUserId, 10, 64)
//	models.AddLog(this.ptCurUser.Id, this.ptCurUser.Name, common.ELogActionGMAny, this.ptCurUser.LastPlatId, ldwServerId, ldwUserId, "", "", strUrl, errMsg)
//	if len(errMsg) > 0 {
//		this.ajaxMsg(errMsg, common.MSG_ERR)
//	}
//	this.ajaxMsg("处理成功", common.MSG_OK)
//}
//
////AJAX-更新配置
//func (this *GmController) AjaxUpdateCfg() {
//	dwPlatId := this.ptCurUser.LastPlatId
//	strServerId := this.GetString("server_id")
//	slcServerId := strings.Split(strServerId, ",")
//	if len(slcServerId) == 0 {
//		this.ajaxMsg("请选择服务器id", common.MSG_ERR)
//	}
//	ptProxyClient := NewProxyClient(common.EProxyClientType_Default, dwPlatId)
//	if ptProxyClient == nil {
//		this.ajaxMsg("平台错误", common.MSG_ERR)
//	}
//
//	logs.Info("[配置更新] 发送给gmproxy，服务器列表:%v ", slcServerId)
//	errMsg, strUrl := ptProxyClient.UpdateCfg(slcServerId)
//	models.AddLog(this.ptCurUser.Id, this.ptCurUser.Name, common.ELogActionUpdateCfg, dwPlatId, 0, 0, "", "", strUrl, errMsg)
//	if len(errMsg) > 0 {
//		logs.Info("[配置更新] 处理失败，服务器列表:%v ", slcServerId)
//		this.ajaxMsg(errMsg, common.MSG_ERR)
//	}
//	logs.Info("[配置更新] 处理成功，服务器列表:%v ", slcServerId)
//	this.ajaxMsg("处理成功", common.MSG_OK)
//}
//
////----------------GM邮件
////导航-网页-邮件
//func (this *GmController) Mail() {
//	this.Data["pageTitle"] = "邮件操作"
//	this.display()
//}
//
////网页-添加收件人
//func (this *GmController) AddRecv() {
//	this.Data["pageTitle"] = "添加收件人"
//	this.Data["PlatList"] = this.getPlatList()
//	this.display2()
//}
//
////网页-添加附件
//func (this *GmController) AddAttach() {
//	this.Data["pageTitle"] = "添加附件"
//
//	this.Data["ResourceList"] = this.getResourceList()
//	this.display2()
//}
//
////网页-选择服务器
//func (this *GmController) SelectServer() {
//	//this.Data["ResourceList"] = this.getResourceList()
//	this.Data["PlatList"] = this.getPlatList()
//	this.display2()
//}
//
//type TJsonPlatServer struct {
//	PlatId     uint32
//	Name       string
//	ServerList []*TJsonServer
//}
//
////AJAX-获取平台以及平台下的服列表
//func (this *GmController) PlatServer() {
//	slcPlatServer := make([]*TJsonPlatServer, 0)
//	for _, ptPlatCfg := range models.GetPlatList() {
//		ptPlatServer := &TJsonPlatServer{}
//		ptPlatServer.PlatId = ptPlatCfg.Id
//		ptPlatServer.Name = ptPlatCfg.Name
//		ptPlatServer.ServerList = this.GetServerJsonList(ptPlatServer.PlatId, common.EServerType_Game)
//
//		slcPlatServer = append(slcPlatServer, ptPlatServer)
//	}
//	this.ajaxMsg(slcPlatServer, common.MSG_OK)
//}
//
//////获取平台渠道列表
////func (this *GmController) getPlatList() []map[string]interface{} {
////	list := make([]map[string]interface{}, 0)
////	if gmconfig.G_ptCommonXmlConfig == nil || len(gmconfig.G_ptCommonXmlConfig.SlcPlatCfg) == 0 {
////		return list
////	}
////	for _, ptPlatCfg := range gmconfig.G_ptCommonXmlConfig.SlcPlatCfg {
////		row := make(map[string]interface{}, 0)
////		row["id"] = ptPlatCfg.PlatId
////		row["name"] = ptPlatCfg.Name
////		list = append(list, row)
////	}
////	return list
////}
//
////AJAX-发送系统邮件
//func (this *GmController) SendSysMail() {
//
//	//plat_id,server_id,server_id;plat_id,server_id,server_id
//	strPlatServers := strings.TrimSpace(this.GetString("plat_server"))
//	if len(strPlatServers) == 0 {
//		this.ajaxMsg("请选择渠道和服务器", common.MSG_ERR)
//	}
//
//	slcPlatServers := strings.Split(strPlatServers, ";")
//	if len(slcPlatServers) == 0 {
//		this.ajaxMsg("请选择渠道和服务器", common.MSG_ERR)
//	}
//
//	ptProxyMgr := NewMultiProxyMgr()
//
//	strIcon := this.GetString("icon")
//	strTitle := this.GetString("title")
//	strContent := this.GetString("content")
//	strFilterType := this.GetString("filter_type")
//	strFilterMinValue := this.GetString("filter_min_value")
//	strFilterMaxValue := this.GetString("filter_max_value")
//	if strFilterType == "" {
//		strFilterType = "0"
//		strFilterMinValue = "0"
//		strFilterMaxValue = "0"
//	}
//
//	//attach
//	strNewAttach := ""
//	strResGroup := this.GetString("attach") //type,value,num;type,value,num;
//	if len(strResGroup) > 0 {
//		slcResGroup := strings.Split(strResGroup, ";")
//		for _, resGroup := range slcResGroup {
//			slcRes := strings.Split(resGroup, ",")
//			if len(slcRes) != 3 {
//				this.ajaxMsg("附件参数有问题", common.MSG_ERR)
//			}
//
//			dwResKey, _ := strconv.ParseUint(slcRes[0], 10, 32)
//			dwResValue, _ := strconv.ParseUint(slcRes[1], 10, 32)
//			dwResNum, _ := strconv.ParseUint(slcRes[2], 10, 32)
//			if dwResNum == 0 {
//				this.ajaxMsg("资源数量为空", common.MSG_ERR)
//			}
//			ptResourceCfg := gmconfig.G_ptCommonXmlConfig.PtResourceCfg.GetResource(uint32(dwResKey))
//			if ptResourceCfg == nil {
//				this.ajaxMsg("资源类型错误", common.MSG_ERR)
//			}
//			dwResType := dwResKey / 1000
//			if ptResourceCfg.PtXml.Value > 0 {
//				dwResValue = uint64(ptResourceCfg.PtXml.Value)
//			}
//			newResGroup := fmt.Sprintf(`%d,%d,%d`, dwResType, dwResValue, dwResNum)
//
//			if len(strNewAttach) == 0 {
//				strNewAttach = newResGroup
//			} else {
//				strNewAttach += ";" + newResGroup
//			}
//		}
//	}
//
//	strParam := fmt.Sprintf(`{"icon":%s,"title":"%s","content":"%s","attach":"%s","filter_type":%s, "filter_min_value":%s,"filter_max_value":%s}`, strIcon, strTitle, strContent, strNewAttach, strFilterType, strFilterMinValue, strFilterMaxValue)
//	for _, strPlatServer := range slcPlatServers {
//		slcServer := strings.Split(strPlatServer, ",")
//		if len(slcServer) < 2 {
//			this.ajaxMsg("渠道和服务器错误", common.MSG_ERR)
//		}
//		dwPlatId, _ := strconv.ParseUint(slcServer[0], 10, 64) //第一个是渠道id，后面是这个渠道的服id
//		ptClient := ptProxyMgr.mutableProxyClient(common.EProxyClientType_MultiServer, uint32(dwPlatId))
//		if ptClient == nil {
//			this.ajaxMsg("平台错误", common.MSG_ERR)
//		}
//		for i := 1; i < len(slcServer); i++ {
//			ldwServerId, _ := strconv.ParseUint(slcServer[i], 10, 64)
//			if ldwServerId == 0 {
//				this.ajaxMsg("服务器id有问题", common.MSG_ERR)
//			}
//			ptProxyMgr.IncTotalCount()
//			ptClient.AddTarget(ldwServerId, 0)
//		}
//		ptClient.SetBackendUser(this.ptCurUser)
//		ptClient.SetGMParam("send_sys_mail", strParam)
//	}
//
//	msg := ptProxyMgr.SendMailAndWait()
//
//	this.ajaxMsg(msg, common.MSG_OK)
//}
//
////AJAX-发送玩家邮件
//func (this *GmController) SendUserMail() {
//
//	//plat_id,server_id,user_id;plat_id,server_id,user_id
//	strTargets := strings.TrimSpace(this.GetString("target"))
//	if len(strTargets) == 0 {
//		this.ajaxMsg("请选择收件人", common.MSG_ERR)
//	}
//
//	slcTargets := strings.Split(strTargets, ";")
//	if len(slcTargets) == 0 {
//		this.ajaxMsg("请选择收件人", common.MSG_ERR)
//	}
//
//	strIcon := this.GetString("icon")
//	strTitle := this.GetString("title")
//	strContent := this.GetString("content")
//
//	//attach
//	strNewAttach := ""
//	strResGroup := this.GetString("attach") //type,value,num;type,value,num;
//	if len(strResGroup) > 0 {
//		slcResGroup := strings.Split(strResGroup, ";")
//		for _, resGroup := range slcResGroup {
//			slcRes := strings.Split(resGroup, ",")
//			if len(slcRes) != 3 {
//				this.ajaxMsg("附件参数有问题", common.MSG_ERR)
//			}
//			dwResKey, _ := strconv.ParseUint(slcRes[0], 10, 32)
//			dwResValue, _ := strconv.ParseUint(slcRes[1], 10, 32)
//			dwResNum, _ := strconv.ParseUint(slcRes[2], 10, 32)
//			if dwResNum == 0 {
//				this.ajaxMsg("资源数量为空", common.MSG_ERR)
//			}
//			ptResourceCfg := gmconfig.G_ptCommonXmlConfig.PtResourceCfg.GetResource(uint32(dwResKey))
//			if ptResourceCfg == nil {
//				this.ajaxMsg("资源类型错误", common.MSG_ERR)
//			}
//			dwResType := dwResKey / 1000
//			if ptResourceCfg.PtXml.Value > 0 {
//				dwResValue = uint64(ptResourceCfg.PtXml.Value)
//			}
//			newResGroup := fmt.Sprintf(`%d,%d,%d`, dwResType, dwResValue, dwResNum)
//
//			if len(strNewAttach) == 0 {
//				strNewAttach = newResGroup
//			} else {
//				strNewAttach += ";" + newResGroup
//			}
//		}
//	}
//
//	strParam := fmt.Sprintf(`{"icon":%s,"title":"%s","content":"%s","attach":"%s"}`, strIcon, strTitle, strContent, strNewAttach)
//
//	ptProxyMgr := NewMultiProxyMgr()
//	for _, strPlatServer := range slcTargets {
//		slcServer := strings.Split(strPlatServer, ",")
//		if len(slcServer) < 3 {
//			this.ajaxMsg("收件人错误", common.MSG_ERR)
//		}
//		dwPlatId, _ := strconv.ParseUint(slcServer[0], 10, 64)
//		ptClient := ptProxyMgr.mutableProxyClient(common.EProxyClientType_MultiUser, uint32(dwPlatId))
//		if ptClient == nil {
//			this.ajaxMsg("平台错误", common.MSG_ERR)
//		}
//		ldwServerId, _ := strconv.ParseUint(slcServer[1], 10, 64)
//		ldwUserId, _ := strconv.ParseUint(slcServer[2], 10, 64)
//		ptProxyMgr.IncTotalCount()
//		ptClient.AddTarget(ldwServerId, ldwUserId)
//		ptClient.SetBackendUser(this.ptCurUser)
//		ptClient.SetGMParam("send_user_mail", strParam)
//	}
//
//	msg := ptProxyMgr.SendMailAndWait()
//	this.ajaxMsg(msg, common.MSG_OK)
//}
//
////---------------一键养成
////---一键养成界面
////导航-网页-一键养成
//func (this *GmController) OneKeyDev() {
//	this.Data["pageTitle"] = "一键养成"
//
//	ptProxyClient := NewProxyClient(common.EProxyClientType_Default, this.ptCurUser.LastPlatId)
//	if ptProxyClient == nil {
//		this.ajaxMsg("平台错误", common.MSG_ERR)
//	}
//
//	list, _ := ptProxyClient.GetServerList(common.EServerType_Game)
//	this.Data["ServerList"] = list
//	this.Data["ServerId"] = this.ldwServerId
//	this.display()
//}
//
////AJAX-一键养成
//func (this *GmController) DoOneKeyDev() {
//	dwPlatId := this.ptCurUser.LastPlatId
//	if dwPlatId == 0 {
//		this.ajaxMsg("请选择区", common.MSG_ERR)
//	}
//	strServerId := this.GetString("server_id")
//	if len(strServerId) == 0 {
//		this.ajaxMsg("请选择服务器id", common.MSG_ERR)
//	}
//
//	//养成配置, xml原始格式
//	strConfig := this.GetString("config")
//	strConfig = url.PathEscape(strConfig)
//	ptProxyClient := NewProxyClient(common.EProxyClientType_Default, dwPlatId)
//	if ptProxyClient == nil {
//		this.ajaxMsg("平台错误", common.MSG_ERR)
//	}
//	ptResult, errMsg, _ := ptProxyClient.OneKeyCultivate(strServerId, strConfig)
//	ldwSeverId, _ := strconv.ParseUint(strServerId, 10, 64)
//	models.AddLog(this.ptCurUser.Id, this.ptCurUser.Name, common.ELogActionOneKey, this.ptCurUser.LastPlatId, ldwSeverId, 0, "", "", "", errMsg)
//	if len(errMsg) > 0 {
//		this.ajaxMsg(errMsg, common.MSG_ERR)
//	}
//	this.ajaxMsg(ptResult.Data, common.MSG_OK)
//}
//
////AJAX-不操作，防止上传控件卡死用的
//func (this *GmController) UploadFile() {
//	this.ajaxMsg("", common.MSG_OK)
//}
//
///********************************
//发布版本
//*/
////导航-网页-客户端包版本
//func (this *GmController) VersionList() {
//	this.Data["pageTitle"] = "版本记录"
//	this.display()
//}
//
////AJAX-请求版本
//func (this *GmController) ReqVersionList() {
//	ptFileVersion := &config.TXmlVersionParaInfo{}
//	err := basecommon.LoadConfig(common.CLIENT_VERSION_FILE, ptFileVersion)
//	if err != nil {
//		this.ajaxList(err.Error(), common.MSG_ERR, 0, nil)
//	}
//
//	list := make([]map[string]interface{}, len(ptFileVersion.Datas))
//	for k, ptVersion := range ptFileVersion.Datas {
//		row := make(map[string]interface{})
//		row["id"] = k
//		row["plat_id"] = ptVersion.PlatForm
//		row["channel_id"] = ptVersion.Chan
//		row["patch"] = ptVersion.Patch
//		row["version"] = ptVersion.Version
//		row["patch_url"] = ptVersion.PatchUrl
//		row["min_version"] = ptVersion.MinVersion
//		row["min_url"] = ptVersion.MinUrl
//		row["release_time"] = utils.Unix2TimeStr(ptVersion.Time)
//		list[k] = row
//	}
//
//	count := int64(len(list))
//	this.ajaxList("成功", common.MSG_OK, count, list)
//}
//
////网页-发布版本页面
//func (this *GmController) ReleaseVersion() {
//	this.Data["pageTitle"] = "发布版本"
//	strPlatId := this.GetString("plat_id")
//	strChannelId := this.GetString("channel_id")
//	if len(strPlatId) > 0 && len(strChannelId) > 0 {
//		ptFileVersion := &config.TXmlVersionParaInfo{}
//		err := basecommon.LoadConfig(common.CLIENT_VERSION_FILE, ptFileVersion)
//		if err != nil {
//			return
//		}
//
//		for _, ptVersion := range ptFileVersion.Datas {
//			if ptVersion.PlatForm == strPlatId && ptVersion.Chan == strChannelId {
//				this.Data["plat_id"] = ptVersion.PlatForm
//				this.Data["channel_id"] = ptVersion.Chan
//				this.Data["patch"] = ptVersion.Patch
//				this.Data["version"] = ptVersion.Version
//				this.Data["patch_url"] = ptVersion.PatchUrl
//				this.Data["min_version"] = ptVersion.MinVersion
//				this.Data["min_url"] = ptVersion.MinUrl
//				this.Data["release_time"] = utils.Unix2TimeStr(ptVersion.Time)
//				break
//			}
//		}
//	}
//	this.display()
//}
//
////网页-发布版本渠道选择页面
//func (this *GmController) ReleaseVersionSelect() {
//	this.Data["PlatList"] = this.getPlatList()
//	this.display2()
//}
//
//type TJsonChannel struct {
//	ChannelId uint32
//	Name      string
//}
//
//type TJsonPlatChannel struct {
//	PlatId      uint32
//	Name        string
//	ChannelList []*models.Channel
//}
//
////AJAX-获取平台以及平台下的渠道列表
//func (this *GmController) PlatChannel() {
//	slcPlatChannel := make([]*TJsonPlatChannel, 0)
//	for _, ptPlatCfg := range models.GetPlatList() {
//		ptPlat := &TJsonPlatChannel{}
//		ptPlat.PlatId = ptPlatCfg.Id
//		ptPlat.Name = ptPlatCfg.Name
//		ptPlat.ChannelList, _ = models.GetChannelList(0, 100000, "plat_id", ptPlatCfg.Id)
//		slcPlatChannel = append(slcPlatChannel, ptPlat)
//	}
//	this.ajaxMsg(slcPlatChannel, common.MSG_OK)
//}
//
////AJAX-发布版本
//func (this *GmController) DoRelease() {
//	dwReSend, _ := this.GetUint32("resend")
//	if dwReSend > 0 {
//		this.ReRelease()
//		return
//	}
//	//plat_id,channel_id,channel_id;plat_id,server_id,channel_id
//	strPlatChannels := strings.TrimSpace(this.GetString("plat_channel"))
//	if len(strPlatChannels) == 0 {
//		this.ajaxMsg("请选择渠道", common.MSG_ERR)
//	}
//
//	slcPlatChannels := strings.Split(strPlatChannels, ";")
//	if len(slcPlatChannels) == 0 {
//		this.ajaxMsg("请选择渠道", common.MSG_ERR)
//	}
//
//	strPatch := this.GetString("patch")
//	strVersion := this.GetString("version")
//	strPatchUrl := this.GetString("patch_url")
//	strMinVersion := this.GetString("min_version")
//	strMinUrl := this.GetString("min_url")
//
//	bCbPatch := this.GetString("cb_patch") == "true"
//	bCbVersion := this.GetString("cb_version") == "true"
//	bCbPatchUrl := this.GetString("cb_patch_url") == "true"
//	bCbMinVersion := this.GetString("cb_min_version") == "true"
//	bCbMinUrl := this.GetString("cb_min_url") == "true"
//
//	if !(bCbPatch || bCbVersion || bCbPatchUrl || bCbMinVersion || bCbMinUrl) {
//		this.ajaxMsg("没有任何更新", common.MSG_ERR)
//	}
//
//	ptFileVersion := &config.TXmlVersionParaInfo{}
//	err := basecommon.LoadConfig(common.CLIENT_VERSION_FILE, ptFileVersion)
//	if err != nil {
//		//this.ajaxMsg("历史版本文件加载错误", common.MSG_ERR)
//	}
//
//	mapVersion := make(map[string]*config.TXmlVersionParaConfig, 0)
//
//	for _, strPlatChannel := range slcPlatChannels {
//		slcChannel := strings.Split(strPlatChannel, ",")
//		if len(slcChannel) < 2 {
//			this.ajaxMsg("渠道错误", common.MSG_ERR)
//		}
//		dwPlatId, _ := strconv.ParseUint(slcChannel[0], 10, 64) //第一个是平台Id
//		for i := 1; i < len(slcChannel); i++ {
//			dwChannelId, _ := strconv.ParseUint(slcChannel[i], 10, 64)
//			if dwChannelId == 0 {
//				this.ajaxMsg("渠道id有问题", common.MSG_ERR)
//			}
//			ptVersion := &config.TXmlVersionParaConfig{}
//			ptVersion.PlatForm = fmt.Sprintf("%d", dwPlatId)
//			ptVersion.Chan = fmt.Sprintf("%d", dwChannelId)
//
//			if bCbPatch {
//				ptVersion.Patch = strPatch
//			}
//			if bCbVersion {
//				ptVersion.Version = strVersion
//			}
//			if bCbPatchUrl {
//				ptVersion.PatchUrl = strPatchUrl
//			}
//			if bCbMinVersion {
//				ptVersion.MinVersion = strMinVersion
//			}
//			if bCbMinUrl {
//				ptVersion.MinUrl = strMinUrl
//			}
//
//			ptVersion.Time = time.Now().Unix()
//			strKey := fmt.Sprintf("%s:%s", ptVersion.PlatForm, ptVersion.Chan)
//			mapVersion[strKey] = ptVersion
//		}
//	}
//
//	for _, ptVersion := range ptFileVersion.Datas {
//		strKey := fmt.Sprintf("%s:%s", ptVersion.PlatForm, ptVersion.Chan)
//		if ptNewVersion, bOk := mapVersion[strKey]; bOk {
//			if bCbPatch {
//				ptVersion.Patch = ptNewVersion.Patch
//			}
//			if bCbVersion {
//				ptVersion.Version = ptNewVersion.Version
//			}
//			if bCbPatchUrl {
//				ptVersion.PatchUrl = ptNewVersion.PatchUrl
//			}
//			if bCbMinVersion {
//				ptVersion.MinVersion = ptNewVersion.MinVersion
//			}
//			if bCbMinUrl {
//				ptVersion.MinUrl = ptNewVersion.MinUrl
//			}
//			ptVersion.Time = ptNewVersion.Time
//			delete(mapVersion, strKey)
//		}
//	}
//	for _, ptVersion := range mapVersion {
//		ptFileVersion.Datas = append(ptFileVersion.Datas, ptVersion)
//	}
//
//	bytesVersion, err := xml.Marshal(ptFileVersion)
//	if err != nil {
//		this.ajaxMsg("xml序列化出错", common.MSG_ERR)
//	}
//	//logs.Info("序列化后的内容", strPatch, string(bytesVersion))
//
//	strMd5 := utils.String2md5(string(bytesVersion), common.VERSION_SERVER_SALT)
//	ptResult := Send2VersionServer(bytesVersion, strMd5)
//	if ptResult == nil {
//		this.ajaxMsg("找不到version服", common.MSG_ERR)
//	}
//	//只要有一个成功了就写文件
//	if ptResult.Success != 0 {
//		basecommon.SaveConfig(common.CLIENT_VERSION_FILE, ptFileVersion)
//	}
//	models.AddLog(this.ptCurUser.Id, this.ptCurUser.Name, common.ELogActionPublishVersion, this.ptCurUser.LastPlatId, 0, 0, "", "", "", "")
//	this.ajaxMsg(ptResult, common.MSG_OK)
//}
//
////AJAX-重新发布
//func (this *GmController) ReRelease() {
//	ptFileVersion := &config.TXmlVersionParaInfo{}
//	err := basecommon.LoadConfig(common.CLIENT_VERSION_FILE, ptFileVersion)
//	if err != nil {
//		this.ajaxMsg("找不到上次发布的版本", common.MSG_ERR)
//	}
//
//	bytesVersion, err := xml.Marshal(ptFileVersion)
//	if err != nil {
//		this.ajaxMsg("xml序列化出错", common.MSG_ERR)
//	}
//
//	strMd5 := utils.String2md5(string(bytesVersion), common.VERSION_SERVER_SALT)
//	ptResult := Send2VersionServer(bytesVersion, strMd5)
//	if ptResult == nil {
//		this.ajaxMsg("找不到version服", common.MSG_ERR)
//	}
//	//只要有一个成功了就写文件
//	if ptResult.Success != 0 {
//		utils.WriteFile(common.CLIENT_VERSION_FILE, bytesVersion)
//	}
//	models.AddLog(this.ptCurUser.Id, this.ptCurUser.Name, common.ELogActionRePublishVersion, this.ptCurUser.LastPlatId, 0, 0, "", "", "", "")
//	this.ajaxMsg(ptResult, common.MSG_OK)
//}
//
////AJAX-关闭服务器
//func (this *GmController) AjaxStopServer() {
//	dwPlatId := this.ptCurUser.LastPlatId
//	if dwPlatId == 0 {
//		this.ajaxMsg("请选择区", common.MSG_ERR)
//	}
//	strServerId := this.GetString("svr_id")
//	if len(strServerId) == 0 {
//		this.ajaxMsg("请选择服务器id", common.MSG_ERR)
//	}
//
//	ptProxyClient := NewProxyClient(common.EProxyClientType_Default, dwPlatId)
//	if ptProxyClient == nil {
//		this.ajaxMsg("平台错误", common.MSG_ERR)
//	}
//	errMsg, strUrl := ptProxyClient.StopServer(strServerId)
//
//	ldwSvrId, _ := strconv.ParseUint(strServerId, 10, 64)
//	models.AddLog(this.ptCurUser.Id, this.ptCurUser.Name, common.ELogActionStopServer, this.ptCurUser.LastPlatId, ldwSvrId, 0, "", "", strUrl, errMsg)
//
//	if len(errMsg) > 0 {
//		this.ajaxMsg(errMsg, common.MSG_ERR)
//	}
//	this.ajaxMsg("", common.MSG_OK)
//}
//
////AJAX-重启服务器
//func (this *GmController) AjaxRestartServer() {
//	dwPlatId := this.ptCurUser.LastPlatId
//	if dwPlatId == 0 {
//		this.ajaxMsg("请选择区", common.MSG_ERR)
//	}
//	strServerId := this.GetString("svr_id")
//	if len(strServerId) == 0 {
//		this.ajaxMsg("请选择服务器id", common.MSG_ERR)
//	}
//
//	ptProxyClient := NewProxyClient(common.EProxyClientType_Default, dwPlatId)
//	if ptProxyClient == nil {
//		this.ajaxMsg("平台错误", common.MSG_ERR)
//	}
//	errMsg, strUrl := ptProxyClient.RestartServer(strServerId)
//	ldwSvrId, _ := strconv.ParseUint(strServerId, 10, 64)
//	models.AddLog(this.ptCurUser.Id, this.ptCurUser.Name, common.ELogActionRestartServer, this.ptCurUser.LastPlatId, ldwSvrId, 0, "", "", strUrl, errMsg)
//
//	if len(errMsg) > 0 {
//		this.ajaxMsg(errMsg, common.MSG_ERR)
//	}
//	this.ajaxMsg("", common.MSG_OK)
//}
//
///********************************
//发布服务器版本
//*/
////导航-网页-服务器版本
//func (this *GmController) ServerVersion() {
//	this.Data["pageTitle"] = "服务器版本"
//	this.display()
//}
//
////AJAX-请求服务器版本
//func (this *GmController) AjaxReqSvrVersion() {
//	svrVerList, count := models.GetServerVersionList()
//	list := make([]map[string]interface{}, len(svrVerList))
//	for k, ptSvr := range svrVerList {
//		row := make(map[string]interface{})
//		row["id"] = ptSvr.Id
//		row["name"] = ptSvr.Name
//		row["svn"] = ptSvr.Svn
//		row["time"] = utils.Unix2TimeStr(ptSvr.Time)
//		row["sync_num"] = GetSyncNum(ptSvr.Name)
//		row["desc"] = ptSvr.Desc
//		list[k] = row
//	}
//
//	this.ajaxList("成功", common.MSG_OK, count, list)
//}
//
//func GetSyncNum(strVerName string) string {
//	strFilePath := "./data/server_version/" + strVerName + "/sync_num.txt"
//	slcBytes, _ := ioutil.ReadFile(strFilePath)
//	if len(slcBytes) > 0 {
//		return string(slcBytes)
//	}
//	return "0"
//}
//
////AJAX-发布服务器版本
//func (this *GmController) AjaxReleaseSvrVer() {
//	dwId, _ := this.GetUint32("id")
//	if dwId == 0 {
//		this.ajaxMsg("版本id为空", common.MSG_ERR)
//	}
//
//	ptSvrVer := models.GetServerVersionById(dwId)
//	if ptSvrVer == nil {
//		this.ajaxMsg("找不到版本配置", common.MSG_ERR)
//	}
//
//	strExecShell := fmt.Sprintf("%s %s %s", "./data/server_version/sync.sh", ptSvrVer.Svn, ptSvrVer.Name)
//	tExecShell := exec.Command("/bin/bash", "-c", strExecShell)
//	if slcRet, err := tExecShell.Output(); err != nil {
//		logs.Error("[GM-发布服务器版本] 脚本执行错误，shell:%s err:%s output:%s", strExecShell, err, slcRet)
//		this.ajaxMsg("同步错误", common.MSG_ERR)
//	}
//
//	ptSvrVer.Time = time.Now().Unix()
//	ptSvrVer.Update()
//
//	this.ajaxMsg("处理成功", common.MSG_OK)
//}
//
////AJAX-更新服务器
//func (this *GmController) AjaxUpdateServer() {
//	dwId, _ := this.GetUint32("id")
//	if dwId == 0 {
//		this.ajaxMsg("版本id为空", common.MSG_ERR)
//	}
//
//	ptSvrVer := models.GetServerVersionById(dwId)
//	if ptSvrVer == nil {
//		this.ajaxMsg("找不到版本配置", common.MSG_ERR)
//	}
//
//	slcClient := gmconfig.G_ptCommonXmlConfig.GetSvrSyncClient(ptSvrVer.Name)
//	if slcClient == nil {
//		this.ajaxMsg("不需要更新服务器", common.MSG_ERR)
//	}
//
//	var wg sync.WaitGroup
//
//	for _, ptSvrClient := range slcClient {
//		strUrl := fmt.Sprintf("%s/update/?name=%s", ptSvrClient.WebAddr, ptSvrClient.Name)
//
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//
//			client := httplib.Get(strUrl).SetTimeout(common.HTTP_CONNECT_TIMEOUT, 4*time.Minute)
//			ptRes, err := client.Response()
//			if err != nil {
//				logs.Error("[更新服务器] 发送给同步服错出错，url:%s %s", strUrl, err)
//				return
//			}
//			ptRes.Body.Close()
//			logs.Info("[更新服务器] 发送给同步服成功，url:%s ", strUrl)
//		}()
//	}
//
//	wg.Wait()
//	this.ajaxMsg("处理成功", common.MSG_OK)
//}
//
///********************************
//服务器版本客户端
//*/
////网页
//func (this *GmController) SyncClientMgr() {
//	this.Data["pageTitle"] = "服务器版本客户端"
//	this.display()
//}
//
////AJAX-请求服务器版本客户端列表
//func (this *GmController) AjaxReqSyncClient() {
//
//	mapList := gmconfig.G_ptCommonXmlConfig.GetAllSvrSyncClient()
//	list := make([]map[string]interface{}, 0, len(mapList))
//
//	var wg sync.WaitGroup
//	count := int64(0)
//	for _, ptClient := range mapList {
//		row := make(map[string]interface{})
//		row["id"] = ptClient.Id
//		row["name"] = ptClient.Name
//		row["web_addr"] = ptClient.WebAddr
//		row["desc"] = ptClient.Desc
//		list = append(list, row)
//		//获取版本同步号
//		wg.Add(1)
//		strUrl := fmt.Sprintf("%s/sync_num/?name=%s", ptClient.WebAddr, ptClient.Name)
//		go func(row map[string]interface{}, strUrl string) {
//			defer wg.Done()
//			client := httplib.Get(strUrl).SetTimeout(common.HTTP_CONNECT_TIMEOUT, 5*time.Second)
//			ptRes, err := client.Response()
//			if err != nil {
//				logs.Error("[服务器版本客户端] 发送给同步服错出错，url:%s %s", strUrl, err)
//				return
//			}
//
//			strBody, err := ioutil.ReadAll(ptRes.Body)
//			if err != nil {
//				ptRes.Body.Close()
//				return
//			}
//			row["sync_num"] = string(strBody)
//			ptRes.Body.Close()
//			logs.Info("[服务器版本客户端] 获取版本同步服同步号成功，url:%s 同步号:%s", strUrl, strBody)
//		}(row, strUrl)
//		count++
//	}
//
//	wg.Wait()
//	this.ajaxList("成功", common.MSG_OK, count, list)
//}
//
////AJAX-更新单个
//func (this *GmController) AjaxUpdateSingleServer() {
//	dwId, _ := this.GetUint32("id")
//	ptSyncClient := gmconfig.G_ptCommonXmlConfig.GetSvrSyncClientById(dwId)
//
//	if ptSyncClient == nil {
//		this.ajaxMsg("失败", common.MSG_ERR)
//	}
//	strUrl := fmt.Sprintf("%s/update/?name=%s", ptSyncClient.WebAddr, ptSyncClient.Name)
//	client := httplib.Get(strUrl).SetTimeout(common.HTTP_CONNECT_TIMEOUT, 4*time.Minute)
//	ptRes, err := client.Response()
//	if err != nil {
//		logs.Error("[更新服务器] 发送给同步服错出错，url:%s %s", strUrl, err)
//		this.ajaxMsg("同步服客户端连接不上", common.MSG_OK)
//	}
//	ptRes.Body.Close()
//	logs.Info("[更新服务器] 发送给同步服成功，url:%s ", strUrl)
//	this.ajaxMsg("处理成功", common.MSG_OK)
//}
//
////网页-选择服务器,当前平台下的game，cross，battle混在一起选
//func (this *GmController) Select_Server2() {
//	//this.Data["ResourceList"] = this.getResourceList()
//	this.Data["PlatList"] = this.getPlatList()
//	this.display2()
//}
//
//////ajax-获取玩家二进制数据
//////http://127.0.0.1:8080/gm/getprotouser?plat_id=1&server_id=2001&user_id=20000001000002
////func (this *GmController) GetProtoUser() {
////	dwPlatId, _ := this.GetUint32("plat_id")
////	ptProxyClient := NewProxyClient(common.EProxyClientType_Default, dwPlatId)
////	if ptProxyClient == nil {
////		this.ajaxMsg("平台错误", common.MSG_ERR)
////	}
////
////	strServerId := this.GetString("server_id")
////	strUserId := this.GetString("user_id")
////	if len(strUserId) == 0 {
////		this.ajaxMsg("玩家id为空", common.MSG_ERR)
////	}
////
////	strCmd := "get_proto_user"
////	errMsg, _, data := ptProxyClient.ExecUserGmcmd(strServerId, strUserId, strCmd, "")
////	if len(errMsg) > 0 {
////		this.ajaxMsg(errMsg, common.MSG_ERR)
////	}
////	this.ajaxMsg(data, common.MSG_OK)
////}
//
////---------------一键导号
////导航-网页-界面
//func (this *GmController) OneKeyDump() {
//	this.Data["pageTitle"] = "一键导号"
//
//	ptProxyClient := NewProxyClient(common.EProxyClientType_Default, this.ptCurUser.LastPlatId)
//	if ptProxyClient == nil {
//		this.ajaxMsg("平台错误", common.MSG_ERR)
//	}
//
//	list, _ := ptProxyClient.GetServerList(common.EServerType_Game)
//	this.Data["ServerList"] = list
//	this.Data["ServerId"] = this.ldwServerId
//	this.display()
//}
//
////ajax-导号
////http://30.17.3.155:8080/gm/dumpuser?server_id=2001&dst_plat_id=2&dst_server_id=1&dst_user_id=1000058
//func (this *GmController) DumpUser() {
//	dwPlatId := this.ptCurUser.LastPlatId
//	ptProxyClient := NewProxyClient(common.EProxyClientType_Default, dwPlatId)
//	if ptProxyClient == nil {
//		this.ajaxMsg("平台错误", common.MSG_ERR)
//	}
//
//	strServerId := this.GetString("server_id")
//
//	dwDstPlatId, _ := this.GetUint32("dst_plat_id")
//	ptDstPlatCfg, err := models.GetPlatById(dwDstPlatId)
//	if err != nil {
//		this.ajaxMsg("目标平台找不到", common.MSG_ERR)
//	}
//	strDstGmproxyAddr := ptDstPlatCfg.GmproxyAddr
//
//	strDstServerId := this.GetString("dst_server_id")
//	strDstUserId := this.GetString("dst_user_id")
//	if len(strDstServerId) == 0 {
//		this.ajaxMsg("请输入目标serverid", common.MSG_ERR)
//	}
//	if len(strDstUserId) == 0 {
//		this.ajaxMsg("请输入目标userid", common.MSG_ERR)
//	}
//
//	strCmd := "dump_user"
//	strParam := fmt.Sprintf(`{"gmproxy":"%s","dst_server_id":%s,"dst_user_id":%s}`, strDstGmproxyAddr, strDstServerId, strDstUserId)
//	errMsg, _, data := ptProxyClient.ExecUserGmcmd(strServerId, "", strCmd, strParam)
//	if len(errMsg) > 0 {
//		this.ajaxMsg(errMsg, common.MSG_ERR)
//	}
//	this.ajaxMsg(data, common.MSG_OK)
//}
//
////---------------账号封禁
////导航-网页-界面
//func (this *GmController) ForbidUser() {
//	this.Data["pageTitle"] = "账号封禁"
//
//	ptProxyClient := NewProxyClient(common.EProxyClientType_Default, this.ptCurUser.LastPlatId)
//	if ptProxyClient == nil {
//		this.ajaxMsg("平台错误", common.MSG_ERR)
//	}
//
//	list, _ := ptProxyClient.GetServerList(common.EServerType_Game)
//	this.Data["ServerList"] = list
//	this.Data["ServerId"] = this.ldwServerId
//	this.display()
//}
//
////AJAX-请求某个服封禁玩家列表
//func (this *GmController) ForbidUserList() {
//	dwPlatId := this.ptCurUser.LastPlatId
//	strServerId := this.GetString("server_id")
//	if len(strServerId) == 0 {
//		this.ajaxMsg("请选择服务器id", common.MSG_ERR)
//	}
//	ldwServerId, _ := this.GetUint64("server_id")
//	this.ldwServerId = ldwServerId
//	ptProxyClient := NewProxyClient(common.EProxyClientType_Default, dwPlatId)
//	if ptProxyClient == nil {
//		this.ajaxMsg("平台错误", common.MSG_ERR)
//	}
//
//	dwCnt, list, errMsg := ptProxyClient.ForbidUserList(strServerId, 0, 10)
//	if len(errMsg) > 0 {
//		this.ajaxMsg(errMsg, common.MSG_ERR)
//	}
//
//	this.ajaxList("成功", common.MSG_OK, int64(dwCnt), list)
//}
//
////网页-新加封禁玩家
//func (this *GmController) Add_Forbid_User() {
//	this.Data["PlatList"] = this.getPlatList()
//	this.display2()
//}
//
////ajax-新加封禁玩家
//func (this *GmController) Ajax_Add_Forbid_User() {
//	strServerId := this.GetString("server_id")
//	if len(strServerId) == 0 {
//		this.ajaxMsg("服务器id为空", common.MSG_ERR)
//	}
//
//	strUuid := this.GetString("uuid")
//	strFreeTime := this.GetString("free_time")
//	strReason := this.GetString("reason")
//
//	dwFreeTime := uint32(0) //0表示解封
//	if len(strFreeTime) != 0 {
//		dwFreeTime = utils.Str2UnixTime(strFreeTime)
//	}
//
//	strCmd := "forbid_user"
//	strParam := fmt.Sprintf(`{"uuid":"%s","free_time":%d,"reason":"%s"}`, strUuid, dwFreeTime, strReason)
//	this.doGM(strServerId, "", strCmd, strParam)
//}
//
////让json传更美观
//func FormatJson(strJson string) string {
//	var buf bytes.Buffer
//	json.Indent(&buf, []byte(strJson), "", "\t")
//	return buf.String()
//}
