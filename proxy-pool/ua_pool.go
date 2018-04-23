package proxypool

import (
	"io/ioutil"
	"math/rand"
	"strings"
)

var (
	uaPool []string
)

func InitUaPool() {
	uaPool = append(uaPool, "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36")
	uaPool = append(uaPool, "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; de-de) AppleWebKit/125.5.5 (KHTML, like Gecko) Safari/125.12")
	uaPool = append(uaPool, "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:46.0) Gecko/20100101 Firefox/46.0")

	uaPool = append(uaPool, getUaFormFile()...)

	for i, ua := range uaPool {
		uaPool[i] = strings.TrimSpace(strings.Replace(ua, "\r", "", -1))
	}
}

func GetRandomUa() string {
	length := len(uaPool)
	if length == 0 {
		return ""
	}
	return uaPool[rand.Intn(length-1)]
}

func getUaFormFile() []string {
	data, err := ioutil.ReadFile("../scripts/ua/ua_pool.txt")
	if err != nil {
		panic(err)
	}

	uaList := strings.Split(string(data), "\n")
	return uaList
}
