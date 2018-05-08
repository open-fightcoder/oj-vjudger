package models

import (
	"testing"
)

func TestSubmitCreate(t *testing.T) {
	InitAllInTest()
	submit := Submit{ProblemId: 1000, UserId: 3, Language: "C++", SubmitTime: 1524905061, Code: "#include<iostream>using namespace std;int main(){int a,b;while(cin>>a>>b){cout<<a+b<<endl;}return 0;}"}

	if _, err := SubmitCreate(&submit); err != nil {
		t.Error("create submit error")
	}

}

func TestSubmitRemove(t *testing.T) {
	InitAllInTest()

	if SubmitRemove(1) != nil {
		t.Error("submit remove error")
	}
}

func TestSubmitUpdate(t *testing.T) {
	InitAllInTest()
	submit := Submit{Id: 8, ProblemId: 1000, UserId: 2, Language: "C", Code: "#include<stdio.h> int main(void) { int a,b; int sum; scanf(\"%d%d\",&a,&b); sum=a+b; printf(\"%d\",sum); return 0; }"}
	if SubmitUpdate(&submit) != nil {
		t.Error("update submit error")
	}
}

func TestSubmitGetById(t *testing.T) {
	InitAllInTest()

	if _, err := SubmitGetById(1); err != nil {
		t.Error("get submit by id error")
	}
}

func TestSubmitGetByUserId(t *testing.T) {
	InitAllInTest()

	if _, err := SubmitGetByUserId(1, 1, 1); err != nil {
		t.Error("get submit by user id error")
	}
}

func TestSubmitGetByProblemId(t *testing.T) {
	InitAllInTest()

	if _, err := SubmitGetByProblemId(1, 1, 1); err != nil {
		t.Error("get submit by problem id error")
	}
}
