package proxypool

import "time"

var t = make(chan int)
var isStop = false

func StartGetter() {

	go func() {
		for {
			if isStop {
				break
			}
			YDL()
			PLP()
			time.Sleep(10 * time.Minute)
		}
	}()

}

func StopGetter() {
	isStop = true
}
