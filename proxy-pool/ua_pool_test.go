package proxypool

import (
	"fmt"
	"testing"
)

func TestGetUaFromFile(t *testing.T) {
	InitUaPool()

	for i := 0; i < 10; i++ {
		ua := GetRandomUa()
		fmt.Println(ua)
	}
}
