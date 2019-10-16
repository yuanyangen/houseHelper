package lianjia

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/yuanyangen/houseHelper/http_utils"
	"github.com/yuanyangen/houseHelper/model"
	"log"
	url2 "net/url"
	"strconv"
	"strings"
	"time"
)

func CrawlAll() {
	communityName := "新龙城"
	CrawlAllPublishedDataByCommunity(communityName)
}

//根据一个小区获取这个小区的所有正在发布的二手房信息
func CrawlAllPublishedDataByCommunity(communityName string) {
	page := 1
	for {
		oneBatchIds := []int64{}

		url := fmt.Sprintf("https://bj.lianjia.com/ershoufang/pg%vrs%s", page, url2.QueryEscape(communityName))
		body := http_utils.Get(url)
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
		if err != nil {
			log.Printf("build dom parser for resp error %v %v", body, err)
			continue
		}

		doc.Find(".item").Each(func(i int, selection *goquery.Selection) {
			houseId,exist := selection.Attr("data-houseid")
			if exist {
				idInt, _ := strconv.ParseInt(houseId, 10, 64)
				oneBatchIds = append(oneBatchIds, idInt)
			}
		})
		if len(oneBatchIds) == 0 {
			log.Printf("all data has has been crawl, finish")
			return
		}

		for _,id:= range oneBatchIds {
			info := CrawlOneHouseInfo(id)
			saveErshouInfo(info)
		}
		page ++
	}
}

func CrawlOneHouseInfo(id int64) *model.HouseInfo {
	houseInfo := model.NewHouseInfo()
	houseInfo.Id = id
	{
		url := fmt.Sprintf("https://bj.lianjia.com/ershoufang/%d.html?is_sem=1", id)
		resp := http_utils.Get(url)
		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp))
		if err != nil {
			log.Fatal(err)
		}

		// 解析房屋的基本属性
		rawLianjiaData := map[string]string{}
		doc.Find(".introContent").Each(func(i int, s *goquery.Selection) {
			// 解析基本信息
			baseSelect := s.Find(".base").Find(".content").First().Children().First().Children()
			baseSelect.Each(func(i int, selection *goquery.Selection) {
				attrName := selection.Children().First().Text()
				attrValue := strings.Replace(selection.Text(), attrName, "", -1)
				rawLianjiaData[attrName] = attrValue
			})
			// 解析交易信息
			tradeSelect := s.Find(".transaction").Find(".content").First().Children().First().Children()
			tradeSelect.Each(func(i int, selection *goquery.Selection) {
				attrName := selection.Children().First().Text()
				attrValue := strings.Replace(selection.Text(), attrName, "", -1)
				rawLianjiaData[attrName] = attrValue
			})
		})

		//关注信息
		doc.Find("#favCount").Each(func(i int, selection *goquery.Selection) {
			houseInfo.FavCount, _ = strconv.ParseInt(selection.Text(), 10, 64)
		})
		doc.Find("#cartCount").Each(func(i int, selection *goquery.Selection) {
			houseInfo.TotalSeeCount, _ = strconv.ParseInt(selection.Text(), 10, 64)
		})

		houseInfo.CommunityName = doc.Find(".communityName").Find(".info").Text()


		for k, v := range rawLianjiaData {
			switch k {
			case "供暖方式":
				{
					if v == "自供暖" {
						houseInfo.WarnSupplyType = model.WarnSupplyTypeSelf
					} else {
						houseInfo.WarnSupplyType = model.WarnSupplyTypeCentral
					}
				}
			case "装修情况":
				{
					var decorateValueMap = map[string]int{
						"精装": model.DecorateTypeGood,
						"简装": model.DecorateTypeNormal,
						"毛坯": model.DecorateTypeNone,
					}
					houseInfo.DecorateType = decorateValueMap[v]
				}
			case "房屋年限":
				{
					var decorateValueMap = map[string]int{
						"满两年": model.YearTypeGreatThanTwo,
						"满五年": model.DecorateTypeNormal,
					}
					v = strings.TrimSpace(v)
					houseInfo.TradeYearType = decorateValueMap[v]

				}
			case "产权所属":
				{
					var decorateValueMap = map[string]model.BelongType{
						"共有":  model.Public,
						"非共有": model.Private,
					}
					v = strings.TrimSpace(v)
					houseInfo.BelongTo = decorateValueMap[v]
				}
			case "所在楼层":
				{
					houseInfo.FloorNumber = strings.TrimSpace(v)
				}
			case "建筑面积":
				{
					v = strings.TrimSpace(v)
					v = strings.TrimSuffix(v, "㎡")
					houseInfo.BuildingSize, _ = strconv.ParseFloat(v, 10)
				}
			case "套内面积":
				{
					v = strings.TrimSpace(v)
					v = strings.TrimSuffix(v, "㎡")
					houseInfo.RealSize, _ = strconv.ParseFloat(v, 10)
				}
			case "户型结构":
				{
					houseInfo.StandardType = strings.TrimSpace(v)
				}
			case "建筑结构":
				{
					houseInfo.BuildingStruct = strings.TrimSpace(v)
				}
			case "梯户比例":
				{
					houseInfo.FloorToUseCount = strings.TrimSpace(v)
				}
			case "配备电梯":
				{
					if v == "有" {
						houseInfo.Elevator = true
					}
				}
			case "产权年限":
				{
					v = strings.TrimSpace(v)
					v = strings.TrimSuffix(v, "年")
					houseInfo.OwnYearCountType, _ = strconv.Atoi(v)
				}
			case "挂牌时间":
				{
					v = strings.TrimSpace(v)
					t,err := time.Parse("2006-01-02", v)
					if err == nil {
						houseInfo.OnlineTime = t.Unix()
					}
				}
			case "上次交易":
				{
					v = strings.TrimSpace(v)
					t,err := time.Parse("2006-01-02", v)
					if err == nil {
						houseInfo.LastTradeTime = t.Unix()
					}
				}
			case "交易权属": {
				v = strings.TrimSpace(v)
				houseInfo.TradeType = v
			}
			case "抵押信息": {
				if strings.Contains(v, "有") {
					houseInfo.Mortgage = true
				}
			}
			case "房屋户型": {
				v = strings.TrimSpace(v)
				houseInfo.Type = v
			}
			case "建筑类型": {
				v = strings.TrimSpace(v)
				houseInfo.BuildingType = v
			}
			case "房屋朝向": {
				v = strings.TrimSpace(v)
				houseInfo.ForwardType = v
			}
			case "房屋用途": {
				v = strings.TrimSpace(v)
				houseInfo.HouseUsageType = v
			}

			}
		}
	}


	//关注信息
	//{
	//	url := fmt.Sprintf("https://bj.lianjia.com/ershoufang/houseseerecord?id=%d", id)
	//	resp := http_utils.Get(url)
	//	seeRecord := SeeRecord{}
	//	err := json.Unmarshal([]byte(resp), &seeRecord)
	//	if err != nil {
	//		log.Printf("get see record error data %s, err:%v", resp, err.Error())
	//	} else {
	//		houseInfo.WeekSeeCount = seeRecord.Data.ThisWeek
	//		houseInfo.MonthSeeCount = seeRecord.Data.TotalCnt
	//	}
	//}
	houseInfo.ModifyTime = time.Now().Unix()
	houseInfo.CrawlDate = time.Now().Format("20060102")

	return houseInfo
}

type SeeRecord struct {
	Data struct {
		ThisWeek int64 `json:"thisWeek"`
		TotalCnt int64 `json:"totalCnt"`
	}
}
