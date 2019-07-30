package main

import (
	"context"
	proto "github.com/micro/go-api/proto"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
)

//切记：事件订阅结构的所有公有方法都会被执行
//方法名没有限制，但是方法一定要接收ctx, event
type Event struct {

}

func (e *Event) Process(ctx context.Context, event *proto.Event) error {
	log.Log("共有方法Process 收到事件，", event.name)
	log.Log("共有方法Process 数据", event.Data)
	return nil
}

func (e *Event) Process2(ctx context.Context, event *proto.Event) error {
	log.Log("共有方法Process2 收到事件，", event.name")
	log.Log("共有方法Process2 数据", event.Data)
	return nil
}

func (e *Event) process3(ctx context.Context, event *proto.Event) error {
	log.Log("私有方法process3, 收到事件，", event.Name)
	log.Log("私有方法process, 数据", event.Data)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("user2"),
		)
	service.Init()

	//register subscriber
	micro.RegisterSubscriber("go.micro.evt.user", service.Server(), new(Event))
	if err := service.Run(); err != nil{
		log.Fatal(err)
	}
}
