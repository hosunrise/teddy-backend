// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/message.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/message.proto

It has these top-level messages:
	InBoxEntry
	NotifyEntry
	SendEmailReq
	SendInBoxReq
	SendNotifyReq
	SendSMSReq
	GetInBoxReq
	GetNotifyReq
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = google_protobuf.Empty{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Message service

type MessageService interface {
	SendEmail(ctx context.Context, in *SendEmailReq, opts ...client.CallOption) (*google_protobuf.Empty, error)
	SendInBox(ctx context.Context, in *SendInBoxReq, opts ...client.CallOption) (*google_protobuf.Empty, error)
	SendNotify(ctx context.Context, in *SendNotifyReq, opts ...client.CallOption) (*google_protobuf.Empty, error)
	SendSMS(ctx context.Context, in *SendSMSReq, opts ...client.CallOption) (*google_protobuf.Empty, error)
	GetInBox(ctx context.Context, in *GetInBoxReq, opts ...client.CallOption) (Message_GetInBoxService, error)
	GetNotify(ctx context.Context, in *GetNotifyReq, opts ...client.CallOption) (Message_GetNotifyService, error)
}

type messageService struct {
	c    client.Client
	name string
}

func NewMessageService(name string, c client.Client) MessageService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "com.teddy.srv.notify"
	}
	return &messageService{
		c:    c,
		name: name,
	}
}

func (c *messageService) SendEmail(ctx context.Context, in *SendEmailReq, opts ...client.CallOption) (*google_protobuf.Empty, error) {
	req := c.c.NewRequest(c.name, "Message.SendEmail", in)
	out := new(google_protobuf.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) SendInBox(ctx context.Context, in *SendInBoxReq, opts ...client.CallOption) (*google_protobuf.Empty, error) {
	req := c.c.NewRequest(c.name, "Message.SendInBox", in)
	out := new(google_protobuf.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) SendNotify(ctx context.Context, in *SendNotifyReq, opts ...client.CallOption) (*google_protobuf.Empty, error) {
	req := c.c.NewRequest(c.name, "Message.SendNotify", in)
	out := new(google_protobuf.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) SendSMS(ctx context.Context, in *SendSMSReq, opts ...client.CallOption) (*google_protobuf.Empty, error) {
	req := c.c.NewRequest(c.name, "Message.SendSMS", in)
	out := new(google_protobuf.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) GetInBox(ctx context.Context, in *GetInBoxReq, opts ...client.CallOption) (Message_GetInBoxService, error) {
	req := c.c.NewRequest(c.name, "Message.GetInBox", &GetInBoxReq{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &messageServiceGetInBox{stream}, nil
}

type Message_GetInBoxService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*InBoxEntry, error)
}

type messageServiceGetInBox struct {
	stream client.Stream
}

func (x *messageServiceGetInBox) Close() error {
	return x.stream.Close()
}

func (x *messageServiceGetInBox) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *messageServiceGetInBox) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *messageServiceGetInBox) Recv() (*InBoxEntry, error) {
	m := new(InBoxEntry)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageService) GetNotify(ctx context.Context, in *GetNotifyReq, opts ...client.CallOption) (Message_GetNotifyService, error) {
	req := c.c.NewRequest(c.name, "Message.GetNotify", &GetNotifyReq{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &messageServiceGetNotify{stream}, nil
}

type Message_GetNotifyService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*NotifyEntry, error)
}

type messageServiceGetNotify struct {
	stream client.Stream
}

func (x *messageServiceGetNotify) Close() error {
	return x.stream.Close()
}

func (x *messageServiceGetNotify) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *messageServiceGetNotify) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *messageServiceGetNotify) Recv() (*NotifyEntry, error) {
	m := new(NotifyEntry)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Message service

type MessageHandler interface {
	SendEmail(context.Context, *SendEmailReq, *google_protobuf.Empty) error
	SendInBox(context.Context, *SendInBoxReq, *google_protobuf.Empty) error
	SendNotify(context.Context, *SendNotifyReq, *google_protobuf.Empty) error
	SendSMS(context.Context, *SendSMSReq, *google_protobuf.Empty) error
	GetInBox(context.Context, *GetInBoxReq, Message_GetInBoxStream) error
	GetNotify(context.Context, *GetNotifyReq, Message_GetNotifyStream) error
}

func RegisterMessageHandler(s server.Server, hdlr MessageHandler, opts ...server.HandlerOption) error {
	type message interface {
		SendEmail(ctx context.Context, in *SendEmailReq, out *google_protobuf.Empty) error
		SendInBox(ctx context.Context, in *SendInBoxReq, out *google_protobuf.Empty) error
		SendNotify(ctx context.Context, in *SendNotifyReq, out *google_protobuf.Empty) error
		SendSMS(ctx context.Context, in *SendSMSReq, out *google_protobuf.Empty) error
		GetInBox(ctx context.Context, stream server.Stream) error
		GetNotify(ctx context.Context, stream server.Stream) error
	}
	type Message struct {
		message
	}
	h := &messageHandler{hdlr}
	return s.Handle(s.NewHandler(&Message{h}, opts...))
}

type messageHandler struct {
	MessageHandler
}

func (h *messageHandler) SendEmail(ctx context.Context, in *SendEmailReq, out *google_protobuf.Empty) error {
	return h.MessageHandler.SendEmail(ctx, in, out)
}

func (h *messageHandler) SendInBox(ctx context.Context, in *SendInBoxReq, out *google_protobuf.Empty) error {
	return h.MessageHandler.SendInBox(ctx, in, out)
}

func (h *messageHandler) SendNotify(ctx context.Context, in *SendNotifyReq, out *google_protobuf.Empty) error {
	return h.MessageHandler.SendNotify(ctx, in, out)
}

func (h *messageHandler) SendSMS(ctx context.Context, in *SendSMSReq, out *google_protobuf.Empty) error {
	return h.MessageHandler.SendSMS(ctx, in, out)
}

func (h *messageHandler) GetInBox(ctx context.Context, stream server.Stream) error {
	m := new(GetInBoxReq)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.MessageHandler.GetInBox(ctx, m, &messageGetInBoxStream{stream})
}

type Message_GetInBoxStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*InBoxEntry) error
}

type messageGetInBoxStream struct {
	stream server.Stream
}

func (x *messageGetInBoxStream) Close() error {
	return x.stream.Close()
}

func (x *messageGetInBoxStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *messageGetInBoxStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *messageGetInBoxStream) Send(m *InBoxEntry) error {
	return x.stream.Send(m)
}

func (h *messageHandler) GetNotify(ctx context.Context, stream server.Stream) error {
	m := new(GetNotifyReq)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.MessageHandler.GetNotify(ctx, m, &messageGetNotifyStream{stream})
}

type Message_GetNotifyStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*NotifyEntry) error
}

type messageGetNotifyStream struct {
	stream server.Stream
}

func (x *messageGetNotifyStream) Close() error {
	return x.stream.Close()
}

func (x *messageGetNotifyStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *messageGetNotifyStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *messageGetNotifyStream) Send(m *NotifyEntry) error {
	return x.stream.Send(m)
}