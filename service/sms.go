package service

import (
	// 公共引入

	aliPB "github.com/gomsa/aliyun/proto/aliyun"
	aliyunClient "github.com/gomsa/aliyun/client"

	pb "github.com/gomsa/nats/proto/nats"
)

//Sms 短信发送接口
type Sms interface {
	Send(role *pb.Request) (*pb.Response, error)
}

// NatsRepository 消息事件仓库
type AliyunSms struct {
}

// List 获取所有消息事件信息
func (repo *AliyunSms) Send(req *pb.Request) (res *pb.Response, err error) {

	aliyunReq := aliPB.Request{
		Method: "POST",
		Scheme: "https",
		Domain: "dysmsapi.aliyuncs.com",
		Version: "2017-05-25",
		ApiName: "SendSms",
		QueryParams: {
			"PhoneNumbers":"13954386521",
			"SignName":"阿里云",
			"TemplateCode":"SMS_135275049",
			"TemplateParam":"{code: '562374'}"
		}
	}
	aliyunRes, err :=aliyunClient.ProcessCommonRequest(ctx, aliyunReq)
	if err != nil {
		return err
	}
	fmt.Println(aliyunRes.Context)
	return nil, nil
}