# cgo by example
>see [chai2010/advanced-go-programming-book](https://github.com/chai2010/advanced-go-programming-book)
1. the first cgo program  
    [01/hello.go](./01/hello.go)
0. use C function in go file
    [02/hello.go](./02/hello.go)
0. declare C function in c file  
    [03/hello.c](./03/hello.c)  
    [03/hello.go](./03/hello.go)    
0. C code Modularization  
    `.h`�ļ���API��������go�е�interface��ʵ������ʹ���ߵ�Լ����
    ʵ����Ҫȷ������ʵ���������ӿڡ��μ�[04/hello/hello.h](./04/hello/hello.h)  
    ����ʵ����ʹ��C����C++����Go����ʲô����������ʵ������ӿڣ�ʹ�����ǲ�֪���������ġ�   
    * �ӿڹ��ܵ�Cʵ�֣�[04/hello/hello.c](04/hello/hello/hello.c)
    * �ӿڹ��ܵ�C++ʵ�֣�[04/hello/hello.cpp](04/hello/hello/hello.cpp)
    * �ӿڹ��ܵ�Goʵ�֣�[04/hello/hello.go](04/hello/hello/hello.go)
0. C-Interface Oriented Programming in Go (CIOPG)
    >Go����C�ӿڱ��  
    
    * ���ӿ�Լ�����ӿڵ�ʵ�֡��ӿڵĵ��÷ŵ�һ��Go�ļ��С�
    [04_2/hello.go](./04_2/hello.go)
    * ʹ��Go�е�Ԥ����C���ͣ�ʹ�ô������gopheric.
    [04_3/hello.go](./04_3/hello.go)  
    
    