package lianjia

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yuanyangen/houseHelper/conf"
	"github.com/yuanyangen/houseHelper/consts"
	"github.com/yuanyangen/houseHelper/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)



func saveErshouInfo(in *model.HouseInfo) {
	date := time.Now().Format("20060102")
	key := fmt.Sprintf("%s_%d",date, in.Id)
	url := fmt.Sprintf("http://%s/%s/%s/%s", conf.GetString("EsHost"), consts.EsIndexNameErshou, consts.EsTypeNameErshou, key)
	bodyB, err :=json.Marshal(in)
	if err != nil {
		return
	}

	req,_ := http.NewRequest("PUT", url, bytes.NewReader(bodyB))
	req.Header.Add("User-Agent", "Mozilla/3.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36")
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("save data to es failed %v", err)
	}
	_, err =  ioutil.ReadAll(resp.Body)
	if err !=nil {
		log.Printf("save data to es failed %v", err)
	}
	return
}



func saveDealInfo(in *model.HouseDealInfo) {
	key := fmt.Sprintf("%d_%d",in.DealTime, in.HouseId)
	url := fmt.Sprintf("http://%s/%s/%s/%s", conf.GetString("EsHost"), consts.EsIndexNameChengjiao, consts.EsTypeNameChengjiao, key)
	bodyB, err :=json.Marshal(in)
	if err != nil {
		return
	}
	log.Printf("start to save deal info %s", string(bodyB))

	req,_ := http.NewRequest("PUT", url, bytes.NewReader(bodyB))
	req.Header.Add("User-Agent", "Mozilla/3.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36")
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("save data to es failed %v", err)
	}
	rb, err :=  ioutil.ReadAll(resp.Body)
	if err !=nil {
		log.Printf("save data to es failed %v %v", string(rb), err)
	}
	return
}

func QueryDealCountByInfo(in *model.HouseDealInfo) int64 {
	sql := fmt.Sprintf("select count(1) from %s where DealTime=%d and HouseId=%d",  consts.EsIndexNameChengjiao, in.DealTime, in.HouseId)

	resp,err := queryWithSql(sql)
	if err != nil {
		log.Printf("query es error: %v", err)
	}
	if strings.Contains(resp, "error") {
		log.Printf("query es error %v", err)
		return 0
	}
	tmp := strings.Split(resp, "\n")
	tmp = tmp[2:]
	c,_ := strconv.ParseInt(strings.TrimSpace(tmp[0]), 10, 64)
	return c
}


func QueryErshouInfoCount(date string) int64 {
	sql := fmt.Sprintf("select count(1) from %s where CrawlDate='%s'",  consts.EsIndexNameErshou, date)

	resp,err := queryWithSql(sql)
	if err != nil {
		log.Printf("query es error %v", err)
		return 0
	}
	if strings.Contains(resp, "error") {
		log.Printf("query es error %v", err)
		return 0
	}
	tmp := strings.Split(resp, "\n")
	tmp = tmp[2:]
	c,_ := strconv.ParseInt(strings.TrimSpace(tmp[0]), 10, 64)
	return c
}


func queryWithSql(sql string) (string, error) {
	url := fmt.Sprintf("http://%s/_sql?format=txt",  conf.GetString("EsHost"))
	body := map[string]string{
		"query": sql,
	}
	bodyB, _ := json.Marshal(body)
	req,_ := http.NewRequest("POST", url, bytes.NewReader(bodyB))
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	rb, err :=  ioutil.ReadAll(resp.Body)
	if err !=nil {
		return "", err
	}
	return string(rb), nil
}


