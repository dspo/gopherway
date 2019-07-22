package main

import (
	"./handler"
	"./proto"
	"context"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"os"
)

func runClient(service micro.Service)  {
	greeter := proto.NewGreeterService("greeter", service.Client())
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name:"tony stack"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Greeting)
}

func main() {
	//create a new service.
	//optionally include some options here
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{"type": "helloworld",}),

		//setup some flags.
		//specify --run_client to run the client
		//add runtime flags
		micro.Flags(cli.BoolFlag{
			Name:"run_client",
			Usage:"launch the client",
		}),
	)

	service.Init(
		//init will parse the command line flags.
		// any flags set will override above settings.
		// options defined here will override anything set on the command line.
		micro.Action(func(context *cli.Context) {
			if context.Bool("run_client"){
				runClient(service)
				os.Exit(0)
			}

		}))

	//by default we'll run the server unless the flags catch us
	//setup the server
	//register handler
	proto.RegisterGreeterHandler(service.Server(), new(handler.Greeter))

	//run the server
	if err := service.Run(); err !=nil{
		fmt.Println(err)
	}
}
