package models

import . "github.com/open-fightcoder/oj-vjudger/common/store"

type SubmitTest struct {
	Id            int64  `xorm:"pk autoincr BIGINT(20)"`
	UserId        int64  `xorm:"not null index BIGINT(20)"`
	Language      string `xorm:"not null VARCHAR(20)"`
	SubmitTime    int64  `xorm:"not null BIGINT(20)"`
	RunningTime   int    `xorm:"INT(11)"`
	RunningMemory int    `xorm:"INT(11)"`
	Result        int    `xorm:"index INT(11)"`
	Input         string `xorm:"VARCHAR(300)"`
	ResultDes     string `xorm:"VARCHAR(300)"`
	Code          string `xorm:"not null VARCHAR(200)"`
}

func SubmitTestCreate(submitTest *SubmitTest) (int64, error) {
	return OrmWeb.Insert(submitTest)
}

func SubmitTestRemove(id int64) error {
	_, err := OrmWeb.Id(id).Delete(&SubmitTest{})
	return err
}

func SubmitTestUpdate(submitTest *SubmitTest) error {
	_, err := OrmWeb.AllCols().ID(submitTest.Id).Update(submitTest)
	return err
}

func SubmitTestGetById(id int64) (*SubmitTest, error) {
	submitTest := new(SubmitTest)
	has, err := OrmWeb.Id(id).Get(submitTest)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return submitTest, nil
}
func SubmitTestGetByUserId(userId int64, currentPage int, perPage int) ([]*SubmitTest, error) {
	submitTestList := make([]*SubmitTest, 0)
	err := OrmWeb.Where("user_id=?", userId).Limit(perPage, (currentPage-1)*perPage).Find(&submitTestList)
	if err != nil {
		return nil, err
	}
	return submitTestList, nil
}