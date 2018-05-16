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

	models.SubmitUpdate(&submit)

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

