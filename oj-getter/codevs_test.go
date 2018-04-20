package ojgetter

import (
	"testing"
	"github.com/open-fightcoder/oj-vjudger/common"
)

func TestCodeVSGetter_GetProblem(t *testing.T) {
	common.Init("../cfg/cfg.toml.debug")

	c := CodeVSGetter{}
	c.getter()
	//c.GetProblem("http://www.codevs.cn/problem/1001/")
	//problem := c.GetProblem(1001)
	//fmt.Println(problem)
}
