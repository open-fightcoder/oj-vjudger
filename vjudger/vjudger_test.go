package vjudger

import (
	"testing"
	"fmt"
	"github.com/open-fightcoder/oj-vjudger/common/g"
	"github.com/open-fightcoder/oj-vjudger/common/store"
	"io/ioutil"
)

func TestGetCode(t *testing.T) {
	g.LoadConfig("../cfg/cfg.toml.debug")
	store.InitMinio()
	workDir, err := createWorkDir("default", 35, 2)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("----------------",workDir)
	err=GetCode("10_1525422879.c",workDir)
	fmt.Println(err)

	file:=workDir+"/"+"10_1525422879.c"
	bytes,_:=ioutil.ReadFile(file)
	fmt.Println("total bytes read：",len(bytes))
	fmt.Println("string read:",string(bytes))

	codeStr:=fmt.Sprintf("%s",string(bytes))
	fmt.Println("+++++++++++",codeStr)
}
