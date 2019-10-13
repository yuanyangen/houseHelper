package lianjia

import "testing"

func TestCrawlAllChengjiao(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestCrawlAllChengjiaoByCommunity(t *testing.T) {
	CrawlAllChengjiaoByCommunity("新龙城")
}