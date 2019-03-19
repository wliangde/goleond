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
	orm.RegisterModel(new(House))
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
	//strUrl := "https://sh.lianjia.com/ershoufang/"
	ptBaseC := colly.NewCollector()
	ptAreaC := ptBaseC.Clone()
	ptSmallAreaC := ptBaseC.Clone()
	ptSingePageC := ptBaseC.Clone()

	slcSmallArea := make([]string, 0, 100)
	slcFullUrl := make([]string, 0, 1000)

	dwCount := uint32(0)
	dwCount2 := uint32(0)

	//获取大区域
	ptBaseC.OnHTML("div[data-role='ershoufang']>div > a", func(ptEle *colly.HTMLElement) {
		if dwCount >= 1 {
			return
		}
		strHref := ptEle.Attr("href")
		if strHref == "/xiaoqu/shanghaizhoubian/" {
			return
		}
		strNewUrl := ptEle.Request.AbsoluteURL(strHref)
		fmt.Println(strNewUrl)
		ptAreaC.Visit(strNewUrl)
		dwCount++

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
		break
	}
	dwCount2++

	//slcFullUrl = append(slcFullUrl, "https://sh.lianjia.com/ershoufang/xidu/pg9/")

	//单页信息分析
	//ptSingePageC.OnHTML("div[class='info clear']", func(ptEle *colly.HTMLElement) {
	//	//fmt.Println(ptEle.Request.URL, ptEle.Text)
	//	ptHouse := &House{}
	//	//标题
	//	ptSelection := ptEle.DOM.Find("div.title > a") //<div class="title"><a class="" href="https://sh.lianjia.com/ershoufang/107001263925.html" target="_blank" data-log_index="1" data-el="ershoufang" data-housecode="107001263925" data-is_focus="1" data-sl="">鹏海东苑，两房朝南，客厅厨卫朝北有窗，光线充足</a><!-- 拆分标签 --><span class="yezhushuo tagBlock">房主自荐</span></div>
	//	if ptSelection != nil {
	//		ptHouse.Url, _ = ptSelection.Attr("href") //https://sh.lianjia.com/ershoufang/107100530553.html
	//		fmt.Fscanf(strings.NewReader(ptHouse.Url), "https://sh.lianjia.com/ershoufang/%d.html", &ptHouse.HouseId)
	//		ptHouse.Title = ptSelection.Text()
	//	}
	//	//房型信息
	//	ptSelection = ptEle.DOM.Find("div.houseInfo")
	//	if ptSelection != nil {
	//		ptHouse.HouseInfo = ptSelection.Text()
	//		//小区
	//		ptSel2 := ptSelection.ChildrenFiltered("a")
	//		if ptSel2 != nil {
	//			ptHouse.XiaoQuName = ptSel2.Text()
	//			ptHouse.XiaoQuUrl, _ = ptSel2.Attr("href")
	//			fmt.Fscanf(strings.NewReader(ptHouse.XiaoQuUrl), "https://sh.lianjia.com/xiaoqu/%d/", &ptHouse.XiaoQuId)
	//		}
	//
	//		//紫叶花园东园  | 2室1厅 | 68.65平米 | 暂无数据 | 简装 | 无电梯
	//		slcStr := strings.Split(ptHouse.HouseInfo, "|")
	//		if len(slcStr) >= 3 {
	//			ptHouse.HuXing = strings.TrimSpace(slcStr[1])
	//			fmt.Fscanf(strings.NewReader(strings.TrimSpace(slcStr[2])), "%d平米", &ptHouse.MianJi)
	//		}
	//	}
	//
	//	//楼层信息 高楼层(共6层)1995年建板楼
	//	ptSelection = ptEle.DOM.Find("div.positionInfo")
	//	if ptSelection != nil {
	//		slcString := strings.Split(ptSelection.Text(), "-")
	//		if len(slcString) >= 2 {
	//			ptHouse.Flood = strings.TrimSpace(slcString[0])
	//			ptHouse.Area = strings.TrimSpace(slcString[1])
	//		}
	//	}
	//
	//	//带看信息
	//	ptSelection = ptEle.DOM.Find("div.followInfo")
	//	if ptSelection != nil {
	//		strFollowInfo := ptSelection.Text()
	//		slcStr := strings.Split(strFollowInfo, "/") //245人关注 / 共7次带看 / 一年前发布
	//		if len(slcStr) == 3 {
	//			fmt.Fscanf(strings.NewReader(strings.TrimSpace(slcStr[0])), "%d人关注", &ptHouse.Follow)
	//			fmt.Fscanf(strings.NewReader(strings.TrimSpace(slcStr[1])), "共%d次带看", &ptHouse.Look)
	//			ptHouse.ReleaseTime = strings.TrimSpace(slcStr[2])
	//			//log.Debug("followinfo %d, %d, %s", ptHouse.Follow, ptHouse.Look, ptHouse.ReleaseTime)
	//		}
	//	}
	//
	//	//tag
	//	ptSelection = ptEle.DOM.Find("div.tag")
	//	if ptSelection != nil {
	//		ptHouse.Tag = ptSelection.Text()
	//	}
	//
	//	//总价
	//	ptSelection = ptEle.DOM.Find("div.totalPrice>span")
	//	if ptSelection != nil {
	//		f, _ := strconv.ParseFloat(ptSelection.Text(), 64)
	//		ptHouse.TotalPrice = uint32(f)
	//	}
	//	//单价
	//	ptSelection = ptEle.DOM.Find("div.unitPrice")
	//	if ptSelection != nil {
	//		strPrice, _ := ptSelection.Attr("data-price")
	//		dwI, _ := strconv.ParseInt(strPrice, 10, 64)
	//		ptHouse.Price = uint32(dwI)
	//	}
	//	this.chanXiaoQu <- ptHouse
	//	logs.Debug("wld======", ptHouse)
	//})

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