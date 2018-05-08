package slack

import (
	"fmt"
	"time"

	"github.com/hunterhug/marmot/miner"
)

var Slack *miner.Worker

func init() {
	Slack = miner.NewAPI()
}

func SentMessage(hook string, message string) ([]byte, error) {
	s := `{"text":"PgToEs: %s | %s"}`
	times := time.Now().UTC().Format("2006-01-02 15:04:05")
	s = fmt.Sprintf(s, times, message)
	Slack.SetUrl(hook)
	Slack.SetBData([]byte(s))
	fmt.Println(hook, s)
	return Slack.PostJSON()
}
