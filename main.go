package main

import (
	// 公共引入
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	k8s "github.com/micro/kubernetes/go/micro"

	npb "github.com/gomsa/nats/proto/nats"
	tpb "github.com/gomsa/nats/proto/template"
	db "github.com/gomsa/nats/providers/database"
	"github.com/gomsa/nats/hander"
	"github.com/gomsa/nats/service"
)

func main() {
	srv := k8s.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
	)
	srv.Init()

	repo := &service.TemplateRepository{db.DB}
	tpb.RegisterTemplatesHandler(srv.Server(), &hander.Template{repo})

	// 消息事件服务实现
	sms := &service.SmsHandler{"aliyun"}
	npb.RegisterNatsHandler(srv.Server(), &hander.Nats{
		Repo : repo,
		Sms : sms.NewHandler(),
	})
	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
