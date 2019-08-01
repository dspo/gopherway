# Web API

本篇演示如何使用Web模式下的Micro API，以下简称API。

在web代理模式下运行API，我们可以将API作为反向代理提供下游服务的http接口，示例中我们演示使用websocket。

API会向注册中心查询服务信息，将请求路由转向合适的后台服务上。故而我们直接使用go-web作为后台服务，因为它可以直接注册，为了方便我们不直接从头写可以注册的服务。

## 使用方法
启动consul
```bash
concult agent -dev
```

以web模式运行API，因为我们的应用是在web空间下，所以我们把api的启动空间设置为go.micro.web  
```bash
micro api --handler=web --namespace=go.micro.web
```
运行web-server应用，如果web-server找不到consul服务，使用 `--registry_address` 指定服务地址。
```bash
go run web-server.go --registry_address=localhost:8300
```

## 演示
打开 `http://127.0.0.1:8080/websocket/`

默认会打开websocket连接。在Name栏中输入文本，点击send按钮便可以与后台websocket服务交互。

### 链接
[目录](/go-micro-by-example)  
[上一节](/go-micro-by-example/micro-api/proxy-api)  
[下一节](/go-micro-by-example/micro-api/event)