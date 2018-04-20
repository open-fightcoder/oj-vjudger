package models

import (
	. "github.com/open-fightcoder/oj-vjudger/common/store"
)

type ProblemUser struct {
	Id                 int64
	UserId             int64  //题目提供者
	CaseData           string //测试数据
	Titile             string //题目标题
	Description        string //题目描述
	InputDes           string //输入描述
	OutputDes          string //输出描述
	InputCase          string //测试输入
	OutputCase         string //测试输出
	Hint               string //题目提示(可以为对样例输入输出的解释)
	TimeLimit          int    //时间限制
	MemoryLimit        int    //内存限制
	Tag                string //题目标签
	IsSpecialJudge     bool   //是否特判
	SpecialJudgeSource string //特判程序源代码
	Code               string //标准程序
	LanguageLimit      string //语言限制
	Note               string //备注(用户不可见)
}

//增加
func (this ProblemUser) Create(problemUser *ProblemUser) (int64, error) {
	_, err := OrmWeb.Insert(problemUser) //第一个参数为影响的行数
	if err != nil {
		return 0, err
	}
	return problemUser.Id, nil
}

//删除
func (this ProblemUser) Remove(id int64) error {
	problemUser := new(ProblemUser)
	_, err := OrmWeb.Id(id).Delete(problemUser)
	return err
}

//修改
func (this ProblemUser) Update(problemUser *ProblemUser) error {
	_, err := OrmWeb.AllCols().ID(problemUser.Id).Update(problemUser)
	return err
}

//查询
func (this ProblemUser) GetById(id int64) (*ProblemUser, error) {
	problemUser := new(ProblemUser)
	has, err := OrmWeb.Id(id).Get(problemUser) //第一个为 bool 类型，表示是否查找到记录

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return problemUser, nil
}

func (this ProblemUser) QueryByUserId(userId int64) ([]*ProblemUser, error) {
	problemUser := make([]*ProblemUser, 0)
	err := OrmWeb.Where("user_id = ?", userId).Find(&problemUser)
	if err != nil {
		return nil, err
	}
	return problemUser, nil
}

func (this ProblemUser) QueryByCaseData(str string) ([]*ProblemUser, error) {
	problemUser := make([]*ProblemUser, 0)
	err := OrmWeb.Where("case_data = ?", str).Find(&problemUser)
	if err != nil {
		return problemUser, err
	}
	return problemUser, nil
}
