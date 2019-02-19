package spider_sage

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gocolly/colly"
	"regexp"
	"strconv"
	"strings"
)

type TTorrent struct {
	Url         string
	DownLoadCnt uint32
}

type TStory struct {
	Name       string
	Size       string
	Ma         string
	Url        string
	SlcTorrent []*TTorrent
}

type TSpiderSage struct {
	strBaseUrl string
}

func NewSpiderSage() *TSpiderSage {
	ptSpider := &TSpiderSage{
		strBaseUrl: "https://www.kongjiee.space/",
	}
	return ptSpider
}

func (this *TSpiderSage) Run() {
	ptBaseC := colly.NewCollector()
	//#normalthread_11544827
	dwCnt := uint32(0)
	ptBaseC.OnHTML("tbody[id^=normalthread_]  a.xst", func(ptEle *colly.HTMLElement) {
		strText := ptEle.Text
		strUrl := ptEle.Attr("href")
		strNewUrl := fmt.Sprintf("https://www.kongjiee.space/%s", strUrl)
		if dwCnt >= 1 {
			return
		}
		logs.Info("%s %s ", strText, strNewUrl)
		this.ParseSingle(strNewUrl, ptBaseC.Clone())
		dwCnt++
	})

	for i := 1; i <= 1000; i++ {
		strUrl := fmt.Sprintf("https://www.kongjiee.space/forum-798-%d.html", i)
		ptBaseC.Visit(strUrl)
		break
	}
}

//分析详细页面
func (this *TSpiderSage) ParseSingle(strUrl string, ptC *colly.Collector) {
	strUrl = "https://www.kongjiee.space/thread-11544782-1-1.html"

	ptStory := &TStory{}
	ptStory.Url = strUrl
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

		logs.Info("name: %s size:%s %s ", ptStory.Name, ptStory.Size, ptStory.Ma)

	})

	ptC.OnHTML("dl.tattl >dd ", func(ptEle *colly.HTMLElement) {
		//logs.Info("%s", ptEle.Text)
		ptTorrent := &TTorrent{}
		ptSel := ptEle.DOM.Find("p.attnm > a")
		if ptSel != nil {
			strTorrentUrl, _ := ptSel.Attr("href")
			ptTorrent.Url = fmt.Sprintf("https://www.kongjiee.space/%s", strTorrentUrl)
		}
		ptSel = ptEle.DOM.Find("p:contains('下载次数')")
		if ptSel != nil {
			reg := regexp.MustCompile(`下载次数: ([\d]+)`)
			slcStr := reg.FindStringSubmatch(strings.TrimSpace(ptSel.Text()))
			if len(slcStr) > 0 {
				a, _ := strconv.ParseUint(slcStr[0], 10, 64)
				ptTorrent.DownLoadCnt = uint32(a)
			}
		}
		logs.Info("torrent %v", ptTorrent)
		ptStory.SlcTorrent = append(ptStory.SlcTorrent, ptTorrent)
	})

	ptC.Visit(strUrl)
}
