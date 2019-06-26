package main

import (
	// 公共引入
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	k8s "github.com/micro/kubernetes/go/micro"

	pb "github.com/gomsa/nats/proto/nats"
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

	// 消息事件服务实现
	repo := &service.NatsRepository{db.DB}
	pb.RegisterNatsHandler(srv.Server(), &hander.Nats{repo})
	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
