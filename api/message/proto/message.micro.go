// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/message.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/message.proto

It has these top-level messages:
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import go_api "github.com/micro/go-api/proto"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = go_api.Response{}

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
	Inbox(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
	Notify(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
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
		name = "com.teddy.api.uaa"
	}
	return &messageService{
		c:    c,
		name: name,
	}
}

func (c *messageService) Inbox(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.name, "Message.Inbox", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) Notify(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.name, "Message.Notify", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Message service

type MessageHandler interface {
	Inbox(context.Context, *go_api.Request, *go_api.Response) error
	Notify(context.Context, *go_api.Request, *go_api.Response) error
}

func RegisterMessageHandler(s server.Server, hdlr MessageHandler, opts ...server.HandlerOption) error {
	type message interface {
		Inbox(ctx context.Context, in *go_api.Request, out *go_api.Response) error
		Notify(ctx context.Context, in *go_api.Request, out *go_api.Response) error
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

func (h *messageHandler) Inbox(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.MessageHandler.Inbox(ctx, in, out)
}

func (h *messageHandler) Notify(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.MessageHandler.Notify(ctx, in, out)
}