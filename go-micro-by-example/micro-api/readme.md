# micro api
> ע:������������������[micro-in-cn/all-in-one](https://github.com/micro-in-cn/all-in-one/tree/master/basic-practices/micro-api)���������÷��ű�ű�ע������Ҳ�������������ݡ�  

>API�����Ͼ�һ���������أ����߱���̬·�ɡ������ֵ���������HTTP��ʽ���������ӳ�䵽����΢�����Զ����ṩ����
>
>ͨ�������֣�������õ�`�����ռ䣨Namespace��`����API���԰�����url����ӳ�䵽ƥ��������ռ����ķ���ӿڡ�
>
>��Micro��ϵ�У����񶼻����Լ��������ռ䣬��API��Ĭ�������ռ���`go.micro.api`��ͨ��������������ṩ����ӿڵ�΢�������ǻ�Ĭ�ϰ��� Micro������ǵ������ռ�����Ϊ`go.micro.api.example`����ʽ��example���Ǿ���ķ������������Ҫ�Ķ����ֵ��������ʱָ��` --namespace=`ָ������ָ��ɡ�

### handler ������
Micro APIĿǰ��5�ִ���ʽ�����Ը�������API���ó�ָ�������͡�
1.	**RPC**  
    Ĭ��ѡ�ͨ��RPC��go-microӦ��ת������ͨ��ֻ��������body��ͷ��Ϣ����װ��ֻ����POST����
0.  **API**  
	������rpc������<u>���������httpͷ��װ���´���</u>�����������󷽷���  
0.	**HTTP**��**proxy**  
	�Է������ķ�ʽʹ��API���൱�ڰ���ͨ��webӦ�ò�����API֮��
	��������api�ӿ�һ������web����  
0.	**web**  
	��http��࣬����֧��websocket��  
0.	**event**  
    ����event�¼��������͵�����
*	**meta**
	Ԫ���ݣ�ͨ���ڴ����е�����ѡ��ʹ�������е�ĳһ����������
	�ⲻ��һ��ģʽ��ֻ�Ƕ�api��rpcģʽ�µ���չʹ�÷�����  


### RPC/API���͵�����ӳ��

Micro�ڲ��н�http����·��ӳ�䵽����Ļ��ƣ�ӳ��������ͨ���±����

http·��    |    ��̨����    |    �ӿڷ���
----    |    ----    |    ----
/foo/bar    |    go.micro.api.foo    |    Foo.Bar
/foo/bar/baz    |    go.micro.api.foo    |    Bar.Baz
/foo/bar/baz/cat    |    go.micro.api.foo.bar    |    Baz.Cat

Ĭ�ϵ������ռ���**go.micro.api**��������ͨ��`--namespace`ָ���Զ��塣

����Щ���汾�ŵ�·����Ҳ����ӳ�䵽��������

����·��    |    ��̨����    |    �ӿڷ���
----    |    ----    |    ----
/foo/bar    |    go.micro.api.foo    |    Foo.Bar
/v1/foo/bar    |    go.micro.api.v1.foo    |    Foo.Bar
/v1/foo/bar/baz    |    go.micro.api.v1.foo    |    Bar.Baz
/v2/foo/bar    |    go.micro.api.v2.foo    |    Foo.Bar
/v2/foo/bar/baz    |    go.micro.api.v2.foo    |    Bar.Baz

�������ӳ������п��Կ�����**RPC/API**ģʽ�£�·����������������ᱻ��ϳ�Golang��������·��������ʣ�µĻ���������ռ�ǰ׺��ɷ����������磺

`/v1/foo/bar/baz`������`bar/baz`����ĸ��дת��`Bar.Baz`����·����ʣ�µ�`/v1/foo/`�������������ռ�ǰ׺`go.micro.api`���
`go.micro.api.v1.foo`��  

### proxy���͵�����ӳ��
�����������**API**ʱ��ָ��`--handler=http`����ô**API**��ᷴ��������󵽾���API�����ռ�ĺ�̨�����С�  
���磺
  
����·��    |    ����    |    ��̨����·��
---    |    ---    |    ---
/greeter    |    go.micro.api.greeter    |    /greeter
/greeter/:name    |    go.micro.api.greeter    |    /greeter/:name

### event���͵�����ӳ��
����**API**ʱ��ָ��`--handler=event`����ô**API**��ᷴ��������󵽾���API�����ռ�ĺ�̨�¼����ѷ����С�  
���磨�����ռ�����Ϊgo.micro.evt����  

����·��    |    ����    |    ����
---    |    ---    |    ---
/user/login    |    go.micro.evt.user    |    ����������ʾ���е�new(Event)�����й����������ҷ���Ҫ��ctx���¼�����


### ����
* [Ŀ¼](../)
* [��һ�ڣ�һ���򵥵�go-microʾ��](../greeterservice)
* [��һ�ڣ�rcp apiʾ��](../micro-api/rpc-api)
