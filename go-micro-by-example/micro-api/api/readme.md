# API
## Step-1 编写proto脚本
[proto/api.proto](/go-micro-by-example/micro-api/api/proto/api.proto)
```bash
protoc --go_out=. --micro_out=. proto/api.proto
```
如果api.proto中有导入的proto，可以使用 `--proto_path=import_proto_path:.` 参数。

## Step-2 编写handler和服务


## Step-2 以api模式运行API网关
以api模式运行API网关
```bash
micro api --handler=api
```
运行api服务
```bash
go run .
```

