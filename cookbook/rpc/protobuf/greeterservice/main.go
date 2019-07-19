package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/prometheus/common/log"
)

var (
	//service = micro.NewService(
	//	micro.Name("greeter"),
	//	micro.Version("latest"),
	//)
	service = micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:"environment",
				Usage:"the environment",
			},
		),
	)
)

func main()  {
	service.Init(
		micro.Action(
			func(context *cli.Context) {
				env := context.GlobalString("environment")
				if env == "" {
					fmt.Print("environment set to ", env)
				}
			}),
	)

	RegisterGreeterHandler(service.Server(), new(Greeter))
	log.Fatal(service.Run())
}





