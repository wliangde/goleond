/**
@user:          liangde.wld
@createtime:    2019/3/18 9:39 PM
@desc:
**/
package spider_lianjia

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/gocolly/colly"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

type XiaoQu struct {
	Id       uint64
	XiaoQuId uint64
	Name     string
	Price    uint32
	SellCnt  uint32
	SoldCnt  uint32 //
	Url      string
	Area     string //所属区域
}

type TLianJiaXiaoQu struct {
	strBaserUrl string
	mapXiaoQu   map[uint64]*XiaoQu
	chanXiaoQu  chan *XiaoQu
}

func NewLianJiaXiaoQu() *TLianJiaXiaoQu {
	ptSpider := &TLianJiaXiaoQu{
		strBaserUrl: "https://sh.lianjia.com/xiaoqu/",
		mapXiaoQu:   make(map[uint64]*XiaoQu),
		chanXiaoQu:  make(chan *XiaoQu, 1000),
	}

	if ptSpider.initDb() == false {
		return nil
	}
	return ptSpider
}

func (this *TLianJiaXiaoQu) initDb() bool {
	dbhost := "47.110.50.49"
	dbport := "3306"
	dbuser := "root"
	dbpassword := "10086"
	dbname := "lianjia"
	timezone := ""
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	err := orm.RegisterDataBase("default", "mysql", dsn)
	if err != nil {
		logs.Error("数据库注册失败, %s error:%s", dsn, err)
		return false
	}
	orm.RegisterModel(new(XiaoQu))
	logs.Info("数据库连接成功", dsn)
	return true
}

func (this *TLianJiaXiaoQu) Run() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		go this.ParseWeb()
	}()

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		this.DbInsert()
	}()
	wg.Wait()
}

func (this *TLianJiaXiaoQu) ParseWeb() {
	ptBaseC := colly.NewCollector()
	ptAreaC := ptBaseC.Clone()
	ptSmallAreaC := ptBaseC.Clone()
	ptSingePageC := ptBaseC.Clone()

	slcSmallArea := make([]string, 0, 100)
	slcFullUrl := make([]string, 0, 1000)

	//获取大区域
	ptBaseC.OnHTML("div[data-role='ershoufang']>div > a", func(ptEle *colly.HTMLElement) {
		strHref := ptEle.Attr("href")
		if strHref == "/xiaoqu/shanghaizhoubian/" {
			return
		}
		strNewUrl := ptEle.Request.AbsoluteURL(strHref)
		fmt.Println(strNewUrl)
		ptAreaC.Visit(strNewUrl)
	})

	//获取小区域
	ptAreaC.OnHTML("div[data-role='ershoufang']>div:nth-child(2) > a", func(ptEle *colly.HTMLElement) {
		strHref := ptEle.Attr("href")
		strNewUrl := ptEle.Request.AbsoluteURL(strHref)
		//fmt.Println("wld", strNewUrl)
		slcSmallArea = append(slcSmallArea, strNewUrl)
	})

	//分页
	ptSmallAreaC.OnHTML("div.contentBottom > div.page-box > div.house-lst-page-box", func(ptEle *colly.HTMLElement) {
		//fmt.Println("wld", ptEle.Request.URL, ptEle.Text)
		//分页
		strAttr := ptEle.Attr("page-data") //{"totalPage":27,"curPage":1}
		dwTotalPage := 0
		dwCurPage := 0
		fmt.Fscanf(strings.NewReader(strAttr), "{\"totalPage\":%d,\"curPage\":%d}", &dwTotalPage, &dwCurPage)
		//fmt.Println(strAttr, dwTotalPage, dwCurPage)
		logs.Info("url:%s totalpage:%d", ptEle.Request.URL, dwTotalPage)
		if dwTotalPage == 0 {
			logs.Error("url:%s no page ", ptEle.Request.URL)
			return
		}
		for i := 1; i < dwTotalPage; i++ {
			strUrl := fmt.Sprintf("%spg%d", ptEle.Request.URL, i)
			slcFullUrl = append(slcFullUrl, strUrl)
		}
	})
	ptBaseC.Visit(this.strBaserUrl)

	//去分页
	for _, strUrl := range slcSmallArea {
		ptSmallAreaC.Visit(strUrl)
		//break
	}
	//单页信息分析
	ptSingePageC.OnHTML("li[class='clear xiaoquListItem']", func(ptEle *colly.HTMLElement) {
		//fmt.Println(ptEle.Request.URL, ptEle.Text)
		ptXiaoQu := &XiaoQu{}
		//标题
		ptSelection := ptEle.DOM.Find("div.title > a")
		if ptSelection != nil {
			ptXiaoQu.Url, _ = ptSelection.Attr("href") //https://sh.lianjia.com/ershoufang/107100530553.html
			fmt.Fscanf(strings.NewReader(ptXiaoQu.Url), "https://sh.lianjia.com/xiaoqu/%d.html", &ptXiaoQu.XiaoQuId)
			ptXiaoQu.Name = ptSelection.Text()
		}

		//区域		浦东 北蔡
		ptSel := ptEle.DOM.Find("div.positionInfo > a.bizcircle")
		if ptSel != nil {
			ptXiaoQu.Area = ptSel.Text()
		}
		//90天成交量
		ptSel = ptEle.DOM.Find("div.houseInfo > a")
		if ptSel != nil {
			strTitle, _ := ptSel.Attr("title") //xxx网签
			if strings.HasSuffix(strTitle, "网签") {
				str := ptSel.Text() //90天成就1套
				fmt.Fscanf(strings.NewReader(str), "90天成交%d套", &ptXiaoQu.SoldCnt)
			}
		}
		//均价
		ptSel = ptEle.DOM.Find("div.totalPrice > span")
		if ptSel != nil {
			lnPrice, _ := strconv.ParseInt(ptSel.Text(), 10, 64)
			ptXiaoQu.Price = uint32(lnPrice)
		}

		//在售套数
		ptSel = ptEle.DOM.Find("a.totalSellCount > span")
		if ptSel != nil {
			lnSell, _ := strconv.ParseInt(ptSel.Text(), 10, 64)
			ptXiaoQu.SellCnt = uint32(lnSell)
		}

		this.chanXiaoQu <- ptXiaoQu
		logs.Debug("wld======", ptXiaoQu)
	})

	//分页后的数据
	for _, strUrl := range slcFullUrl {
		ptSingePageC.Visit(strUrl)
		//break
	}

	close(this.chanXiaoQu)
}

func (this *TLianJiaXiaoQu) DbInsert() {
	for ptHouse := range this.chanXiaoQu {
		_, err := orm.NewOrm().Insert(ptHouse)
		if err != nil {
			logs.Error("数据库插入失败, houseurl:%s error:%s", ptHouse.Url, err)
		} else {
			logs.Info("数据库插入成功, houseurl:%s", ptHouse.Url)

		}
	}
}
