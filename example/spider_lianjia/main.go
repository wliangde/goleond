/**********************************************
** @Author: liangde.wld
** @Desc:

***********************************************/
package main

import "github.com/wliangde/goleond/spider/spider_lianjia"

func main() {
	spiderXiaoQu()
}

func spiderHouse() {
	ptSpider := spider_lianjia.NewSpiderLianJia()
	if ptSpider == nil {
		return
	}
	ptSpider.Run()
}

func spiderXiaoQu() {
	ptSpider := spider_lianjia.NewLianJiaXiaoQu()
	if ptSpider == nil {
		return
	}
	ptSpider.Run()
}
