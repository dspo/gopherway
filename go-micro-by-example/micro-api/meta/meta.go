package main

import (
	"./proto"
	"context"
	"github.com/micro.bak/go-api"
	"github.com/micro.bak/go-api/handler/rpc"
	"github.com/micro.bak/go-micro/errors"
	"github.com/micro/go-micro"
	rapi "github.com/micro/go-micro/api/handler/api"
	"log"
)

type Example struct {
}

type Foo struct {
}

//Call方法在下面main中我们通过endpoint将其注册到/example/call路由
func (e *Example) Call(ctx context.Context, req *proto.CallRequest,
	rsp *proto.CallResponse) error {
	log.Println("Meta Example.Call接收到请求")
	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.api.example", "no content")
	}
	rsp.Message = "Meta已经收到你的请求，" + req.Name
	return nil
}

//Bar方法在下面main中我们通过endpoint将其注册到/example/foo/bar路由
func (f *Foo) Bar(ctx context.Context, req *proto.EmptyRequest, rsp *proto.EmptyResponse) error {
	log.Println("Meta Foo.Bar接口收到请求")
	return nil
}

func main() {
	service := micro.NewService(micro.Name("go.micro.api.example"))
	service.Init()

	//注册Example接口处理器
	proto.RegisterExampleHandler(service.Server(), new(Example),
		api.WithEndpoint(&api.Endpoint{
			Name:        "Example.Call",   //接口方法名，一定proto中存在，不能是类的自有方法
			Description: "Example.Call",   //
			Handler:     rpc.Handler,      //该接口的API转发模式
			Method:      []string{"POST"}, //支持的方法类型
			Path:        []string{"POST"}, //http请求路由，支持POSIX风格
		}))

	//注册Foo接口处理器
	proto.RegisterFooHandler(service.Server(), new(Foo), api.WithEndpoint(
		&api.Endpoint{
			Name:        "Foo.Bar",
			Description: "Foo.Bar",
			Handler:     rapi.Handler,
			Method:      []string{"POST"},
			Path:        []string{"/foo/bar"},
		}))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
