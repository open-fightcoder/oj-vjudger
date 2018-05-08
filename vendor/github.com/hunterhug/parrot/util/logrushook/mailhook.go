package logrushook

import (
	"github.com/hunterhug/parrot/util/gomail"
	"github.com/sirupsen/logrus"
)

const (
	format = "20060102 15:04:05"
)

type EmailHook struct {
	Auth   gomail.MailAuth
	Config gomail.MailConfig
}

func (e *EmailHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}
}

func (e *EmailHook) Fire(data *logrus.Entry) error {
	// only has email key then send
	if _, ok := data.Data["email"]; ok {
		subject := data.Level.String() + ":" + data.Time.Format(format)
		//body := data.Message
		e.Config.Subject = subject
		s, _ := data.String()
		e.Config.Body = []byte(s)
		return gomail.SendMail(e.Auth, e.Config)
	}
	return nil
}
