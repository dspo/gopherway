# Gopher千千问 - 网络篇

1. - 请简述TCP三次握手, 四次挥手.
0. - 请简述TCP状态机, 可以画图解释.
0. - 阅读如下代码, 回答问题
    ```go
    import (
     "fmt"
     "net"
     "io"
     "log"
    )
   
    func ListenAndServe(address string, handle func(conn net.Conn)) {
        listener, err := net.Listen("tcp", address)
        if err != nil {
            log.Fatal(fmt.Sprintf("listen err: %v", err))
        }
        defer listener.Close()
    
        for {
            conn, err := listener.Accept() // (1)
            if err != nil {
                log.Fatal(fmt.Sprintf("accept err: %v", err))
            }
            go handle(conn)
        }
    }
    ```
    代码中注释(1)的行, 其中的`conn`对应TCP的什么状态?