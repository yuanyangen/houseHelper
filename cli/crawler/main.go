package main

import (
	"github.com/yuanyangen/houseHelper/lianjia"
	"log"
	"time"
)

var communityNames = []string {
		"新龙城",
		"流星花园",
		"龙泽苑",
		"金榜园",
		"和谐家园",
		"龙腾苑",
		"龙跃苑",
	}

func main () {
	for {
		go oneRound()
		time.Sleep(time.Hour * 4)
	}
}



func oneRound() {
	log.Printf("start one round\n")
	crawlpublish()
	crawlChengjiao()

}

func crawlChengjiao() {
	for _, v := range communityNames {
		lianjia.CrawlAllChengjiaoByCommunity(v)
	}
}

func crawlpublish() {
	date := time.Now().Format("20060102")
	log.Printf("start chengjiao of %s \n", date)

	count:= lianjia.QueryErshouInfoCount(date)

	if count >= 0 {
		log.Printf("already has data , ignore\n")
		return
	}

	for _, v := range communityNames {
		lianjia.CrawlAllPublishedDataByCommunity(v)
	}
}