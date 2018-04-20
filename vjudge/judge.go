package vjudge

type Judger interface {
	Submit(problemId, language, code string) string
	GetResult(submitId string) *Result
}

func newJudger(ojname string) Judger {
	switch ojname {
	case "HDU":
		return &HDUJudger{}
	default:
		panic("No such oj: " + ojname)
	}
}
