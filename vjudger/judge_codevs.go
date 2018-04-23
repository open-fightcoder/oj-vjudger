package vjudger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"sync"

	"strconv"

	log "github.com/sirupsen/logrus"
)

var (
	codevsUserList = []string{"tyming@fightcoder.com"}
	codevsPassList = []string{"tyming"}
	codevsMutexMap map[string]*sync.Mutex
)

var CodeVSRes = map[string]int{

	"等待测试 Pending":                0,
	"测试通过 Accepted":               1,
	"编译错误 Compile Error":          2,
	"测试失败 Rejected":               3,
	"正在评测 Running":                4,
	"答案错误 Wrong Answer":           5,
	"题目无效 Invalid Problem":        6,
	"非法调用 Restricted Call":        7,
	"运行错误 Runtime Error":          8,
	"暂无数据 Data Missed":            9,
	"超出时间 Time Limit Exceeded":    10,
	"超出空间 Memory Limit Exceeded":  11,
	"过多输出 Output Limit Exceeded":  12,
	"等待重测 Rejudge Pending":        13,
	"运行错误(内存访问非法) Runtime Error":  14,
	"运行错误(浮点错误)    Runtime Error": 15,
	"正在编译 COMPILING":              16}

var CodeVSLang = map[string]string{
	"C":      "c",
	"C++":    "cpp",
	"Pascal": "pas"}

type CodeVSJudger struct {
}

func (this *CodeVSJudger) Submit(problemId, language, code string) (string, *http.Client) {

	//init jar
	jar, _ := cookiejar.New(nil)
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	client := &http.Client{Jar: jar}
	//获取cookie
	client.Get("http://www.codevs.cn/")

	index := rand.Intn(5)
	index = index % len(codevsUserList)
	//login data
	values := map[string]string{"username": codevsUserList[index], "password": codevsPassList[index]}
	jsonStr, _ := json.Marshal(values)
	req, err := http.NewRequest("POST", "https://login.codevs.com/api/auth/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("POST https://login.codevs.com/api/auth/login error:", err)
		return "", nil
	}

	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.86 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err", err)
		return "", nil
	}
	defer resp.Body.Close()
	jwtResp, _ := ioutil.ReadAll(resp.Body)

	var f interface{}
	err = json.Unmarshal(jwtResp, &f)
	if err != nil {
		fmt.Println(err)
	}

	jwtInterface := f.(map[string]interface{})
	jwt := jwtInterface["jwt"]
	//判断帐号密码
	html := string(jwtResp)
	if strings.Index(html, "Unable to login with provided credentials.") >= 0 {
		log.Println("username or password error")
		return "", nil
	}

	client.Get("https://login.codevs.com/auth/redirect/?next=http://codevs.cn/accounts/token/login/&token")
	req, err = http.NewRequest("GET", "https://login.codevs.com/api/auth/token", nil)
	if err != nil {
		log.Println("GET https://login.codevs.com/api/auth/token error")
		return "", nil
	}

	req.Header.Add("Authorization", "JWT "+jwt.(string))
	resp, _ = client.Do(req)
	getTokenResp, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(getTokenResp, &f)
	if err != nil {
		fmt.Println(err)
	}

	tokenInterface := f.(map[string]interface{})
	token := tokenInterface["token"]

	client.Get("http://codevs.cn/accounts/token/login/?token=" + token.(string))
	req, err = http.NewRequest("GET", "http://codevs.cn/problem/1000/", nil)
	if err != nil {
		log.Println("GET http://codevs.cn/problem/1000/ error")
		return "", nil
	}
	q, _ := client.Do(req)
	w, _ := ioutil.ReadAll(q.Body)
	html = string(w)
	//fmt.Println(html)
	uv := url.Values{}
	uv.Add("code", code)
	uv.Add("id", problemId)
	uv.Add("format", CodeVSLang[language])
	uv.Add("csrfmiddlewaretoken", client.Jar.Cookies(req.URL)[0].Value)

	req, err = http.NewRequest("POST", "http://codevs.cn/judge/", strings.NewReader(uv.Encode()))
	if err != nil {
		return "", nil
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.86 Safari/537.36")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Referer", "http://codevs.cn/problem/"+problemId+"/")

	resp, err = client.Do(req)
	if err != nil {
		log.Println("POST http://codevs.cn/judge/ error:", err)
		return "", nil
	}

	//fmt.Println(resp.StatusCode)
	judgeResp, _ := ioutil.ReadAll(resp.Body)
	html = string(judgeResp)
	err = json.Unmarshal(judgeResp, &f)
	if err != nil {
		log.Println(err)
	}

	idInterface := f.(map[string]interface{})
	submitIdFloat := idInterface["id"].(float64)
	submitIdString := strconv.Itoa(int(submitIdFloat))
	//fmt.Println(submitIdString)
	return submitIdString, client
}

