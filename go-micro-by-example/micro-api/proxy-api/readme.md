# Proxy API
在proxy代理模式下运行API，我们可以自行决定使用何种语言或类库来写接口层应用。  
API会向注册中心查询服务信息，将请求路由转向合适的后台服务上。故而我们直接使用go-web作为后天服务，因为它可以直接注册，为了方便我们不直接从头写可以注册的服务。  

## 使用方法
以http模式运行API
```bash
micro api --handler=http
```
运行代理应用
```bash
go run proxy.go
```

## 示例1 `/example/call`
发送请求到 /example/call，该请求会被API反向代理到go.micro.api.example服务的 /example/call路由  
```bash
curl "http://localhost:8080/example/call?name=micro"
```

## 示例2 `/example/foo/bar`
POST请求到 /example/foo/bar会调用go.micro.api.example的 /example/foo/bar路由。  
```bash
 curl -H 'Content-Type: application/json' -d '{"name": "micro"}' http://localhost:8080/example/foo/bar curl -H 'Content-Type: application/json' -d '{"name": "micro"}' http://localhost:8080/example/foo/bar
```
 
## 示例3 文件上传 `/example/foo/upload`
我们可以请求`http://localhost:8080/example/foo/upload`，
获取上传页面，选择适当的文件上传，测试上传功能。
为了方便和直观，请保证上传保存的目录存在，且上传小文件。
 
### 链接
[目录](/go-micro-by-example)  
[上一节](/go-micro-by-example/micro-api/api)  
[下一节](/go-micro-by-example/micro-api/web-api)