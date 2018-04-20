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
	hduBaseUrl = "http://acm.hdu.edu.cn/showproblem.php?pid="
	HduUserId  = 1
)

type HDUGetter struct {
	BaseGetter
}

func (this HDUGetter) getter() {

	end := this.getProblemIdMax()
	for i := 1000; i < end; i++ {
		h := HDUGetter{}
		problem := h.getProblem(i)
		h.Save(problem)
	}

}

func (this HDUGetter) Save(problem models.Problem) {

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

func (this HDUGetter) update(problem models.Problem, problemType string) {
	problemJson, err := json.Marshal(problem)
	if err != nil {
		panic("HDUGetter update: " + err.Error())
	}

	switch problemType {
	case "user":
		problemUser := models.ProblemUser{}
		if err := json.Unmarshal(problemJson, &problemUser); err != nil {
			panic("HDUGetter update: " + err.Error())
		}

		models.ProblemUser{}.Create(&problemUser)
	case "check":
		problemCheck := models.ProblemCheck{}
		if err := json.Unmarshal(problemJson, &problemCheck); err != nil {
			panic("HDUGetter save: " + err.Error())
		}

		problemCheck.UserId = HduUserId

		models.ProblemCheck{}.Create(&problemCheck)
	default:
		panic("HDUGetter save: not match problemType " + problemType)
	}
}

func (this HDUGetter) save(problem models.Problem, problemType string) {
	problemJson, err := json.Marshal(problem)
	if err != nil {
		panic("HDUGetter save: " + err.Error())
	}

	switch problemType {
	case "user":
		problemUser := models.ProblemUser{}
		if err := json.Unmarshal(problemJson, &problemUser); err != nil {
			panic("HDUGetter save: " + err.Error())
		}

		problemUser.UserId = HduUserId

		models.ProblemUser{}.Create(&problemUser)
	case "check":
		problemCheck := models.ProblemCheck{}
		if err := json.Unmarshal(problemJson, &problemCheck); err != nil {
			panic("HDUGetter save: " + err.Error())
		}

		problemCheck.UserId = HduUserId

		models.ProblemCheck{}.Create(&problemCheck)
	default:
		panic("HDUGetter save: not match problemType " + problemType)
	}
}

func (this HDUGetter) getProblemIdMax() int {
	return 1005
}

//获取对应的题目,例如:"
// /showproblem.php?pid=1000
func (this HDUGetter) getProblem(id int) models.Problem {
	problem := models.Problem{}

	url := hduBaseUrl + strconv.Itoa(id)
	resp, err := http.Get(url)
	if err != nil {
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

	reGetTitle, _ := regexp.Compile(`<h1 style=.*?>([\S\s]+?)</h1><font><b><span style=.*?>([\S\s]+?)</span></b></font>`)
	title := reGetTitle.FindStringSubmatch(src)
	//fmt.Println(title[1])
	//fmt.Println(title[2])

	reLimit, _ := regexp.Compile(`(\d+) MS[\S\s]+?(\d+) K`)
	limit := reLimit.FindStringSubmatch(title[2])
	//fmt.Println("----------", limit[1], limit[2])
	time, _ := strconv.Atoi(limit[1])
	memory, _ := strconv.Atoi(limit[2])
	fmt.Println(time, memory)
	//匹配需要的数据,添加外层div防止非目标p标签的干扰
	re, _ = regexp.Compile(`<div class=panel_content>[\S\s]+?</div>`)
	temps := re.FindAllString(src, -1)

	for i := 0; i < len(temps); i++ {

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
	problem.Hint = temps[6]
	problem.TimeLimit = time
	problem.MemoryLimit = memory

	if problem.Hint == "&nbsp;" {
		problem.Hint = ""
	}
	fmt.Println(problem)
	return problem
}
