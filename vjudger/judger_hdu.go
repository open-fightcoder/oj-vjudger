package vjudger

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/parnurzeal/gorequest"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const (
	hduUrl = "http://acm.hdu.edu.cn/status.php?first="
)

var (
	userList = []string{"hduFightcoder"}
	passList = []string{"fightcoder"}
	mutexMap map[string]*sync.Mutex
)

var HDURes = map[string]int{"Queuing": 0,
	"Compiling":                                 Compiling,
	"Running":                                   Running,
	"Accepted":                                  Accepted,
	"Compilation Error":                         CompilationError,
	"Runtime Error<br>(STACK_OVERFLOW)":         RuntimeError,
	"Runtime Error<br>(ACCESS_VIOLATION)":       RuntimeError,
	"Runtime Error<br>(ARRAY_BOUNDS_EXCEEDED)":  RuntimeError,
	"Runtime Error<br>(FLOAT_DENORMAL_OPERAND)": RuntimeError,
	"Runtime Error<br>(FLOAT_DIVIDE_BY_ZERO)":   RuntimeError,
	"Runtime Error<br>(FLOAT_OVERFLOW)":         RuntimeError,
	"Runtime Error<br>(FLOAT_UNDERFLOW )":       RuntimeError,
	"Runtime Error<br>(INTEGER_OVERFLOW)":       RuntimeError,
	"Runtime Error<br>(INTEGER_DIVIDE_BY_ZERO)": RuntimeError,
	"Wrong Answer":                              WrongAnswer,
	"Time Limit Exceeded":                       TimeLimitExceeded,
	"Memory Limit Exceeded":                     MemoryLimitExceeded,
	"Output Limit Exceeded":                     OutputLimitExceeded,
	"Presentation Error":                        SystemError,
	"System Error":                              SystemError}

type HDUJudger struct {
}

func (this *HDUJudger) Submit(problemId, language, code string) string {

	//init jar
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar, Timeout: time.Second * 10}
	client.Get("http://acm.hdu.edu.cn")

	index := rand.Intn(5)
	index = index % len(userList)
	//login data
	uvLogin := url.Values{}
	uvLogin.Add("username", userList[index])
	uvLogin.Add("userpass", passList[index])
	uvLogin.Add("login", "Sign In")

	req, err := http.NewRequest("POST", "http://acm.hdu.edu.cn/userloginex.php?action=login", strings.NewReader(uvLogin.Encode()))
	if err != nil {
		return ""
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Host", "acm.hdu.edu.cn")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.91 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err", err)
		return ""
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	html := string(b)
	if strings.Index(html, "No such user or wrong password.") >= 0 {
		return ""
	}

	uv := url.Values{}
	uv.Add("check", "0")
	uv.Add("problemid", problemId)
	uv.Add("language", language)
	uv.Add("usercode", code)
	//uv.Add("problemid", "1000")
	//uv.Add("language", "2")
	//uv.Add("usercode", "#include<iostream>using namespace std;int main(){int a,b;while(cin>>a>>b){cout<<a+b<<endl;}return 0;}")

	req, err = http.NewRequest("POST", "http://acm.hdu.edu.cn/submit.php?action=submit", strings.NewReader(uv.Encode()))
	if err != nil {
		return ""
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Host", "acm.hdu.edu.cn")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.91 Safari/537.36")

	resp, err = client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	fmt.Println(time.Now())

	b, _ = ioutil.ReadAll(resp.Body)
	html = string(b)

	if strings.Index(html, "One or more following ERROR(s) occurred.") >= 0 {
		return ""
	}

	resp, _ = client.Get("http://acm.hdu.edu.cn/status.php?first=&pid=1000&user=mysake&lang=3&status=0")
	b, _ = ioutil.ReadAll(resp.Body)
	html = string(b)

	re := regexp.MustCompile(`(\d+)</td><td>(.*?)</td><td>(?s:.*?)<font color=.*?>(.*?)</font>.*?<td>(\d+)MS</td>`)
	temps := re.FindAllStringSubmatch(html, -1)
	//fmt.Println("=============", temps[0][1])
	return temps[0][1]

}

func (this *HDUJudger) GetResult(submitId string) *Result {

	//url := hduUrl + strconv.FormatInt(int64(submitId), 10) + "&pid=&user=&lang=0&status=0"
	url := hduUrl + submitId + "&pid=&user=&lang=0&status=0"

	for {
		resp, _, _ := gorequest.New().Get(url).Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.91 Safari/537.36").End()
		b, _ := ioutil.ReadAll(resp.Body)
		html := string(b)
		//fmt.Println(html)

		re := regexp.MustCompile(`(\d+)</td><td>(.*?)</td><td>(?s:.*?)<font color=.*?>(.*?)</font>.*?<td>(\d+)MS</td><td>(\d+)K</td><td>(\d+)B</td>`)
		temps := re.FindAllStringSubmatch(html, -1)

		temp := temps[0]
		fmt.Println(temp[3], temp[4], temp[5], temp[6])

		result := new(Result)
		result.ResultCode = HDURes[temp[3]]
		result.ResultDes = ""
		timeLimit, _ := strconv.ParseInt(temp[4], 10, 64)
		memoryLimit, _ := strconv.ParseInt(temp[5], 10, 64)
		result.RunningTime = timeLimit
		result.RunningMemory = memoryLimit
		fmt.Println(result)
		time.Sleep(time.Second * 1)
		if HDURes[temp[3]] > 1 {
			if HDURes[temp[3]] == 3 {
				result.ResultDes, _ = this.GetCEInfo(submitId)
			}
			return result
		}

	}

}

func (h *HDUJudger) GetCEInfo(rid string) (string, error) {
	resp, _, _ := gorequest.New().Get("http://acm.hdu.edu.cn/viewerror.php?rid=" + rid).End()
	//if err != nil {
	//	log.Println(err)
	//	return "", err
	//}

	b, _ := ioutil.ReadAll(resp.Body)
	pre := "(?s)<pre>(.*?)</pre>"
	re := regexp.MustCompile(pre)
	match := re.FindStringSubmatch(string(b))
	data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(match[1])), simplifiedchinese.GB18030.NewDecoder()))

	return string(data), nil
}
