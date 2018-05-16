package vjudger

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

const DefaultTimeOut = 2

func NewJar() *cookiejar.Jar {
	cookieJar, _ := cookiejar.New(nil)
	return cookieJar
}

func Second(times int) time.Duration {
	return time.Duration(times) * time.Second
}

func NewProxyClient(proxystring string) (*http.Client, error) {
	proxy, err := url.Parse(proxystring)
	if err != nil {
		return nil, err
	}

	// This a alone client, diff from global client.
	client := &http.Client{
		// Allow redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			log.Printf("Redirect:%v", req.URL)
			return nil
		},
		// Allow proxy
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
		// Allow keep cookie
		Jar: NewJar(),
		// Allow Timeout
		Timeout: Second(DefaultTimeOut),
	}
	return client, nil
}

func NewClient() (*http.Client, error) {
	client := &http.Client{
		// Allow redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			log.Printf("Redirect:%v", req.URL)
			return nil
		},
		Jar:     NewJar(),
		Timeout: Second(DefaultTimeOut),
	}
	return client, nil
}
