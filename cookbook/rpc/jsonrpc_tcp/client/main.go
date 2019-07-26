package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//基于TCP协议的JSONRPC客户端程序
//远程调用jsonrpc_tcp/server下的函数
func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "三体", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
