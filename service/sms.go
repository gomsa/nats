package service

import (
	// 公共引入

	"fmt"

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
func (s *SmsHandler) NewHandler() (handler Sms, err error) {
	switch s.Drive {
	case "aliyun":
		handler = &sms.Aliyun{
			RegionId:        "default",
			AccessKeyId:     env.Getenv("SMS_KEY_ID", ""),
			AccessKeySecret: env.Getenv("SMS_KEY_SECRET", ""),
			SignName:        env.Getenv("SMS_SIGN_NAME", "阿里云短信测试专用"),
		}
	default:
		return handler, fmt.Errorf("未找 %s SMS 驱动", s.Drive)
	}
	return handler, err
}
