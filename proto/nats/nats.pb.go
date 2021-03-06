// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/nats/nats.proto

package nats

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	// 手机号 email 微信号 等等消息接收方
	Addressee string `protobuf:"bytes,1,opt,name=addressee,proto3" json:"addressee,omitempty"`
	// 事件 配套魔板 默认发送方式
	Event string `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	// 留空调用模板默认设置 默认实现事件的类型 sms,email,wechat ...
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// 消息参数 验证码 等魔板参数参数
	QueryParams          map[string]string `protobuf:"bytes,4,rep,name=queryParams,proto3" json:"queryParams,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_86369f7f3ad51865, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetAddressee() string {
	if m != nil {
		return m.Addressee
	}
	return ""
}

func (m *Request) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *Request) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Request) GetQueryParams() map[string]string {
	if m != nil {
		return m.QueryParams
	}
	return nil
}

type Response struct {
	Valid                bool     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_86369f7f3ad51865, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "nats.Request")
	proto.RegisterMapType((map[string]string)(nil), "nats.Request.QueryParamsEntry")
	proto.RegisterType((*Response)(nil), "nats.Response")
}

func init() { proto.RegisterFile("proto/nats/nats.proto", fileDescriptor_86369f7f3ad51865) }

var fileDescriptor_86369f7f3ad51865 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0x1b, 0x6d, 0x3b, 0x45, 0x29, 0x43, 0x85, 0xa5, 0x88, 0x94, 0x9c, 0x7a, 0x8a,
	0x50, 0x0f, 0x4a, 0x0f, 0x2a, 0x88, 0x57, 0xa9, 0xfb, 0x0f, 0x56, 0x33, 0x88, 0x68, 0x76, 0xd3,
	0x9d, 0x4d, 0x21, 0xbf, 0xd3, 0x3f, 0x24, 0x99, 0x64, 0xa9, 0x7a, 0x09, 0xf3, 0x4d, 0xde, 0xe3,
	0xbd, 0x1d, 0x38, 0xaf, 0xbc, 0x0b, 0xee, 0xca, 0x9a, 0xc0, 0xf2, 0xc9, 0x85, 0x31, 0x6d, 0xe7,
	0xec, 0x3b, 0x81, 0x91, 0xa6, 0x5d, 0x4d, 0x1c, 0xf0, 0x02, 0x26, 0xa6, 0x28, 0x3c, 0x31, 0x13,
	0xa9, 0x64, 0x99, 0xac, 0x26, 0xfa, 0xb0, 0xc0, 0x39, 0x1c, 0xd3, 0x9e, 0x6c, 0x50, 0x03, 0xf9,
	0xd3, 0x01, 0x22, 0xa4, 0xa1, 0xa9, 0x48, 0x0d, 0x65, 0x29, 0x33, 0x3e, 0xc0, 0x74, 0x57, 0x93,
	0x6f, 0xb6, 0xc6, 0x9b, 0x92, 0x55, 0xba, 0x1c, 0xae, 0xa6, 0xeb, 0xcb, 0x5c, 0xb2, 0xfb, 0xac,
	0xfc, 0xe5, 0x20, 0x78, 0xb2, 0xc1, 0x37, 0xfa, 0xb7, 0x65, 0x71, 0x07, 0xb3, 0xff, 0x02, 0x9c,
	0xc1, 0xf0, 0x93, 0x9a, 0xbe, 0x57, 0x3b, 0xb6, 0x8d, 0xf6, 0xe6, 0xab, 0xa6, 0xd8, 0x48, 0x60,
	0x33, 0xb8, 0x4d, 0xb2, 0x0d, 0x8c, 0x35, 0x71, 0xe5, 0x2c, 0x53, 0xaf, 0xfa, 0x28, 0xc4, 0x39,
	0xd6, 0x1d, 0xa0, 0x82, 0x51, 0x49, 0xcc, 0xe6, 0x3d, 0xba, 0x23, 0xae, 0xef, 0x21, 0x7d, 0x36,
	0x81, 0xf1, 0x06, 0xe6, 0x5b, 0xef, 0xde, 0x88, 0xf9, 0xd1, 0x95, 0xa5, 0xb3, 0xf1, 0x4a, 0xa7,
	0x7f, 0x1e, 0xb2, 0x38, 0x8b, 0xd8, 0xc5, 0x65, 0x47, 0xaf, 0x27, 0x72, 0xdf, 0xeb, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x8a, 0x4d, 0xe4, 0x4e, 0x78, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Nats service

type NatsClient interface {
	// 共处理方法
	ProcessCommonRequest(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type natsClient struct {
	c           client.Client
	serviceName string
}

func NewNatsClient(serviceName string, c client.Client) NatsClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "nats"
	}
	return &natsClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *natsClient) ProcessCommonRequest(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Nats.ProcessCommonRequest", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Nats service

type NatsHandler interface {
	// 共处理方法
	ProcessCommonRequest(context.Context, *Request, *Response) error
}

func RegisterNatsHandler(s server.Server, hdlr NatsHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Nats{hdlr}, opts...))
}

type Nats struct {
	NatsHandler
}

func (h *Nats) ProcessCommonRequest(ctx context.Context, in *Request, out *Response) error {
	return h.NatsHandler.ProcessCommonRequest(ctx, in, out)
}
