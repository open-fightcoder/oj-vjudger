package redis

import (
	"testing"

	"github.com/open-fightcoder/oj-vjudger/common/g"
	"github.com/open-fightcoder/oj-vjudger/common/store"
)

func TestRankGet(t *testing.T) {
	g.LoadConfig("../cfg/cfg.toml.debug")
	store.InitRedis()

	RankListUpdate(10, 15)
}
