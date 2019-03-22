package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/wliangde/goleond/spider_web/models"
	_ "github.com/wliangde/goleond/spider_web/routers"
)

func initLog() {
	strCfg := fmt.Sprintf(`{"filename":"%s", "level":%s}`, beego.AppConfig.String("log.filename"), beego.AppConfig.String("log.level"))
	fmt.Println(strCfg)
	logs.SetLogger(logs.AdapterFile, strCfg)
}

//从 conf/app.cnf 读取配置
func initAppConfig() {
	initLog()
}

func main() {
	initAppConfig()
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
