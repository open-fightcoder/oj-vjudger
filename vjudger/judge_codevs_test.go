package vjudger

import (
	"fmt"
	"testing"
)

func TestCodeVSJudger_Submit(t *testing.T) {
	c := CodeVSJudger{}
	a := c.Submit("1000", "c++", "#include<iostream>using namespace std;int main(){int a,b;while(cin>>a>>b){cout<<a+b<<endl;}return 0;}")
	fmt.Println(a)
}

func TestGetResult(t *testing.T) {
	c := CodeVSJudger{}
	a:=c.Submit("1000", "C++", "#include<iostream>using namespace std;int main(){int a,b;while(cin>>a>>b){cout<<a+b<<endl;}return 0;}")
	fmt.Println(a)
	//time.Sleep(3*time.Second)
	//c.GetResult("3434468")
	//c.GetResult("3434306")
	//c.GetResult("3443012")

}
