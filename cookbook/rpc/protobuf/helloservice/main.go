package main

import (
	"github.com/micro/go-micro"
	"github.com/prometheus/common/log"
)

var SERVICE_NAME  = "go.micro.service.user"

func main()  {
	service := micro.NewService(
		micro.Name(SERVICE_NAME),
	)
	service.Init()
	_ = RegisterHelloServiceHandler(service.Server(), new(TheManSayHello))
	if err := service.Run(); err != nil {
		log.Fatal("error when service start run ")
	}
}
