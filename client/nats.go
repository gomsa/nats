package client

import (
	"os"

	"github.com/gomsa/tools/k8s/client"

	natsPB "github.com/gomsa/nats/proto/nats"
	templatePB "github.com/gomsa/nats/proto/template"
)

var (
	// Nats 消息服务客户端
	Nats natsPB.NatsClient
	// Template 模板管理客户端
	Template templatePB.TemplatesClient
)

func init() {
	srvName := os.Getenv("NATS_NAME")

	Nats = natsPB.NewNatsClient(srvName, client.DefaultClient)
	Template = templatePB.NewTemplatesClient(srvName, client.DefaultClient)
}
