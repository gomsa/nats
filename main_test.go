package main

import (
	"fmt"
	"context"
	"testing"

	_ "github.com/gomsa/nats/database/migrations"
	"github.com/gomsa/nats/hander"
	npb "github.com/gomsa/nats/proto/nats"
	db "github.com/gomsa/nats/providers/database"
	"github.com/gomsa/nats/service"
	"github.com/gomsa/tools/env"
)

func TestProcessCommonRequest(t *testing.T) {

	nrepo := &service.SmsHandler{
		env.Getenv("SMS_DRIVE", "cloopen"),
	}
	repo := &service.TemplateRepository{db.DB}
	sms,err := nrepo.NewHandler()
	if err != nil {
		fmt.Println(err)
	}
	h := hander.Nats{Repo:repo,Sms:sms}
	req := &npb.Request{
		Addressee: `15550251272`,
		Event:     `register_verify`,
		Type:      `sms`,
		QueryParams: map[string]string{
			`code`: `123456`,
			`time`: `30`,
		},
	}
	res := &npb.Response{}
	err = h.ProcessCommonRequest(context.TODO(), req, res)
	fmt.Println(req, res, err)
	t.Log(req, res, err)
}
