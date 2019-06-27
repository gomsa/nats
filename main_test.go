package main

import (
	"context"
	// "fmt"
	"testing"

	_ "github.com/gomsa/nats/database/migrations"
	"github.com/gomsa/nats/hander"
	npb "github.com/gomsa/nats/proto/nats"
	"github.com/gomsa/nats/service"
	"github.com/gomsa/tools/env"
)

func TestProcessCommonRequest(t *testing.T) {

	nrepo := &service.SmsHandler{
		env.Getenv("SMS_DRIVE", "aliyun"),
	}

	h := hander.Nats{nrepo.NewHandler()}
	req := &npb.Request{
		Addressee: `13954386521`,
		Event:     `123456`,
		Type:      []string{`sms`},
		QueryParams: map[string]string{
			`code`: `123456`,
		},
	}
	res := &npb.Response{}
	err := h.ProcessCommonRequest(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}
