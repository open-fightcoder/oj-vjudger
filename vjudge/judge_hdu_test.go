package vjudge

import (
	"fmt"
	"testing"
)

func TestHDUJudger_GetSubmitId(t *testing.T) {
	h := HDUJudger{}

	a := h.GetResult(h.Submit("1000", "2", "#include<iostream>using namespace std;int main(){int a,b;while(cin>>a>>b){cout<<a+b<<endl;}return 0;}"))
	fmt.Println(a)
	//for {
	//	if (h.GetResult(h.Submit("1000", "2", "#include<iostream>using namespace std;int main(){int a,b;while(cin>>a>>b){cout<<a+b<<endl;}return 0;}")).ResultCode) != HDURes["Queuing"] {
	//		break
	//	}
	//	time.Sleep(1 * time.Second)
	//}

}

func TestHDUJudger_GetCEInfo(t *testing.T) {
	h := HDUJudger{}
	res, _ := h.GetCEInfo("23466507")
	fmt.Println(res)
}
