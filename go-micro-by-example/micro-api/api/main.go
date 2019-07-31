package main

import (
	"github.com/micro/go-micro"
	"github.com/prometheus/common/log"
	"gopherway/go-micro-by-example/micro-api/api/handler"
	"gopherway/go-micro-by-example/micro-api/api/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)
	service.Init()

	proto.RegisterExampleHandler(service.Server(), new(handler.Example))
	proto.RegisterFooHandler(service.Server(), new(handler.Foo))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
