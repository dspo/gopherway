# Web API

��ƪ��ʾ���ʹ��Webģʽ�µ�Micro API�����¼��API��

��web����ģʽ������API�����ǿ��Խ�API��Ϊ��������ṩ���η����http�ӿڣ�ʾ����������ʾʹ��websocket��

API����ע�����Ĳ�ѯ������Ϣ��������·��ת����ʵĺ�̨�����ϡ��ʶ�����ֱ��ʹ��go-web��Ϊ��̨������Ϊ������ֱ��ע�ᣬΪ�˷������ǲ�ֱ�Ӵ�ͷд����ע��ķ���

## ʹ�÷���
����consul
```bash
concult agent -dev
```

��webģʽ����API����Ϊ���ǵ�Ӧ������web�ռ��£��������ǰ�api�������ռ�����Ϊgo.micro.web  
```bash
micro api --handler=web --namespace=go.micro.web
```
����web-serverӦ�ã����web-server�Ҳ���consul����ʹ�� `--registry_address` ָ�������ַ��
```bash
go run web-server.go --registry_address=localhost:8300
```

## ��ʾ
�� `http://127.0.0.1:8080/websocket/`

Ĭ�ϻ��websocket���ӡ���Name���������ı������send��ť��������̨websocket���񽻻���

### ����
[Ŀ¼](/go-micro-by-example)  
[��һ��](/go-micro-by-example/micro-api/proxy-api)  
[��һ��](/go-micro-by-example/micro-api/event)