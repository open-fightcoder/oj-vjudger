package redis

import (
	"testing"

	"fmt"

	"github.com/open-fightcoder/oj-judger/common/g"
	"github.com/open-fightcoder/oj-judger/common/store"
)

func TestProblemCountSet(t *testing.T) {
	g.LoadConfig("../cfg/cfg.toml.debug")
	store.InitRedis()

	for i := 1; i <= 2300; i++ {
		ProblemCountSet(int64(i), "{\"ac_num\":0,\"total_num\":0}")
	}
}

func TestProblemCountGet(t *testing.T) {
	g.LoadConfig("../cfg/cfg.toml.debug")
	store.InitRedis()

	aa, err := ProblemCountGet(1)
	fmt.Println(aa, err)
}
