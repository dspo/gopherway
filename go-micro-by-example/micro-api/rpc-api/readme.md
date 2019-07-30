# RPC API
该模式下允许我们通过RPC的方式把HTTP请求转发到go-micro微服务上。  
RPC模式下API只接收POST方式的请求，并且只支持内容格式content-type为application/json或者application/protobuf。  

## Step-1 编写proto协议脚本
1. 编写proto协议脚本  
    [rpc-api/proto/api.proto](proto/api.proto)
0. 生成go代码
    ```proto
    protoc --go_out=. --micro_out=. proto/api.proto
    ```
    [rpc-api/proto/api.pb.go](proto/api.pb.go)  
    [rpc-api/proto/api.micro.go](proto/api.micro.go)  

## Step-2 编写handler
[rpc-api/handler/rpc.go](handler/rpc.go)  

## Step-3 编写service
[rpc-api/main.go](main.go)  

## Step-4 以rpc模式运行API
分别运行  
```bash
micro api --handler=rpc
```
```bash
go run rpc.go
```

## Step-5 接口测试
[test.py](test.py)  

### 链接
[目录](/go-micro-by-example)  
[上一节](/go-micro-by-example/micro-api)  
[下一节](/go-micro-by-example/micro-api/api)