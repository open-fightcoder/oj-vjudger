package vjudger

type JudgeJob struct {
	SubmitType string `json:"submit_type"`
	SubmitId int64 `json:"submit_id"`
}

func DoJudger(job *JudgeJob) {
	GetData(job.SubmitId)

	judger := newJudger("HDU")
	judger.Submit("","","")
	judger.GetResult("")

	saveResult(nil)
}

func GetData(submitId int64) {

}

func saveResult(result *Result) {

}
