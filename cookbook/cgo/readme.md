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
    `.h`文件是API（区别于go中的interface）实现者与使用者的约定，
    实现者要确保函数实现满足对外接口。参见[04/hello/hello.h](./04/hello/hello.h)  
    至于实现者使用C还是C++还是Go还是什么其他语言来实现这个接口，使用者是不知道、不理会的。   
    * 接口功能的C实现：[04/hello/hello.c](04/hello/hello/hello.c)
    * 接口功能的C++实现：[04/hello/hello.cpp](04/hello/hello/hello.cpp)
    * 接口功能的Go实现：[04/hello/hello.go](04/hello/hello/hello.go)
0. C-Interface Oriented Programming in Go (CIOPG)
    >Go面向C接口编程  
    
    * 将接口约定、接口的实现、接口的调用放到一个Go文件中。
    [04_2/hello.go](./04_2/hello.go)
    * 使用Go中的预定义C类型，使得代码更加gopheric.
    [04_3/hello.go](./04_3/hello.go)  
    
    