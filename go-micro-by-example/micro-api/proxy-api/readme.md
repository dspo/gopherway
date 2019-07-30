# Proxy API
��proxy����ģʽ������API�����ǿ������о���ʹ�ú������Ի������д�ӿڲ�Ӧ�á�  
API����ע�����Ĳ�ѯ������Ϣ��������·��ת����ʵĺ�̨�����ϡ��ʶ�����ֱ��ʹ��go-web��Ϊ���������Ϊ������ֱ��ע�ᣬΪ�˷������ǲ�ֱ�Ӵ�ͷд����ע��ķ���  

## ʹ�÷���
��httpģʽ����API
```bash
micro api --handler=http
```
���д���Ӧ��
```bash
go run proxy.go
```

## ʾ��1 `/example/call`
�������� /example/call��������ᱻAPI�������go.micro.api.example����� /example/call·��  
```bash
curl "http://localhost:8080/example/call?name=micro"
```

## ʾ��2 `/example/foo/bar`
POST���� /example/foo/bar�����go.micro.api.example�� /example/foo/bar·�ɡ�  
```bash
 curl -H 'Content-Type: application/json' -d '{"name": "micro"}' http://localhost:8080/example/foo/bar curl -H 'Content-Type: application/json' -d '{"name": "micro"}' http://localhost:8080/example/foo/bar
```
 
## ʾ��3 �ļ��ϴ� `/example/foo/upload`
���ǿ�������`http://localhost:8080/example/foo/upload`��
��ȡ�ϴ�ҳ�棬ѡ���ʵ����ļ��ϴ��������ϴ����ܡ�
Ϊ�˷����ֱ�ۣ��뱣֤�ϴ������Ŀ¼���ڣ����ϴ�С�ļ���
 
### ����
[Ŀ¼](/go-micro-by-example)  
[��һ��](/go-micro-by-example/micro-api/api)  
[��һ��](/go-micro-by-example/micro-api/web-api)