# API
## Step-1 ��дproto�ű�
[proto/api.proto](/go-micro-by-example/micro-api/api/proto/api.proto)
```bash
protoc --go_out=. --micro_out=. proto/api.proto
```
���api.proto���е����proto������ʹ�� `--proto_path=import_proto_path:.` ������

## Step-2 ��дhandler�ͷ���


## Step-2 ��apiģʽ����API����
��apiģʽ����API����
```bash
micro api --handler=api
```
����api����
```bash
go run .
```

