/**********************************************
** @Des:
** @Author: wangliangde
** @Date:   2018/4/3 18:35
***********************************************/

package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Index() {
	this.Data["pageTitle"] = "系统首页"
	//self.display()
	this.TplName = "public/main.html"
}

func (this *HomeController) Start() {
	this.Data["pageTitle"] = "控制面板"

	o := orm.NewOrm()
	dwTotalCount := uint32(0)
	fAvgPrice := float32(0)
	fAvgTotalPrice := float32(0)
	err := o.Raw("select count(*) as total_count,avg(price) as avg_price, avg(total_price) as avg_total_price from house").QueryRow(&dwTotalCount, &fAvgPrice, &fAvgTotalPrice)
	if err != nil {
		logs.Error(err)
	}
	this.Data["total_count"] = dwTotalCount       //总挂牌数
	this.Data["avg_price"] = fAvgPrice            //均价
	this.Data["avg_total_price"] = fAvgTotalPrice //均总价

	t := time.Now()
	//mmdd
	strTime := t.Format("0102")

	logs.Info("ddddd", strTime)

	this.display()
}

//选择渠道
func (this *HomeController) ChoosePlatId() {
	dwPlatId, _ := this.GetUint32("plat_id")
	//logs.Info("选择平台id %d", dwPlatId)
	this.ptCurUser.LastPlatId = dwPlatId
	this.setUser2Session(this.ptCurUser)
	this.redirect(beego.URLFor("HomeController.Index"))
}
