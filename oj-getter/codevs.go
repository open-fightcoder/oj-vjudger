package ojgetter

import (
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
	return 3646
}

func (this CodeVSGetter) getter() {

	end := this.getProblemIdMax()
	for i := 1000; i < end; i++ {
		c := CodeVSGetter{}
		problem := c.getProblem(i)
		if problem.Description == "" {
			continue
		}
		c.Save(problem)
	}

}
func (this CodeVSGetter) Save(problem models.Problem) {

	srcProblemId := problem.Remark
	problems, err := models.ProblemQueryByRemark(srcProblemId)
	if err != nil {
		panic("QueryByRemark error:" + err.Error())
	}

	if len(problems) > 0 {
		problem.Id = problems[0].Id
		this.update(problem)
	} else {
		this.save(problem)
	}

}

func (this CodeVSGetter) update(problem models.Problem) {
	problem.UserId = codevsUserId
	models.ProblemUpdate(&problem)
}

func (this CodeVSGetter) save(problem models.Problem) {
	problem.UserId = codevsUserId
	problem.LanguageLimit = "c,c++,pascal"
	models.ProblemCreate(&problem)
}

//处理对应的题目"http://www.codevs.cn/problem/1001/
func (this CodeVSGetter) getProblem(id int) models.Problem {

	problem := models.Problem{}
	url := codevsUrl + strconv.Itoa(id)

	//fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Println("http.Get() error!")
		return models.Problem{}
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
	//fmt.Println("------------",title,title[1][4:len(title[1])-1])
	//获取难度


	//获取limit
	reLimit, _ := regexp.Compile(`<i class="fa fa-clock-o fa fa-2x fa icon-muted v-middle"></i>[\S\s]+?(\d+) s[\S\s]+?</span>[\S\s]+?<i class="fa fa-flask fa fa-2x fa icon-muted v-middle"></i>[\S\s]+?(\d+)[\S\s]+?</span>[\S\s]+?<i class="fa fa-trophy fa fa-2x fa icon-muted v-middle"></i>[\S\s]+?(\w+)[\S\s]+?</span>[\S\s]+?`)
	limit := reLimit.FindStringSubmatch(src)
	//fmt.Println(limit, "---------", limit[1], limit[2],limit[3])
	if limit == nil {
		return models.Problem{}
	}
	time, _ := strconv.Atoi(limit[1])
	memory, _ := strconv.Atoi(limit[2])
	difficulty:=limit[3]
	//匹配需要的数据,添加外层div防止非目标p标签的干扰
	//re, _ = regexp.Compile(`<div class="panel-body">[\S\s]+?<p>[\S\s]+?</p>[\S\s]+?</div>`)
	re, _ = regexp.Compile(`<div class="panel-body">[\S\s]+?</div>`)
	temps := re.FindAllString(src, -1)

	for i := 0; i < len(temps); i++ {

		//读取p中的内容
		//re, _ = regexp.Compile(`<p>[\S\s]+?</p>`)
		//temps[i] = re.FindString(temps[i])

		//去除html标签
		//re, _ = regexp.Compile(`<[\S\s]+?>`)
		//temps[i] = re.ReplaceAllString(temps[i], "")
		//替换<div class="panel-body">,</div>

		temps[i] = strings.Replace(temps[i], "<div class=\"panel-body\">", "", -1)
		temps[i] = strings.Replace(temps[i], "</div>", "", -1)
		temps[i] = strings.Replace(temps[i], "<img src=\"", "<img src=\"http://codevs.cn", -1)
		//fmt.Println(i, temps[i], "---------------", len(strings.TrimSpace(temps[i])))
		//清除<p>,<span>中的属性
		re, _ = regexp.Compile(`<p [\S\s]+?>`)
		temps[i] = re.ReplaceAllString(temps[i], "<p>")
		re, _ = regexp.Compile(`<span [\S\s]+?>`)
		temps[i] = re.ReplaceAllString(temps[i], "<span>")
		//fmt.Println(i, temps[i], "---------------", len(strings.TrimSpace(temps[i])))
	}

	problem.Remark = strconv.Itoa(id)
	problem.Title = strings.TrimSpace(title[1][4 : len(title[1])-1])
	problem.Description = strings.TrimSpace(temps[0])
	problem.InputDes = strings.TrimSpace(temps[1])
	problem.OutputDes = strings.TrimSpace(temps[2])
	problem.InputCase = strings.TrimSpace(temps[3])
	//fmt.Println(temps[3], "+++++++++++++", problem.InputCase)
	problem.OutputCase = strings.TrimSpace(temps[4])
	problem.Hint = strings.TrimSpace(temps[5])
	//fmt.Println("-------------------", problem.Title)
	problem.TimeLimit = time * 1000
	problem.MemoryLimit = memory
	//fmt.Println("------------", problem)

	if difficulty == "Bronze"{
		problem.Difficulty = "简单"
	}else if difficulty == "Silver"{
		problem.Difficulty = "中等"
	}else if difficulty == "Gold"{
		problem.Difficulty = "困难"
	}else{
		problem.Difficulty = "极难"
	}

	return problem
}
