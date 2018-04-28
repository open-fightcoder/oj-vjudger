package models

import "testing"

func TestSubmitTestCreate(t *testing.T) {
	InitAllInTest()

	submitTest:=SubmitTest{UserId:1,Language:"GO",SubmitTime:123456,RunningTime:12345,RunningMemory:1234,Result:1,Input:"qqqq",Code:"qqq"}

	if _,err:=SubmitTestCreate(&submitTest);err!=nil{
		t.Error("create submitTest error")
	}
}

func TestSubmitTestRemove(t *testing.T) {
	InitAllInTest()
	if SubmitTestRemove(1)!=nil{
		t.Error("remove submitTest error")
	}
}

func TestSubmitTestUpdate(t *testing.T) {
	InitAllInTest()

	submitTest:=SubmitTest{UserId:1,Language:"GO",SubmitTime:123456,RunningTime:12345,RunningMemory:1234,Result:1,Input:"qqqq",Code:"qqq"}

	if SubmitTestUpdate(&submitTest)!=nil{
		t.Error("update submitTest error")
	}
}

func TestSubmitTestGetById(t *testing.T) {
	InitAllInTest()

	if _,err:=SubmitGetById(1);err!=nil{
		t.Error("get submitTest error")
	}
}

func TestSubmitTestGetByUserId(t *testing.T) {
	InitAllInTest()

	if _,err:=SubmitTestGetByUserId(1,1,1);err!=nil{
		t.Error("get submitTest by userId error")
	}
}




