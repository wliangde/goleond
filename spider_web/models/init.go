/**********************************************
** @Des:
** @Author: wangliangde
** @Date:   2018/4/3 15:46
***********************************************/

package models

import (
	//"net/url"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
)

func Init() {
	initDb()
}

func initDb() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	err := orm.RegisterDataBase("default", "mysql", dsn)
	if err != nil {
		logs.Error("数据库注册失败，%s", dsn, err)
		return
	}
	orm.RegisterModel(new(BackendUser),
		new(Auth),
		new(Group),
		new(GroupAuth),
		new(Log),
	)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	logs.Info("数据库初始化完成")
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}
