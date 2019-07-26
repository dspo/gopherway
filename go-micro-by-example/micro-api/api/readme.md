# API
## Step-1 ��дproto�ű�
* [proto/api.proto](/go-micro-by-example/micro-api/api/proto/api.proto)
```bash
protoc --go_out=. --micro_out=. proto/api.proto
```
* [proto/api.micro.go](/go-micro-by-example/micro-api/api/proto/api.micro.go)
* [proto/api.pb.go](/go-micro-by-example/micro-api/api/proto/api.pb.go)
���api.proto���е����proto������ʹ�� `--proto_path=import_proto_path:.` ������

## Step-2 ��дhandler�ͷ���
* [handler/api.go](/go-micro-by-example/micro-api/api/handler/api.go)
* [main.go](/go-micro-by-example/micro-api/api/main.go)

## Step-3 ��apiģʽ����API����
��apiģʽ����API����
```bash
micro api --handler=api
```
����api����
```bash
go run .
```

<t> TODO: /example/foo/bar�ӿ�û����ͨ�� --> [handler/api.go](/go-micro-by-example/micro-api/api/handler/api.go)

