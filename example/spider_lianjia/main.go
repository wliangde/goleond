/**********************************************
** @Author: liangde.wld
** @Desc:

***********************************************/
package main

import "github.com/wliangde/goleond/spider/spider_lianjia"

func main() {
	ptSpider := spider_lianjia.NewSpiderLianJia()
	if ptSpider == nil {
		return
	}
	ptSpider.Run()
}
