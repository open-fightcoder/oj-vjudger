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
	hduBaseUrl = "http://acm.hdu.edu.cn/showproblem.php?pid="
	hduUserId  = 1
)

type HDUGetter struct {
	BaseGetter
}

func (this HDUGetter) getter() {

	end := this.getProblemIdMax()
	for i := 1095; i < end; i++ {
		h := HDUGetter{}
		problem := h.getProblem(i)
		if problem.Description==""{
			continue
		}
		h.Save(problem)
	}

}

func (this HDUGetter) Save(problem models.Problem) {

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

func (this HDUGetter) update(problem models.Problem) {
	problem.UserId = hduUserId
	models.ProblemUpdate(&problem)
}

func (this HDUGetter) save(problem models.Problem) {
	problem.UserId = hduUserId
	models.ProblemCreate(&problem)
}

func (this HDUGetter) getProblemIdMax() int {
	return 1500
}

//获取对应的题目,例如:"
// /showproblem.php?pid=1000
func (this HDUGetter) getProblem(id int) models.Problem {
	problem := models.Problem{}

	url := hduBaseUrl + strconv.Itoa(id)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get() error!")
		return models.Problem{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll() error!")
		return models.Problem{}
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
	//fmt.Println(time, memory)
	//匹配需要的数据,添加外层div防止非目标p标签的干扰
	re, _ = regexp.Compile(`<div class=panel_content>[\S\s]+?</div>`)
	temps := re.FindAllString(src, -1)
	//fmt.Println("------------",len(temps))

	for i := 0; i < len(temps); i++ {

		re, _ = regexp.Compile(`<[\S\s]+?>`)
		temps[i] = re.ReplaceAllString(temps[i], "")

	}

	problem.Remark = strconv.Itoa(id)
	problem.Title = strings.TrimSpace(title[1])
	problem.Description = strings.TrimSpace(temps[0])
	problem.InputDes = strings.TrimSpace(temps[1])
	problem.OutputDes = strings.TrimSpace(temps[2])
	problem.InputCase = strings.TrimSpace(temps[3])
	problem.OutputCase = strings.TrimSpace(temps[4])
	if len(temps)>5{
		problem.Hint = strings.TrimSpace(temps[6])
	}

	problem.TimeLimit = time
	problem.MemoryLimit = memory

	if problem.Hint == "&nbsp;" {
		problem.Hint = ""
	}
	//fmt.Println(problem)
	return problem
}
