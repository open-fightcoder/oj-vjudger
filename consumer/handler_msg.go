package consumer

import (
	"encoding/json"

	"github.com/nsqio/go-nsq"
	//"github.com/open-fightcoder/oj-vjudger/common/g"
	"github.com/open-fightcoder/oj-vjudger/vjudger"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Topic string
}

//var handlerChan = make(chan struct{}, g.Conf().Nsq.HandlerCount)
var handlerChan = make(chan struct{}, 2)

func (this *Handler) HandleMessage(m *nsq.Message) error {
	log.Infof("HandbleMessage: ", string(m.Body))

	judgerJob := new(vjudger.JudgeJob)
	if err := json.Unmarshal(m.Body, judgerJob); err != nil {
		log.Errorf("unmarshal JudgerData from NsqMessage failed, err: %v, event:%s", err, m.Body)
		return nil
	}

	log.Infof("consume Message from dispatch: %#v", judgerJob)

	handlerChan <- struct{}{}
	go Do(judgerJob)

	// 返回err为nil表示消费成功
	return nil
}

func Do(judgerJob *vjudger.JudgeJob) {
	defer func() {
		<-handlerChan
	}()

	vjudger.DoJudger(judgerJob)
}
