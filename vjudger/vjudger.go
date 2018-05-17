package vjudger

import (
	"strconv"
	"time"

	"fmt"

	"github.com/hunterhug/GoTool/log"
	"github.com/open-fightcoder/oj-vjudger/models"
	"github.com/open-fightcoder/oj-vjudger/common/g"
	"github.com/open-fightcoder/oj-vjudger/common/store"
	"os"
	"path/filepath"
	"github.com/minio/minio-go"
	"io/ioutil"
	"github.com/open-fightcoder/oj-vjudger/redis"
	"encoding/json"

)

type JudgeJob struct {
	SubmitType string `json:"submit_type"`
	SubmitId   int64  `json:"submit_id"`
}
type ProblemCount struct {
	AcNum    int64 `json:"ac_num"`
	TotalNum int64 `json:"total_num"`
}

type SubmitCount struct {
	Accepted            int64 `json:"accepted"`
	FailNum             int64 `json:"fail_num"`
	WrongAnswer         int64 `json:"wrong_answer"`
	CompilationError    int64 `json:"compilation_error"`
	TimeLimitExceeded   int64 `json:"time_limit_exceeded"`
	MemoryLimitExceeded int64 `json:"memory_limit_exceeded"`
	OutputLimitExceeded int64 `json:"output_limit_exceeded"`
	RuntimeError        int64 `json:"runtime_error"`
	SystemError         int64 `json:"system_error"`
}

func DoJudger(job *JudgeJob) {
	var judger Judger
	submit := GetData(job.SubmitId)
	//if submit.UserId != 1 || submit.UserId != 2 {
	//	log.Error("no such a user")
	//	return
	//}
	problem,err:=models.ProblemGetById(submit.ProblemId)
	if err!=nil{
		log.Debugf("problem get by id error in DoJudger")
	}

	if problem.UserId == 4 {
		judger = newJudger("HDU")
	} else if problem.UserId == 3 {
		judger = newJudger("CODEVS")
	} else {
		log.Error("no such a user on vjudger")
		return
	}
	//problem, err := models.ProblemGetById(submit.ProblemId)
	if err != nil {
		fmt.Println("get problem error in DoJudger")
	}
	srcProblemId, err := strconv.ParseInt(problem.Remark, 10, 64)
	if err != nil {
		fmt.Println("string to int64 error in srcProblemId")
	}
	//minio get code
	workDir, err := createWorkDir("default", submit.Id, submit.UserId)
	if err!=nil{
		fmt.Println(err)
	}
	err=GetCode(submit.Code,workDir)
	log.Debugf("getcoder error:%v",err)

	file:=workDir+"/"+submit.Code
	bytes,_:=ioutil.ReadFile(file)
	//fmt.Println("total bytes read：",len(bytes))
	codeStr:=fmt.Sprintf("%s",string(bytes))
	//fmt.Println(codeStr,"string read:",string(bytes))




	submitStr := judger.Submit(strconv.FormatInt(srcProblemId, 10), submit.Language, codeStr)
	//result := judger.GetResult(strconv.FormatInt(job.SubmitId, 10))
	result := judger.GetResult(submitStr)

	for {
		if result.ResultCode <= 3 {
			time.Sleep(1000 * time.Millisecond)
			result = judger.GetResult(submitStr)
		} else {
			break
		}
	}
	fmt.Printf("-------------%v\n", result)
	saveResult(submit, result)
}

func GetData(submitId int64) models.Submit {
	submit, err := models.SubmitGetById(submitId)
	if err != nil {
		log.Debug("get submit by id error")
		return models.Submit{}
	}
	return *submit
	//return strconv.FormatInt(submit.ProblemId,10),submit.Language,submit.Code

}

