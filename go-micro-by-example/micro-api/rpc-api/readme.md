# RPC API
��ģʽ����������ͨ��RPC�ķ�ʽ��HTTP����ת����go-micro΢�����ϡ�  
RPCģʽ��APIֻ����POST��ʽ�����󣬲���ֻ֧�����ݸ�ʽcontent-typeΪapplication/json����application/protobuf��  

## Step-1 ��дprotoЭ��ű�
1. ��дprotoЭ��ű�  
    [rpc-api/proto/api.proto](proto/api.proto)
0. ����go����
    ```proto
    protoc --go_out=. --micro_out=. proto/api.proto
    ```
    [rpc-api/proto/api.pb.go](proto/api.pb.go)  
    [rpc-api/proto/api.micro.go](proto/api.micro.go)  

## Step-2 ��дhandler
[rpc-api/handler/rpc.go](handler/rpc.go)  

## Step-3 ��дservice
[rpc-api/main.go](main.go)  

## Step-4 ��rpcģʽ����API
�ֱ�����  
```bash
micro api --handler=rpc
```
```bash
go run rpc.go
```

## Step-5 �ӿڲ���
[test.py](test.py)  

### ����
[Ŀ¼](/go-micro-by-example)  
[��һ��](/go-micro-by-example/micro-api)  
[��һ��](/go-micro-by-example/micro-api/api)