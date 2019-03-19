/**********************************************
** @Author: liangde.wld
** @Desc:

***********************************************/
package main

import (
	"flag"
	"github.com/wliangde/goleond/spider/spider_lianjia"
)

var gbXiaoQu bool

func main() {
	flag.BoolVar(&gbXiaoQu, "xq", true, "是否是爬取小区")

	flag.Parse()
	
if gbXiaoQu {
		spiderXiaoQu()
	} else {
		spiderHouse()
	}
}

//爬取房子
func spiderHouse() {
	ptSpider := spider_lianjia.NewSpiderLianJia()
	if ptSpider == nil {
		return
	}
	ptSpider.Run()
}

//爬取小区
func spiderXiaoQu() {
	ptSpider := spider_lianjia.NewLianJiaXiaoQu()
	if ptSpider == nil {
		return
	}
	ptSpider.Run()
}