func saveResult(submit models.Submit, result *Result) {
	submit.Result = result.ResultCode
	submit.ResultDes = result.ResultDes
	submit.RunningTime = result.RunningTime
	submit.RunningMemory = result.RunningMemory

	//更新redis
	if submit.Result == Accepted {
		log.Debug("AC")
		log.Debug(isAc(&submit))
		if !isAc(&submit) {
			log.Debug("更新排行榜")
			err := redis.PersonWeekRankUpdate(1, submit.UserId)
			if err != nil {
				log.Debug("PersonWeekRankUpdate:", err.Error())
			}
			err = redis.PersonMonthRankUpdate(1, submit.UserId)
			if err != nil {
				log.Debug("PersonMonthRankUpdate:", err.Error())
			}
			err = redis.RankListUpdate(1, submit.UserId)
			if err != nil {
				log.Debug("RankListUpdate:", err.Error())
			}
		}
		log.Debug("set 1")
		redis.ProblemStatusSet(submit.UserId, submit.ProblemId, 1)
	}

	if submit.Result > Running {
		if submit.Result != Accepted && !isAc(&submit) {
			log.Debug("set 2")
			redis.ProblemStatusSet(submit.UserId, submit.ProblemId, 2)
		}

		jsonStr, err := redis.ProblemCountGet(submit.ProblemId)
		if err != nil {
			log.Error("problemCount get failure:", err.Error())
			return
		}
		var problemCount ProblemCount
		err = json.Unmarshal([]byte(jsonStr), &problemCount)
		if err != nil {
			log.Error("problemCount get failure:", err.Error())
			return
		}
		problemCount.TotalNum += 1
		if submit.Result == Accepted {
			problemCount.AcNum += 1
		}
		data, err := json.Marshal(problemCount)
		if err != nil {
			log.Error("problemCount set failure:", err.Error())
			return
		}
		flag := redis.ProblemCountSet(submit.ProblemId, string(data))
		if !flag {
			log.Error("problemCount set failure")
		}

		jsonStr, err = redis.SubmitCountGet(submit.UserId)
		if err != nil {
			log.Error("submitcount get failure:", err.Error())
			return
		}
		var submitCount SubmitCount
		err = json.Unmarshal([]byte(jsonStr), &submitCount)
		if err != nil {
			log.Error("submitcount get failure:", err.Error())
			return
		}
		switch submit.Result {
		case Accepted:
			submitCount.Accepted += 1
		case WrongAnswer:
			submitCount.WrongAnswer += 1
		case CompilationError:
			submitCount.CompilationError += 1
		case TimeLimitExceeded:
			submitCount.TimeLimitExceeded += 1
		case MemoryLimitExceeded:
			submitCount.MemoryLimitExceeded += 1
		case OutputLimitExceeded:
			submitCount.OutputLimitExceeded += 1
		case RuntimeError:
			submitCount.RuntimeError += 1
		case SystemError:
			submitCount.SystemError += 1
		}

		data, err = json.Marshal(submitCount)
		if err != nil {
			log.Error("submitcount set failure:", err.Error())
			return
		}
		flag = redis.SubmitCountSet(submit.UserId, string(data))
		if !flag {
			log.Error("submitcount set failure")
		}
	}


	models.SubmitUpdate(&submit)

}


func isAc(submit *models.Submit) bool {
	submitList, _ := models.SubmitGetByUserId(submit.UserId)
	for _, sub := range submitList {
		if sub.Result == Accepted && sub.ProblemId == submit.ProblemId{
			return true
		}
	}
	return false
}


func GetCode(code string, workDir string) error {
	err := store.MinioClient.FGetObject(g.Conf().Minio.CodeBucket,
		code, workDir+"/"+code, minio.GetObjectOptions{})
	return err
}

func createWorkDir(judgeType string, submitId int64, userId int64) (string, error) {
	dir := fmt.Sprintf("%s/work/%s/%d_%d", getCurrentPath(), judgeType, submitId, userId)
	err := os.MkdirAll(dir, 0777)
	return dir, err
}

func getCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Error("getCurrentPath: " + err.Error())
	}
	return dir
}

