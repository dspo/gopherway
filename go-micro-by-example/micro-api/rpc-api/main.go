package main

import (
	"./handler"
	"./proto"
	"github.com/micro/go-micro"
	"github.com/prometheus/common/log"
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
