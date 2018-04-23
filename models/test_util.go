package models

import (
	"sync"

	"github.com/open-fightcoder/oj-vjudger/common"
)

var once sync.Once

func InitAllInTest() {
	once.Do(func() {
		common.Init("../cfg/cfg.toml.debug")
	})
}
