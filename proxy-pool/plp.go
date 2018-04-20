package proxypool

import (
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

//PLP get ip from proxylistplus.com
func PLP() {
	pollURL := "https://list.proxylistplus.com/Fresh-HTTP-Proxy-List-1"
	_, body, errs := gorequest.New().Get(pollURL).End()
	if errs != nil {
		log.Println(errs)
		return
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Println(err.Error())
		return
	}
	doc.Find("#page > table.bg > tbody > tr").Each(func(i int, s *goquery.Selection) {
		node := strconv.Itoa(i + 1)
		ss, _ := s.Find("tr:nth-child(" + node + ") > td:nth-child(2)").Html()
		sss, _ := s.Find("tr:nth-child(" + node + ") > td:nth-child(3)").Html()
		ssss, _ := s.Find("tr:nth-child(" + node + ") > td:nth-child(7)").Html()
		if ssss == "yes" {
			ssss = "http,https"
		} else if ssss == "no" {
			ssss = "http"
		}
		ipData := ss + ":" + sss
		ipType := ssss
		//result = append(result, ip)
		//fmt.Println("size", len(ipChan), ip)
		//if ip.Data != "" && ip.Type != "" && util.CheckIP(ip) {
		//	ipChan <- ip
		//}
		if ipData != "" && ipType != "" {
			if strings.ToLower(ipType) == "http" {
				PutHttp(ipData)
			} else if strings.ToLower(ipType) == "https" {
				PutHttps(ipData)
			}
		}

	})
	log.Println("PLP done.")
	return
}
