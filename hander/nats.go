package hander

import (
	"context"
	"fmt"
	
	aliPB "github.com/gomsa/aliyun/proto/aliyun"
	aliyunClient "github.com/gomsa/aliyun/client"

	pb "github.com/gomsa/nats/proto/nats"
	"github.com/gomsa/nats/service"
)

// Nats 消息事件结构
type Nats struct {
	Repo service.Repository
	Sms service.Sms
}

// ProcessCommonRequest 处理公共请求
func (srv *Nats) ProcessCommonRequest(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
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
	return
}

// List 获取所有消息事件
func (srv *Nats) List(ctx context.Context, req *pb.ListQuery, res *pb.Response) (err error) {
	nats, err := srv.Repo.List(req)
	total, err := srv.Repo.Total(req)
	if err != nil {
		return err
	}
	res.Nats = nats
	res.Total = total
	return err
}

// Get 获取消息事件
func (srv *Nats) Get(ctx context.Context, req *pb.Nat, res *pb.Response) (err error) {
	nat, err := srv.Repo.Get(req)
	if err != nil {
		return err
	}
	res.Nat = nat
	return err
}

// Create 创建消息事件
func (srv *Nats) Create(ctx context.Context, req *pb.Nat, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("添加消息事件失败")
	}
	res.Valid = true
	return err
}

// Update 更新消息事件
func (srv *Nats) Update(ctx context.Context, req *pb.Nat, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新消息事件失败")
	}
	res.Valid = valid
	return err
}

// Delete 删除消息事件
func (srv *Nats) Delete(ctx context.Context, req *pb.Nat, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除消息事件失败")
	}
	res.Valid = valid
	return err
}
