package proxypool

import (
	"time"
	"github.com/parnurzeal/gorequest"
)

var httpChan = make(chan string, 20)
var httpsChan = make(chan string, 20)

func StartIpPoool() {
	//PutHttp("")
	StartGetter()
}

//取出立刻放进去
func GetHttp() string {

	for {
		temp := <-httpChan
		if checkHttp(temp) {
			PutHttp(temp)
			return temp
		}

	}
}

func GetHttps() string {
	for {
		temp := <-httpChan
		if checkHttps(temp) {
			PutHttps(temp)
			return temp
		}

	}
}

func PutHttp(ip string) {

	if checkHttp(ip) {
		httpChan <- ip
	}
}

func PutHttps(ip string) {
	if checkHttps(ip) {
		httpsChan <- ip
	}
}

func checkHttp(ip string) bool {

	pollURL := "http://httpbin.org/get"
	resp, _, errs := gorequest.New().Timeout(5000 * time.Millisecond).Proxy("//" + ip).Post(pollURL).End()
	if errs != nil {
		return false
	}
	if resp.StatusCode == 200 {
		return true
	}
	return false
}

func checkHttps(ip string) bool {
	pollURL := "https://httpbin.org/post"
	resp, _, errs := gorequest.New().Timeout(5000 * time.Millisecond).Proxy("//" + ip).Post(pollURL).End()
	if errs != nil {
		return false
	}
	if resp.StatusCode == 200 {
		return true
	}
	return false
}
