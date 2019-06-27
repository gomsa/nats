package sms

import (
	// 公共引入

	"encoding/json"
	"errors"

	"github.com/micro/go-micro/util/log"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"

	pb "github.com/gomsa/nats/proto/nats"
)

// ret 返回请求
var ret map[string]string

// Aliyun 阿里云驱动
type Aliyun struct {
	RegionId        string
	AccessKeyId     string
	AccessKeySecret string
	SignName        string
	TemplateCode    string
}

func (repo *Aliyun) Event(req *pb.Request) {

}

// Send 获取所有消息事件信息
func (repo *Aliyun) Send(req *pb.Request) (res *pb.Response, err error) {
	// 创建连接
	client, err := sdk.NewClientWithAccessKey(
		repo.RegionId,
		repo.AccessKeyId,
		repo.AccessKeySecret,
	)
	if err != nil {
		log.Log(err)
		return res, err
	}
	// 配置参数
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["PhoneNumbers"] = req.Addressee
	request.QueryParams["SignName"] = repo.SignName
	request.QueryParams["TemplateCode"] = repo.TemplateCode
	queryParams, err := json.Marshal(req.QueryParams)
	if err != nil {
		log.Log(err)
		return res, err
	}
	request.QueryParams["TemplateParam"] = string(queryParams)

	// 请求
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		log.Log(err)
		return res, err
	}
	// 返回请求数据
	err = json.Unmarshal([]byte(response.GetHttpContentString()), ret)
	if err != nil {
		log.Log(err)
		return res, err
	}
	if ret["Code"] == "OK" {
		res.Valid = true
	} else {
		err = errors.New(ret["Message"])
	}
	return res, err
}
