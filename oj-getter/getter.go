package ojgetter

import (
	//log "github.com/sirupsen/logrus"
)

func StartGetter() {
	//定时
	getter()
}

func getter() {
	//defer func() {
	//	if err := recover(); err != nil {
	//		log.Errorf("OJ Getter Error: %v", err)
	//	}
	//}()

	//CodeVSGetter{}.getter()
	HDUGetter{}.getter()
}
