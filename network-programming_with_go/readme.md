# 网络编程 Go语言描述

## TCP client
> 原文p44
一旦一个客户端通过一个TCP地址连接到服务, 就会"dial"这个服务.
 如果成功了, 这个dial就会返回一个TCPConn的会话.
 客户端与服务端通过这个TCPConn交换消息.
 一般来说客户端向这个TCPConn写入消息以发送数据, 从这个TCPConn读取消息以接收来自服务端的响应.
 直到一方或双方中断了连接, 这个读写的过程才会结束.
 Go语言中, 可以通过以下函数建立一个客户端TCP连接:
```go
func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Eerror)
```
其中`laddr`是本地地址,一般设为nil, `raddr`是服务端的地址, `net`可以是`tcp`, `tcp4`, `tcp6`中的一个,
 具体选择哪个取决于你想要TCPv4连接还是TCPv6连接, 如果不关心TCPv4还是TCPv6, 就直接写`tcp`.
 
HTTP服务和客户端就是TCP服务和客户端的一个简单的例子. 见原文p45 .

## A daytime server
原文p46, [代码](network-programming_with_go/example/p47)

## Multi-threaded server
原文p48, 

## Controlling TCP connections

### timeout
```go
func (c *TCPConn) SetTimeout(nsec int64) os.Error
``` 

### staying alive
```go
func (c *TCPConn) SetKeepAlive(keepalive bool) os.Error
```

