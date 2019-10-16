package lianjia

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	CrawlAllPublishedDataByCommunity("新龙城")
}


func TestCrawlOneHouseInfo(t *testing.T) {
	r :=  CrawlOneHouseInfo(101105228948)
	log.Println(r)
}
