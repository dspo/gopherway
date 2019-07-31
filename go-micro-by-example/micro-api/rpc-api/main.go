package main

import (
	"github.com/micro/go-micro"
	"github.com/prometheus/common/log"
	"gopherway/go-micro-by-example/micro-api/rpc-api/handler"
	"gopherway/go-micro-by-example/micro-api/rpc-api/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)
	service.Init()

	//注册example接口
	proto.RegisterExampleHandler(service.Server(), new(handler.Exampel))
	//注册foo接口
	proto.RegisterFooHandler(service.Server(), new(handler.Foo))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
