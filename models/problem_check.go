package models

import (
	. "github.com/open-fightcoder/oj-vjudger/common/store"
)

type ProblemCheck struct {
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
	CheckStatus        string //审核状态
	ProblemId          string //所在正式题库的Id
	ProblemUserId      string //所在私人题库的Id
}

//增加
func (this ProblemCheck) Create(problemCheck *ProblemCheck) (int64, error) {
	_, err := OrmWeb.Insert(problemCheck) //第一个参数为影响的行数
	if err != nil {
		return 0, err
	}
	return problemCheck.Id, nil
}

//删除
func (this ProblemCheck) Remove(id int64) error {
	problemCheck := new(ProblemCheck)
	_, err := OrmWeb.Id(id).Delete(problemCheck)
	return err
}

//修改
func (this ProblemCheck) Update(problemCheck *ProblemCheck) error {
	_, err := OrmWeb.AllCols().ID(problemCheck.Id).Update(problemCheck)
	return err
}

//查询
func (this ProblemCheck) GetById(id int64) (*ProblemCheck, error) {
	problemCheck := new(ProblemCheck)
	has, err := OrmWeb.Id(id).Get(problemCheck) //第一个为 bool 类型，表示是否查找到记录

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return problemCheck, nil
}

func (this ProblemCheck) QueryByCaseData(str string) ([]*ProblemCheck, error) {
	problemUser := make([]*ProblemCheck, 0)
	err := OrmWeb.Where("case_data = ?", str).Find(&problemUser)
	if err != nil {
		return problemUser, err
	}
	return problemUser, nil
}
