# Event

��ƪ��ʾ���ʹ��Eventģʽ�µ�Micro API��  
API���http����ӳ�䵽ƥ���ʱ�䴦������ϡ�  

# ����
��Ϊ�����ڴ������������¼�����topic��`go.micro.evt.user`������˵�¼���������������ռ���
`go.micro.evt`���������ǵ�APIҲҪ����������ռ䣬����API�����ҵ�����  
```bash
micro api --handler=event --namespace=go.micro.evt
```

���з���
```bash
go run main.go
```

�����¼�
```bash
curl -d '{"message": "hello, ���}' http://localhost:8080/user/login
```

### ����
[Ŀ¼](/go-micro-by-example)  
[��һ��](/go-micro-by-example/micro-api/web-api)  
[��һ��](/go-micro-by-example/micro-api/meta)