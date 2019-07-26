package handler

import (
	"../proto"
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/errors"
	"log"
	"strings"
)

type Example struct {
}

//Example.Call通过API向外暴露/example/call，接收请求
//--> /example/call  ----调用----->   go.micro.api.example服务的Example.Call方法
func (e *Example) Call(ctx context.Context, req *proto.Request, rsp *proto.Resposne) error {
	log.Println("Example.Call接口收到请求")
	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.example", "参数不正确")
	}

	//打印请求头
	for k, v := range req.Header {
		log.Println("请求头", k, ": ", v)
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": "我们已经收到你的请求，" + strings.Join(name.Values, " "),
	})
	rsp.Body = string(b)
	return nil

}

type Foo struct {
}

//Bar 方法全称是Foo.Bar，故而它会以/example/foo/bar为路由提供服务
func (f *Foo) Bar(ctx context.Context, req *proto.Request, rsp *proto.Resposne) error {
	log.Println("Foo.Bar接口接收到请求")
	if req.Method != "POST" {
		return errors.BadRequest("go.micro.api.example", "require post")
	}
	ct, ok := req.Header["Content-Type"]
	fmt.Println(ct)
	if !ok || len(ct.Values) == 0 {
		return errors.BadRequest("go.micro.api.example", "need content-type")
	}
	contentType := ct.Values[0]
	if contentType != "application/json" {
		return errors.BadRequest("go.micro.api.example", "except application/json but " + contentType)
	}

	log.Println("set body")
	var body map[string]interface{}
	json.Unmarshal([]byte(req.Body), &body)
	log.Printf("req.Body values: %s type: %T", req.Body, req.Body)

	//设置返回值
	rsp.Body = "收到消息：" + string([]byte(req.Body))
	return nil
}
