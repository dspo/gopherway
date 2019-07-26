// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/api.proto

package proto

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

// Client API for Example service

type ExampleService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Resposne, error)
}

type exampleService struct {
	c    client.Client
	name string
}

func NewExampleService(name string, c client.Client) ExampleService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proto"
	}
	return &exampleService{
		c:    c,
		name: name,
	}
}

func (c *exampleService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Resposne, error) {
	req := c.c.NewRequest(c.name, "Example.Call", in)
	out := new(Resposne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Example service

type ExampleHandler interface {
	Call(context.Context, *Request, *Resposne) error
}

func RegisterExampleHandler(s server.Server, hdlr ExampleHandler, opts ...server.HandlerOption) error {
	type example interface {
		Call(ctx context.Context, in *Request, out *Resposne) error
	}
	type Example struct {
		example
	}
	h := &exampleHandler{hdlr}
	return s.Handle(s.NewHandler(&Example{h}, opts...))
}

type exampleHandler struct {
	ExampleHandler
}

func (h *exampleHandler) Call(ctx context.Context, in *Request, out *Resposne) error {
	return h.ExampleHandler.Call(ctx, in, out)
}

// Client API for Foo service

type FooService interface {
	Bar(ctx context.Context, in *Request, opts ...client.CallOption) (*Resposne, error)
}

type fooService struct {
	c    client.Client
	name string
}

func NewFooService(name string, c client.Client) FooService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proto"
	}
	return &fooService{
		c:    c,
		name: name,
	}
}

func (c *fooService) Bar(ctx context.Context, in *Request, opts ...client.CallOption) (*Resposne, error) {
	req := c.c.NewRequest(c.name, "Foo.Bar", in)
	out := new(Resposne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Foo service

type FooHandler interface {
	Bar(context.Context, *Request, *Resposne) error
}

func RegisterFooHandler(s server.Server, hdlr FooHandler, opts ...server.HandlerOption) error {
	type foo interface {
		Bar(ctx context.Context, in *Request, out *Resposne) error
	}
	type Foo struct {
		foo
	}
	h := &fooHandler{hdlr}
	return s.Handle(s.NewHandler(&Foo{h}, opts...))
}

type fooHandler struct {
	FooHandler
}

func (h *fooHandler) Bar(ctx context.Context, in *Request, out *Resposne) error {
	return h.FooHandler.Bar(ctx, in, out)
}
