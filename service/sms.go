package service

import (
	// 公共引入

	"github.com/gomsa/tools/env"

	pb "github.com/gomsa/nats/proto/nats"

	"github.com/gomsa/nats/service/sms"
)

//Sms 短信发送接口
type Sms interface {
	Send(*pb.Request) (*pb.Response, error)
	Event(*pb.Request)
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
			AccessKeyId:     env.Getenv("SMS_KEY_ID", ""),
			AccessKeySecret: env.Getenv("SMS_KEY_SECRET", ""),
			SignName:        env.Getenv("SMS_SIGN_NAME", "阿里云短信测试专用"),
		}
	}
	return handler
}
