package vjudger

import (
	"strconv"
	"time"

	"fmt"

	"github.com/hunterhug/GoTool/log"
	"github.com/open-fightcoder/oj-vjudger/models"
)

type JudgeJob struct {
	SubmitType string `json:"submit_type"`
	SubmitId   int64  `json:"submit_id"`
}

func DoJudger(job *JudgeJob) {
	var judger Judger
	submit := GetData(job.SubmitId)
	//if submit.UserId != 1 || submit.UserId != 2 {
	//	log.Error("no such a user")
	//	return
	//}
	if submit.UserId == 1 {
		judger = newJudger("HDU")
	} else if submit.UserId == 2 {
		judger = newJudger("CODEVS")
	} else {
		log.Error("no such a user on vjudger")
		return
	}
	submitStr := judger.Submit(strconv.FormatInt(submit.ProblemId, 10), submit.Language, submit.Code)
	//result := judger.GetResult(strconv.FormatInt(job.SubmitId, 10))
	result := judger.GetResult(submitStr)

	for {
		if result.ResultDes == "" {
			time.Sleep(1000 * time.Millisecond)
			result = judger.GetResult(submitStr)
		} else {
			break
		}
	}
	fmt.Printf("-------------%v", result)
	saveResult(submit, result)
}

func GetData(submitId int64) models.Submit {
	submit, err := models.SubmitGetById(submitId)
	if err != nil {
		log.Debug("get submit by id error")
	}
	return *submit
	//return strconv.FormatInt(submit.ProblemId,10),submit.Language,submit.Code

}

func saveResult(submit models.Submit, result *Result) {
	submit.Code = strconv.Itoa(result.ResultCode)
	submit.ResultDes = result.ResultDes
	submit.RunningTime = result.RunningTime
	submit.RunningMemory = result.RunningMemory
	models.SubmitUpdate(&submit)

}
