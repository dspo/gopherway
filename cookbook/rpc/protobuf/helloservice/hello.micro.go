// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: hello.proto

package main

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for HelloService service

type HelloService interface {
	Hello(ctx context.Context, in *String, opts ...client.CallOption) (*String, error)
}

type helloService struct {
	c    client.Client
	name string
}

func NewHelloService(name string, c client.Client) HelloService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "main"
	}
	return &helloService{
		c:    c,
		name: name,
	}
}

func (c *helloService) Hello(ctx context.Context, in *String, opts ...client.CallOption) (*String, error) {
	req := c.c.NewRequest(c.name, "HelloService.Hello", in)
	out := new(String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HelloService service

type HelloServiceHandler interface {
	Hello(context.Context, *String, *String) error
}

func RegisterHelloServiceHandler(s server.Server, hdlr HelloServiceHandler, opts ...server.HandlerOption) error {
	type helloService interface {
		Hello(ctx context.Context, in *String, out *String) error
	}
	type HelloService struct {
		helloService
	}
	h := &helloServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&HelloService{h}, opts...))
}

type helloServiceHandler struct {
	HelloServiceHandler
}

func (h *helloServiceHandler) Hello(ctx context.Context, in *String, out *String) error {
	return h.HelloServiceHandler.Hello(ctx, in, out)
}