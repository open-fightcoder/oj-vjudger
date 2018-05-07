package vjudger

import (
	"fmt"
	"testing"
)

func TestSecond(t *testing.T) {

}
func TestNewProxyClient(t *testing.T) {

}

func TestNewJar(t *testing.T) {
	client, _ := NewClient()
	resp, _ := client.Get("http://www.baidu.com")

	fmt.Println(resp.StatusCode)
}
