package service

import (
	// 公共引入

	"github.com/gomsa/tools/env"

	pb "github.com/gomsa/nats/proto/nats"
	tpd "github.com/gomsa/nats/proto/template"

	"github.com/gomsa/nats/service/sms"
)

//Sms 短信发送接口
type Sms interface {
	Send(*pb.Request, *tpd.Template) (bool, error)
}

// SmsHandler sms 驱动
type SmsHandler struct {
	Drive string
}

// NewHandler 使用对应驱动使用 sms
func (s *SmsHandler) NewHandler() (handler Sms) {
	switch s.Drive {
	case "aliyun":
		handler = &sms.Aliyun{
			RegionId:        "default",
			AccessKeyId:     env.Getenv("SMS_KEY_ID", "LTAI8sBGWa9JOHXX"),
			AccessKeySecret: env.Getenv("SMS_KEY_SECRET", "NtQPsr1iNr6H9CBfv7lNNqNk6mkoEz"),
			SignName:        env.Getenv("SMS_SIGN_NAME", "汇丰采驿"),
		}
	}
	return handler
}
