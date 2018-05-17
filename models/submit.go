package models

import . "github.com/open-fightcoder/oj-vjudger/common/store"

type Submit struct {
	Id            int64  `xorm:"pk autoincr BIGINT(20)"`
	ProblemId     int64  `xorm:"not null index BIGINT(20)"`
	UserId        int64  `xorm:"not null index BIGINT(20)"`
	Language      string `xorm:"not null VARCHAR(20)"`
	SubmitTime    int64  `xorm:"not null BIGINT(20)"`
	RunningTime   int64    `xorm:"INT(11)"`
	RunningMemory int64    `xorm:"INT(11)"`
	Result        int    `xorm:"index INT(11)"`
	ResultDes     string `xorm:"VARCHAR(500)"`
	Code          string `xorm:"not null VARCHAR(200)"`
}

func SubmitCreate(submit *Submit) (int64, error) {
	return OrmWeb.Insert(submit)
}

func SubmitRemove(id int64) error {
	_, err := OrmWeb.Id(id).Delete(&Submit{})
	return err
}
func SubmitUpdate(submit *Submit) error {
	_, err := OrmWeb.AllCols().ID(submit.Id).Update(submit)
	return err
}

func SubmitGetById(id int64) (*Submit, error) {
	submit := new(Submit)
	has, err := OrmWeb.Id(id).Get(submit)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return submit, nil
}

<<<<<<< HEAD
func SubmitGetByUserId(userId int64) ([]*Submit, error) {
	submitList := make([]*Submit, 0)
	err := OrmWeb.Where("user_id=?", userId).Find(&submitList)
=======
func SubmitGetByUserId(userId int64, currentPage int, perPage int) ([]*Submit, error) {
	submitList := make([]*Submit, 0)
	err := OrmWeb.Where("user_id=?", userId).Limit(perPage, (currentPage-1)*perPage).Find(&submitList)
>>>>>>> bd277c7ed08a2ecae5ddf41d9ae870e984e38ef2
	if err != nil {
		return nil, err
	}
	return submitList, nil
}

<<<<<<< HEAD
//func SubmitGetByUserId(userId int64, currentPage int, perPage int) ([]*Submit, error) {
//	submitList := make([]*Submit, 0)
//	err := OrmWeb.Where("user_id=?", userId).Limit(perPage, (currentPage-1)*perPage).Find(&submitList)
//	if err != nil {
//		return nil, err
//	}
//	return submitList, nil
//}

=======
>>>>>>> bd277c7ed08a2ecae5ddf41d9ae870e984e38ef2
func SubmitGetByProblemId(problemId int64, currentPage int, perPage int) ([]*Submit, error) {
	submitList := make([]*Submit, 0)
	err := OrmWeb.Where("problem_id=?", problemId).Limit(perPage, (currentPage-1)*perPage).Find(&submitList)
	if err != nil {
		return nil, err
	}
	return submitList, nil
}

