/**********************************************
** @Des:
** @Author: wangliangde
** @Date:   2018/4/4 11:37
***********************************************/

package common

import (
	"time"
)

var (
	TBNAME_BACKEND_USER   string = "uc_backend_user"
	TBNAME_AUTH           string = "uc_auth"
	TBNAME_GROUP          string = "uc_group"
	TBNAME_GROUP_AUTH     string = "uc_group_auth"
	TBNAME_LOG            string = "log"
	TBNAME_CHANNEL        string = "channel"
	TBNAME_VERSION_SVR    string = "version_svr"
	TBNAME_PLAT           string = "plat"
	TBNAME_SERVER_VERSION string = "server_version"
)

var (
	SESSION_USER  string = "sesuser"
	GROUP_SUPERGM int    = 1
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

var (
	CONFIG_PATH_GAME = "../../config/game_config.xml"
)

type ELogAction string

var (
	//ELogActionLogin         ELogAction = 1
	ELogActionGMAny            ELogAction = "任意GM"
	ELogActionUpdateCfg        ELogAction = "更新配置"
	ELogActionMail             ELogAction = "发送邮件"
	ELogActionOneKey           ELogAction = "一键养成"
	ELogActionPublishVersion   ELogAction = "发布版本"
	ELogActionRePublishVersion ELogAction = "重新发布版本"
	ELogActionStopServer       ELogAction = "关闭服务器"
	ELogActionRestartServer    ELogAction = "重启服务器"
)

const (
	EServerType_Min    uint32 = 0
	EServerType_Game   uint32 = 1
	EServerType_Cross  uint32 = 2
	EServerType_Battle uint32 = 3
	EServerType_Gate   uint32 = 4
	EServerType_Max    uint32 = 5
)

const (
	HTTP_CONNECT_TIMEOUT    = 10 * time.Second
	HTTP_READ_WRITE_TIMEOUT = 30 * time.Second
)

type EProxyClientType uint32

const (
	EProxyClientType_Default     EProxyClientType = 0 //无要求
	EProxyClientType_MultiServer EProxyClientType = 1 //多服无玩家
	EProxyClientType_MultiUser   EProxyClientType = 2 //多服多玩家
)

const (
	VERSION_SERVER_SALT = "#@!#@DW"                 //发布版本加密盐
	CLIENT_VERSION_FILE = "data/version.xml"        //版本文件路径
	SERVER_VERSION_FILE = "data/server_version.xml" //服务器版本文件
)
const (
	PRIVATEKEY = "!#$#@12da"
)

const (
	NO_AUTH = "/loginin/dologin/loginout/home/start/getresourceidlist/ajaxsave/ajaxdel/table/getnodes/add/edit/ajaxnew/pwdsave/ajaxaddnewuser/" +
		"ajaxedituser/addrecv/searchuserbrief/addattach/" +
		"selectserver/platserver/sendsysmail/sendusermail/doonekeydev/" +
		"gm/uploadfile/gm/releaseversionselect/gm/platchannel/gm/dorelease/gm/reqversionlist/" +
		"gm/getbriefseverlist/gm/ajaxstopserver/gm/ajaxrestartserver/select_game/ajaxreqsvrversion/ajaxreleasesvrver/ajaxupdateserver/ajaxupdateservercfg" +
		"ajaxreqsyncclient/ajaxupdatesingleserver/select_server2/dumpuser/forbiduserlist/addforbiduser/add_forbid_user/ajax_add_forbid_user" //不需要认证的请求
)
