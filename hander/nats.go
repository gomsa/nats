package hander

import (
	"context"

	pb "github.com/gomsa/nats/proto/nats"
	"github.com/gomsa/nats/service"
)

// Nats 消息事件结构
type Nats struct {
	Sms service.Sms
}

// ProcessCommonRequest 处理公共请求
func (srv *Nats) ProcessCommonRequest(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	for _, Type := range req.Type {
		switch Type {
		case "sms":
			res, err = srv.Sms.Send(req)
		}
	}
	return
}
