package ojgetter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"github.com/open-fightcoder/oj-vjudger/models"
)

const (
	codevsUrl    = "http://www.codevs.cn/problem/"
	codevsUserId = 2
)

type CodeVSGetter struct {
	BaseGetter
}

//func (c *CodeVSGetter) getter() {
//
//}
func (this CodeVSGetter) getProblemIdMax() int {
	return 1005
}

func (this CodeVSGetter) getter() {

	end := this.getProblemIdMax()
	for i := 1000; i < end; i++ {
		c := CodeVSGetter{}
		problem := c.getProblem(i)
		c.Save(problem)
	}

}
func (this CodeVSGetter) Save(problem models.Problem) {

	problemId := problem.CaseData
	userProblems, err := models.ProblemUser{}.QueryByCaseData(problemId)
	if err != nil {
		panic("QueryByCaseData error:" + err.Error())
	}

	if len(userProblems) > 0 {
		problem.Id = userProblems[0].Id
		this.update(problem, "user")

		checkProblem, err := models.ProblemCheck{}.QueryByCaseData(problemId)
		if err != nil {
			panic("QueryByCaseData error:" + err.Error())
		}

		if len(checkProblem) > 0 {
			this.update(problem, "check")
		} else {
			this.save(problem, "check")
		}

	} else {
		this.save(problem, "user")
	}

	//if problem.Description != "" && problem.InputDes != "" && problem.InputCase != "" && problem.OutputDes != "" && problem.OutputCase != "" {
	//	problem.Create(&problem)
	//	fmt.Println(problem)
	//}

}

func (this CodeVSGetter) update(problem models.Problem, problemType string) {
	problemJson, err := json.Marshal(problem)
	if err != nil {
		panic("CodeVSGetter update: " + err.Error())
	}

	switch problemType {
	case "user":
		problemUser := models.ProblemUser{}
		if err := json.Unmarshal(problemJson, &problemUser); err != nil {
			panic("CodeVSGetter update: " + err.Error())
		}

		models.ProblemUser{}.Create(&problemUser)
	case "check":
		problemCheck := models.ProblemCheck{}
		if err := json.Unmarshal(problemJson, &problemCheck); err != nil {
			panic("CodeVSGetter save: " + err.Error())
		}

		problemCheck.UserId = codevsUserId

		models.ProblemCheck{}.Create(&problemCheck)
	default:
		panic("CodeVSGetter save: not match problemType " + problemType)
	}
}

func (this CodeVSGetter) save(problem models.Problem, problemType string) {
	problemJson, err := json.Marshal(problem)
	if err != nil {
		panic("CodeVSGetter save: " + err.Error())
	}

	switch problemType {
	case "user":
		problemUser := models.ProblemUser{}
		if err := json.Unmarshal(problemJson, &problemUser); err != nil {
			panic("CodeVSGetter save: " + err.Error())
		}

		problemUser.UserId = codevsUserId

		models.ProblemUser{}.Create(&problemUser)
	case "check":
		problemCheck := models.ProblemCheck{}
		if err := json.Unmarshal(problemJson, &problemCheck); err != nil {
			panic("CodeVSGetter save: " + err.Error())
		}

		problemCheck.UserId = codevsUserId

		models.ProblemCheck{}.Create(&problemCheck)
	default:
		panic("CodeVSGetter save: not match problemType " + problemType)
	}
}

//处理对应的题目"http://www.codevs.cn/problem/1001/
func (this CodeVSGetter) getProblem(id int) models.Problem {

	problem := models.Problem{}
	url := codevsUrl + strconv.Itoa(id)

	//fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		//fmt.Println(err)
		fmt.Println("http.Get() error!")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll() error!")
		return problem
	}

	src := string(body)

	//将html标签全部转换成小写
	re, _ := regexp.Compile(`<[\S\s]+?>`)
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//获取title
	reGetTitle, _ := regexp.Compile(`<h3 class="m-t m-b-sm" style="display:inline-block">  <b>([\S\s]+?)</b></h3>`)
	title := reGetTitle.FindStringSubmatch(src)
	//fmt.Println(title, title[1])
	//获取limit
	reLimit, _ := regexp.Compile(`<i class="fa fa-clock-o fa fa-2x fa icon-muted v-middle"></i>[\S\s]+?(\d+) s[\S\s]+?</span>[\S\s]+?<i class="fa fa-flask fa fa-2x fa icon-muted v-middle"></i>[\S\s]+?(\d+)[\S\s]+?</span>`)
	limit := reLimit.FindStringSubmatch(src)
	fmt.Println(limit, "---------", limit[1], limit[2])
	time, _ := strconv.Atoi(limit[1])
	memory, _ := strconv.Atoi(limit[2])
	//匹配需要的数据,添加外层div防止非目标p标签的干扰
	//re, _ = regexp.Compile(`<div class="panel-body">[\S\s]+?<p>[\S\s]+?</p>[\S\s]+?</div>`)
	re, _ = regexp.Compile(`<div class="panel-body">[\S\s]+?</div>`)
	temps := re.FindAllString(src, -1)

	for i := 0; i < len(temps); i++ {

		//读取p中的内容
		//re, _ = regexp.Compile(`<p>[\S\s]+?</p>`)
		//temps[i] = re.FindString(temps[i])
		re, _ = regexp.Compile(`<[\S\s]+?>`)
		temps[i] = re.ReplaceAllString(temps[i], "")

	}

	problem.CaseData = strconv.Itoa(id)
	problem.Titile = title[1]
	problem.Description = temps[0]
	problem.InputDes = temps[1]
	problem.OutputDes = temps[2]
	problem.InputCase = temps[3]
	problem.OutputCase = temps[4]
	problem.Hint = temps[5]
	fmt.Println(temps[5])
	problem.TimeLimit = time * 1000
	problem.MemoryLimit = memory
	return problem
}
