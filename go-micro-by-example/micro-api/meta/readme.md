# Meta API

������ʾ���ʹ��metadataģʽ�µ�Micro API�����¼��**API**��go-micro֧�ֽ�����·�ɵ�����Ԫ���������ķ�����Ҳ���ǻ���Ԫ���ݵķ����֡�

## ʹ�÷���
### Step-1 ��дproto�ű������ɴ���
ʹ��protoc����go����
```
protoc --go_out=. --micro_out=. proto/api.proto
```

### Step-2 ��������
����API���أ����Կ�����**API**����ʱ����û������handlerģʽ���ʶ�ʹ�õ�**RPC**ģʽ������**Meta API**��ʵ����RPCģʽ�Ļ����ϣ�ͨ���ڽӿڲ������˵�Ԫ���ݶ�ָ������ġ�
```
micro api
```
### Step-3 ����ʾ������
����ʾ�������ڴ�����ע�����ʱ��������endpoint������д����Ԫ���ݣ������ӿ�Ϊ **/example**�� **/foo/bar**
```
go run meta.go
```

### Step-4 �ͻ��˷�������
�� **/example** POST����ʱ���ᱻת��**go.micro.api.example**��**Example.Call**������
```
curl -H 'Content-Type: application/json' -d '{"name": "john"}' "http://localhost:8080/example"
```

�� **/example** POST����ʱ���ᱻת��**go.micro.api.example**��**Foo.Bar**������
```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/foo/bar
```
### ����
[Ŀ¼](/go-micro-by-example)  
[��һ��](/go-micro-by-example/micro-api/event)  
