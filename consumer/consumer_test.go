package consumer

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/nsqio/go-nsq"
)

const topic = "vjudger"

func TestSendMessDefault(t *testing.T) {
	mess := SendMess{SubmitType: "aaaaaa", SubmitId: 2}
	Nsq{}.send(mess)

}

type Nsq struct{}

type SendMess struct {
	SubmitType string `json:"submit_type"` //提交类型
	SubmitId   int64  `json:"submit_id"`   //提交id
}

func (this Nsq) send(sendMess SendMess) {
	adds := [1]string{"xupt2.fightcoder.com:9002"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(len(adds))
	msg, err := json.Marshal(sendMess)
	if err != nil {
		fmt.Println(err)
	}
	postNsq(adds[num], topic, msg)
}

func postNsq(address, topic string, msg []byte) {
	config := nsq.NewConfig()
	if w, err := nsq.NewProducer(address, config); err != nil {
		panic(err)
	} else {
		w.Publish(topic, msg)
	}
}
