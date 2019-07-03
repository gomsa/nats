package hander

import (
	"context"
	"strings"

	pb "github.com/gomsa/nats/proto/nats"
	tpd "github.com/gomsa/nats/proto/template"
	"github.com/gomsa/nats/service"
	"github.com/micro/go-micro/util/log"
)

// Nats 消息事件结构
type Nats struct {
	Repo service.Repository
	Sms  service.Sms
}

// ProcessCommonRequest 处理公共请求
func (srv *Nats) ProcessCommonRequest(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 查找对应模板
	template, err := srv.Repo.Get(&tpd.Template{
		Event: req.Event,
	})

	if err != nil {
		return
	}

	if req.Type == "" {
		req.Type = template.Type
	}
	Type := strings.Split(req.Type, ",")
	valid := false
	if srv.InSliceString(Type, "sms") {
		valid, err = srv.Sms.Send(req, template)
		if err != nil {
			log.Log(err)
		}
	}
	res.Valid = valid
	return
}

// InSliceString 判断切片是否包含对应字符串
func (srv *Nats) InSliceString(array []string, val string) bool {
	for _, item := range array {
		switch item {
		case val:
			return true
		}
	}
	return false
}
