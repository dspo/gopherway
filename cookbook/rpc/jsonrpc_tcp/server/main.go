package main

import (
	"net"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

//基于TCP协议的JsonRPC服务
//运行，开启TCP服务
//一下是在Python中连接TCP服务，进行远程调用的示例
//>>> import socket
//>>> s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
//>>> s.connect(('localhost', 1234))
//>>> s.send('{"method":"HelloService.Hello","params":["hello"],"id":1}'.encode())
//57
//>>> str(s.recv(1024), encoding='utf8')
//'{"id":1,"result":"hello:hello","error":null}\n'
//>>>
//>>>
//>>>
//>>>
//>>> s.send('{"method":"HelloService.Hello","params":["tony stack"],"id":1}'.encode())
//62
//>>> str(s.recv(1024), encoding='utf8')
//'{"id":1,"result":"hello:tony stack","error":null}\n'
func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
