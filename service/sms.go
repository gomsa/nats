package service

import (
	// 公共引入

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
	return nil, nil
}