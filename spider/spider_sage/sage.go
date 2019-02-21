package spider_sage

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type TTorrent struct {
	Url         string
	DownLoadCnt uint32
}

type SageStory struct {
	Id          uint64
	Url         string
	Name        string
	Size        string
	Ma          string
	Pic         string
	TorUrl      string //种子地址
	DownloadCnt uint32 //下载次数
}

type TSpiderSage struct {
	strBaseUrl string
	ChanStory  chan *SageStory
	wg         sync.WaitGroup
}

func NewSpiderSage() *TSpiderSage {
	ptSpider := &TSpiderSage{
		strBaseUrl: "https://www.kongjiee.space/",
		ChanStory:  make(chan *SageStory, 10000),
	}
	if ptSpider.initDb() == false {
		return nil
	}
	return ptSpider
}

func (this *TSpiderSage) initDb() bool {
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
	orm.RegisterModel(new(SageStory))
	logs.Info("数据库连接成功", dsn)
	return true
}

func (this *TSpiderSage) Run() {
	this.wg.Add(1)
	go this.InsertDb()

	this.ParseWeb(1, 1001)
	this.wg.Wait()
}

func (this *TSpiderSage) ParseWeb(nFrom, nTo int) {
	ptBaseC := colly.NewCollector()
	ptC2 := ptBaseC.Clone()
	//#normalthread_11544827
	dwCnt := uint32(0)
	ptBaseC.OnHTML("tbody[id^=normalthread_]  a.xst", func(ptEle *colly.HTMLElement) {
		strUrl := ptEle.Attr("href")
		strNewUrl := fmt.Sprintf("https://www.kongjiee.space/%s", strUrl)
		//if dwCnt >= 1 {
		//	return
		//}
		//logs.Info("%s %s ", strText, strNewUrl)
		this.ParseSingle(strNewUrl, ptC2)
		dwCnt++
	})

	for i := nFrom; i < nTo; i++ {
		strUrl := fmt.Sprintf("https://www.kongjiee.space/forum-798-%d.html", i)
		ptBaseC.Visit(strUrl)
		//break
	}

	close(this.ChanStory)
}

//分析详细页面
func (this *TSpiderSage) ParseSingle(strUrl string, ptC *colly.Collector) {
	//strUrl = "https://www.kongjiee.space/thread-11544782-1-1.html"

	ptStory := &SageStory{}
	ptStory.Url = strUrl

	//抓取标题和size 预览图片
	ptC.OnHTML("td[id^=postmessage_]", func(ptEle *colly.HTMLElement) {
		strText := strings.TrimSpace(ptEle.Text)
		if !strings.Contains(strText, "【影片名称】") {
			return
		}
		//logs.Info("wldddd %s", strText)
		slcStr := strings.Split(strText, "\n")
		for _, str := range slcStr {
			str = strings.Trim(str, "\n")
			str = strings.TrimSpace(str)
			if len(str) == 0 {
				continue
			}
			if strings.HasPrefix(str, "【影片名称】：") {
				ptStory.Name = strings.Trim(str, "【影片名称】：")
			}
			if strings.HasPrefix(str, "【影片大小】：") {
				ptStory.Size = strings.Trim(str, "【影片大小】：")
			}
			if strings.HasPrefix(str, "【是否有码】：") {
				ptStory.Ma = strings.Trim(str, "【是否有码】：")
			}
		}

		//预览图片
		slcPic := make([]string, 0, 0)
		ptEle.DOM.Find("img[id^=aimg_]").Each(func(i int, selection *goquery.Selection) {
			str, _ := selection.Attr("file")
			slcPic = append(slcPic, str)
		})
		//不存了
		//ptStory.Pic = strings.Join(slcPic, ";")
		//logs.Info("name: %s size:%s %s ", ptStory.Name, ptStory.Size, ptStory.Ma)
	})

	//抓取种子地址和下载次数
	bFind := false
	ptC.OnHTML("dl.tattl >dd ", func(ptEle *colly.HTMLElement) {
		if bFind {
			return
		}
		//logs.Info("%s", ptEle.Text)
		//可能存在一个页面多个种子的情况
		//ptTmpStory := &SageStory{
		//	Url:  ptStory.Url,
		//	Name: ptStory.Name,
		//	Ma:   ptStory.Ma,
		//	Size: ptStory.Size,
		//	Pic:  ptStory.Pic,
		//}
		ptSel := ptEle.DOM.Find("p.attnm > a")
		if ptSel != nil {
			strTorrentUrl, _ := ptSel.Attr("href")
			ptStory.TorUrl = fmt.Sprintf("https://www.kongjiee.space/%s", strTorrentUrl)
		}
		ptSel = ptEle.DOM.Find("p:contains('下载次数')")
		if ptSel != nil {
			reg := regexp.MustCompile(`下载次数: ([\d]+)`)
			slcStr := reg.FindStringSubmatch(strings.TrimSpace(ptSel.Text()))
			if len(slcStr) > 1 {
				a, _ := strconv.ParseUint(slcStr[1], 10, 64)
				ptStory.DownloadCnt = uint32(a)
			}
		}
		this.ChanStory <- ptStory
		logs.Info("%v", ptStory)
		bFind = true
	})

	ptC.Visit(strUrl)

}

func (this *TSpiderSage) InsertDb() {
	defer this.wg.Done()

	for ptStory := range this.ChanStory {
		_, err := orm.NewOrm().Insert(ptStory)
		if err != nil {
			logs.Error("插入数据库失败 %s %s", ptStory.Url, err)
		} else {
			logs.Info("插入数据成功 %s %s ", ptStory.Url, ptStory.Name)
		}
	}
}
