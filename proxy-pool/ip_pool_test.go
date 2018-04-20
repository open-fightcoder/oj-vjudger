package proxypool

import (
	"fmt"
	"testing"
)

func TestStartIpPoool(t *testing.T) {
	StartIpPoool()
	fmt.Println(GetHttp())
}
