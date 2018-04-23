package vjudger

import (
	"testing"
)

func TestCodeVSJudger_Submit(t *testing.T) {
	c := CodeVSJudger{}
	c.Submit("1000", "C++", "#include<iostream>using namespace std;int main(){int a,b;while(cin>>a>>b){cout<<a+b<<endl;}return 0;}")
}

func TestGetResult(t *testing.T) {
	//c := CodeVSJudger{}
	//a,b:=c.Submit("1000", "C++", "#include<iostream>using namespace std;int main(){int a,b;while(cin>>a>>b){cout<<a+b<<endl;}return 0;}")
	//fmt.Println(a)
	//time.Sleep(3*time.Second)
	GetResult("3428952")
}
