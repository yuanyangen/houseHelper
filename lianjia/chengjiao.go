package lianjia

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/yuanyangen/houseHelper/http_utils"
	"github.com/yuanyangen/houseHelper/model"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func CrawlAllChengjiao() {
	communityName := "新龙城"
	CrawlAllChengjiaoByCommunity(communityName)
}

//根据一个小区获取这个小区的所有正在发布的二手房信息
func CrawlAllChengjiaoByCommunity(communityName string) {
	page := 1
	for {
		u := fmt.Sprintf("https://bj.lianjia.com/chengjiao/pg%vrs%s", page, url.QueryEscape(communityName))
		body := http_utils.Get(u)
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
		if err != nil {
			log.Fatal(err)
		}

		oneBatchInfo := make([]*model.HouseDealInfo, 0)
		doc.Find(".listContent").Children().Each(func(i int, selection *goquery.Selection) {
			title := selection.Find(".title").Children().First()
			jmpUrl,_ := title.Attr("href")
			if jmpUrl == ""{
				return
			}
			t := strings.Split(jmpUrl, "/")
			t = strings.Split(t[len(t)-1], ".")
			id,_ := strconv.ParseInt(t[0], 10, 64)
			dealInfo := &model.HouseDealInfo{}
			dealInfo.HouseId = id

			dealDateS := selection.Find(".dealDate").Text()
			dealTime, _ := time.Parse("2006.01.02", dealDateS)

			dealInfo.DealTime = dealTime.Unix()

			dealPrice := selection.Find(".totalPrice").Find(".number").Text()

			dealInfo.DealPrice, _ = strconv.ParseFloat(dealPrice, 10)

			selection.Find(".dealCycleTxt").Children().Each(func(i int, selection *goquery.Selection) {
				if i == 0 {
					rawPriceS := selection.Text()
					rawPriceS = strings.Replace(rawPriceS, "挂牌", "", -1)
					rawPriceS = strings.Replace(rawPriceS, "万", "", -1)
					dealInfo.RawPrice,_ = strconv.ParseFloat(rawPriceS, 10)
				} else if i == 1 {
					rawPriceS := selection.Text()
					rawPriceS = strings.Replace(rawPriceS, "成交周期", "", -1)
					rawPriceS = strings.Replace(rawPriceS, "天", "", -1)
					dealInfo.Duration,_ = strconv.ParseInt(rawPriceS, 10, 64)
				}
			})
			dealInfo.ModifyTime = time.Now().Unix()
			oneBatchInfo = append(oneBatchInfo, dealInfo)
		})

		for _,oneDeal := range oneBatchInfo {
			saveDealInfo(oneDeal)
		}

		page ++
	}
}
