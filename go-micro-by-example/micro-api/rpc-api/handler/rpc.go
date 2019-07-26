package handler

import (
	"context"
	"github.com/micro/go-micro/errors"
	"log"
	"../proto"
)

type Exampel struct {

}

//Example.Call方法会由API层转发，路由为/example/call的HTTP请求
func (e *Exampel) Call(ctx context.Context, req *proto.CallRequest, rsp *proto.CallResponse) error {
	log.Printf("copy the requests of Example.Call")
	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.api.example", "no context")
	}

	rsp.Message = "RPC Call copies your request 我接到你的请求啦，" + req.Name
	return nil
}

type Foo struct {

}

func (f *Foo) Bar(ctx context.Context, req *proto.EmptyRequest, rsp *proto.EmptyResponse) error {
	log.Println("copy the request of Foo.Bar")
	return nil
}