func GetResult(submitId string) *Result {
	jar, _ := cookiejar.New(nil)
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	client := &http.Client{Jar: jar}
	//获取cookie
	client.Get("http://www.codevs.cn/")

	index := rand.Intn(5)
	index = index % len(codevsUserList)
	//login data
	values := map[string]string{"username": codevsUserList[index], "password": codevsPassList[index]}
	jsonStr, _ := json.Marshal(values)
	req, err := http.NewRequest("POST", "https://login.codevs.com/api/auth/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("POST https://login.codevs.com/api/auth/login error:", err)
		return nil
	}

	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.86 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err", err)
		return nil
	}
	defer resp.Body.Close()
	jwtResp, _ := ioutil.ReadAll(resp.Body)

	var f interface{}
	err = json.Unmarshal(jwtResp, &f)
	if err != nil {
		fmt.Println(err)
	}

	jwtInterface := f.(map[string]interface{})
	jwt := jwtInterface["jwt"]
	//判断帐号密码
	html := string(jwtResp)
	if strings.Index(html, "Unable to login with provided credentials.") >= 0 {
		log.Println("username or password error")
		return nil
	}

	client.Get("https://login.codevs.com/auth/redirect/?next=http://codevs.cn/accounts/token/login/&token")
	req, err = http.NewRequest("GET", "https://login.codevs.com/api/auth/token", nil)
	if err != nil {
		log.Println("GET https://login.codevs.com/api/auth/token error")
		return nil
	}

	req.Header.Add("Authorization", "JWT "+jwt.(string))
	resp, _ = client.Do(req)
	getTokenResp, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(getTokenResp, &f)
	if err != nil {
		fmt.Println(err)
	}

	tokenInterface := f.(map[string]interface{})
	token := tokenInterface["token"]

	client.Get("http://codevs.cn/accounts/token/login/?token=" + token.(string))
	req, err = http.NewRequest("GET", "http://codevs.cn/problem/1000/", nil)
	if err != nil {
		log.Println("GET http://codevs.cn/problem/1000/ error")
		return nil
	}
	q, _ := client.Do(req)
	w, _ := ioutil.ReadAll(q.Body)
	html = string(w)

	///////////////////////////////
	url := "http://codevs.cn/submission/api/refresh/?id=" + submitId + "&waiting_time=0"

	resp, err = client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	src := string(body)
	fmt.Printf("%v", src)
	//var f interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		fmt.Println(err)
	}

	respInterface := f.(map[string]interface{})
	status := respInterface["status"]
	results := respInterface["results"]
	statusStr := status.(string)
	statusStr = fmt.Sprintf("%v", statusStr)
	resultsStr := results.(string)
	resultsStr = fmt.Sprintf("%v", resultsStr)
	memoryCostStr := status.(string)
	memoryCostStr = fmt.Sprintf("%v", memoryCostStr)
	timeCostStr := status.(string)
	timeCostStr = fmt.Sprintf("%v", timeCostStr)
	memoryCostInt, err := strconv.ParseInt(memoryCostStr, 10, 64)
	timeCostInt, err := strconv.ParseInt(timeCostStr, 10, 64)
	//fmt.Println(statusStr, resultsStr)
	result := &Result{}
	result.ResultDes = resultsStr
	result.RunningMemory = memoryCostInt
	result.RunningTime = timeCostInt

	//将html标签全部转换成小写
	//re, _ := regexp.Compile(`<[\S\s]+?>`)
	//src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//<label class="label bg-.*?">(.*?)</label>
	//re,err = regexp.Compile(`<tr run_id=`+problemId)
	//if err!=nil{
	//	fmt.Println("---------------",err)
	//}
	//temps := re.FindAllStringSubmatch(src, -1)
	//fmt.Println(temps)
	//for i:=1;i<len(temps);i++{
	//	fmt.Println("------ID---------",temps[i][1])
	//	fmt.Println("-------res----------",temps[i][2])
	//	if temps[i][1] == problemId{
	//		re,_=regexp.Compile(`<label class="label .*?">(.*?)</label>`)
	//		result:=re.FindStringSubmatch(temps[i][2])
	//		fmt.Println(result)
	//		return nil
	//	}
	//}

	return result
}
