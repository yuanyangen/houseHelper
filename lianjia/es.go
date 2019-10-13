package lianjia

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yuanyangen/houseHelper/conf"
	"github.com/yuanyangen/houseHelper/consts"
	"github.com/yuanyangen/houseHelper/model"
	"io/ioutil"
	"net/http"
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
		fmt.Println(err)
	}
	_, err =  ioutil.ReadAll(resp.Body)
	if err !=nil {
		fmt.Println(err)
	}
	return
}



func saveDealInfo(in *model.HouseDealInfo) {
	date := time.Now().Format("20060102")
	key := fmt.Sprintf("%s_%d",date, in.HouseId)
	url := fmt.Sprintf("http://%s/%s/%s/%s", conf.GetString("EsHost"), consts.EsIndexNameChengjiao, consts.EsTypeNameChengjiao, key)
	bodyB, err :=json.Marshal(in)
	if err != nil {
		return
	}

	req,_ := http.NewRequest("PUT", url, bytes.NewReader(bodyB))
	req.Header.Add("User-Agent", "Mozilla/3.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36")
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	rb, err :=  ioutil.ReadAll(resp.Body)
	if err !=nil {
		fmt.Println(err, string(rb))
	}
	return
}


