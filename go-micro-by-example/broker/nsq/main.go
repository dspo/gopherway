package main

import (
	"fmt"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/nsq"
	"log"
	"time"
)

var (
	topic = "mu.micro.book.topic.payment.done"
	b     = nsq.NewBroker()
)

func pub() {
	tick := time.NewTicker(time.Second)
	i := 0
	for range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		if err := b.Publish(topic, msg); err != nil {
			log.Println("[pub] 发布消息失败：%v", err)
		} else {
			fmt.Println("[pub] 发布消息：", string(msg.Body))
		}
		i++
	}
}

func sub() {
	_, err := b.Subscribe(topic, func(event broker.Event) error {
		fmt.Println("[sub] Received Body: %s, Header: %s",
			string(event.Message().Body), event.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	if err := b.Init(); err != nil {
		log.Fatalf("Broker 初始化错误：%v", err)
	}
	if err := b.Connect(); err != nil {
		log.Fatalf("Broker 连接错误：%v", err)
	}

	go pub()
	go sub()

	<-time.After(time.Second * 10)
}
