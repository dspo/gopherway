package main

import (
	"context"
)

type Greeter struct {

}

func (g *Greeter) Hello(ctx context.Context, req *HelloRequest, rsp *HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}