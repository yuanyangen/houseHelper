package proxy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

type ProxyStat struct {
	UsedCount int
	FailedCount int

}

const ProxyPoolSize = 2
const MaxFailCount = 5

var mu sync.Mutex
var AllProxys = map[string]*ProxyStat{}

func GetProxyAddress() string {
	if len(AllProxys) <ProxyPoolSize {
		crawlProxyAddress()
	}
	randNum := time.Now().Nanosecond() % len(AllProxys)
	i := 0
	mu.Lock()
	defer mu.Unlock()
	for proxyAddr,stat := range  AllProxys {
		if i == randNum {
			stat.UsedCount = stat.UsedCount + 1
			return proxyAddr
		}
		i ++
	}
	return ""
}

func crawlProxyAddress()  {
	url := fmt.Sprintf( "http://tpv.daxiangdaili.com/ip/?tid=557974608618518&num=%d&delay=1&filter=on", ProxyPoolSize)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	respB, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	t := strings.Split(string(respB), "\n")
	for _,v :=range t {
		mu.Lock()
		v = strings.TrimSpace(v)
		AllProxys[v] = &ProxyStat{}
		mu.Unlock()
	}
}

func CountFail(proxyAddr string) {
	stat := AllProxys[proxyAddr]
	if stat.FailedCount >= MaxFailCount {
		delete(AllProxys, proxyAddr)
	}
	stat.FailedCount  = stat.FailedCount + 1
}