package models

import (
	. "github.com/open-fightcoder/oj-vjudger/common/store"
)

type Problem struct {
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
}

//增加
func (this Problem) Create(problem *Problem) (int64, error) {
	_, err := OrmWeb.Insert(problem) //第一个参数为影响的行数
	if err != nil {
		return 0, err
	}
	return problem.Id, nil
}

//删除
func (this Problem) Remove(id int64) error {
	problem := new(Problem)
	_, err := OrmWeb.Id(id).Delete(problem)
	return err
}

//修改
func (this Problem) Update(problem *Problem) error {
	_, err := OrmWeb.AllCols().ID(problem.Id).Update(problem)
	return err
}

//查询
func (this Problem) GetById(id int64) (*Problem, error) {
	problem := new(Problem)
	has, err := OrmWeb.Id(id).Get(problem) //第一个为 bool 类型，表示是否查找到记录

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return problem, nil
}

func (this Problem) GetByUserId(userId int64) (*Problem, error) {
	problem := new(Problem)
	has, err := OrmWeb.Where("user_id = ?", userId).Get(problem)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return problem, nil
}

func (this Problem) QueryByTitile(titile string) ([]*Problem, error) {
	problemList := make([]*Problem, 0)
	err := OrmWeb.Where("titile like ?", "%"+titile+"%").Find(&problemList)
	if err != nil {
		return nil, err
	}
	return problemList, nil
}
