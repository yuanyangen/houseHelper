package proxy

import (
	"fmt"
	"testing"
)



func Test_CrawlProxyAddress(t *testing.T) {
	r := crawlProxyAddress()
	fmt.Print(r)
}
