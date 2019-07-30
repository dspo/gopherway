# Event

本篇演示如何使用Event模式下的Micro API。  
API会把http请求映射到匹配的时间处理服务上。  

# 运行
因为我们在代码中声明了事件主题topic是`go.micro.evt.user`，即是说事件服务的命名所属空间是
`go.micro.evt`，所以我们的API也要是这个命名空间，这样API才能找到它。  
```bash
micro api --handler=event --namespace=go.micro.evt
```

运行服务
```bash
go run main.go
```

发送事件
```bash
curl -d '{"message": "hello, 悟空}' http://localhost:8080/user/login
```

### 链接
[目录](/go-micro-by-example)  
[上一节](/go-micro-by-example/micro-api/web-api)  
[下一节](/go-micro-by-example/micro-api/meta)