package ojgetter

import (
	"testing"
	"github.com/open-fightcoder/oj-vjudger/common"
)

func TestGetProblem(t *testing.T) {
	common.Init("../cfg/cfg.toml.debug")
	h := HDUGetter{}
	h.getter()
	//h.getProblem("http://acm.hdu.edu.cn/showproblem.php?pid=1000")

	//problem := h.getProblem(1000)
	//fmt.Println(problem)

	//
	//for i := 1000; i < 1500; i++ {
	//	h.getProblem("")
	//}
}
