package main

import (
	"fmt"
	"github.com/yuanyangen/houseHelper/http_utils"
	"github.com/yuanyangen/houseHelper/model"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strconv"
	"strings"
)

func main () {
	communityName := "新龙城"
	crawlAllPublishedDataByCommunity(communityName)
}

//根据一个小区获取这个小区的所有正在发布的二手房信息
func crawlAllPublishedDataByCommunity(communityName string)  {
	// get all lianjia info by page
	i := 1
	for {
		body := getRawDataByCommunityName(communityName, i)
		tmp :=strings.Split(body, "data-housecode=\"")
		var oneBatchInfo []*model.HouseInfo
		for _,v := range tmp {
			tt := strings.Split(v, "\"")
			id := tt[0]
			idInt, _ := strconv.ParseInt(id, 10, 64)

			oneBatchInfo = append(oneBatchInfo, parseOneHouseInfo(idInt))
		}

		fmt.Println(string(b))



	}

	// get

}

func getRawDataByCommunityName(communityName string, page int) string {
	url := fmt.Sprintf("https://bj.lianjia.com/ershoufang/pg%vrs%s", page, url2.QueryEscape(communityName))
	return http_utils.Get(url)
}


type RawLianjiHouseInfo struct {
	Data RawLianjiHouseInfoData
}

type RawLianjiHouseInfoData struct {
	Selector deal_property
}

type Selector struct {
	deal_property []KVS
	is_unique []KVS
	buy_house_time []KVS

}

type KVS struct {
	Name string
	Value int64
	selected bool
}

type Params struct {
	city_id int64
	price_listing float64
	is_unique int64
	house_area float64 `json:house_area,string`
	inside_area	 float64 `json:inside_area,string`
	floor int64

}

func parseOneHouseInfo(id int64) *model.HouseInfo {
	url1 := fmt.Sprintf("https://bj.lianjia.com/tools/calccost?house_code=%d", id)
	body := http_utils.Get(url1)



	url := fmt.Sprintf("https://bj.lianjia.com/ershoufang/%d.html?is_sem=1", id)
	resp := http_utils.Get(url)
	fmt.Println(id)
	houseInfo := &model.HouseInfo{}




	return houseInfo
}