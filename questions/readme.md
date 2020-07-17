# Go开发工程师千千问

## 语法篇
1. - 阅读如下代码, 回答问题
    ```go
    package main
    
    type Naming interface {
        SetName(name string)
    }
    
    type Person struct {
        name string
    }
    
    func (p *Person) SetName(name string) {
        p.name = name
    }
    
    type Student struct {
        Person
    }
    
    func main() {
        var (
            naming Naming
            student Student
        )
        naming = student
        naming.SetName("some name")
    }
    ```
    请说明上述代码中是否存在问题或错误, 如无错误, 请说明代码含义; 如有错误, 请指出是编译错误还是运行时错误, 并以最少改动修改为正确代码.
    (注意请不要借助IDE等工具检查代码) 
    
0. - 阅读如下代码, 回答问题
    ```go
    package main
    
    import (
        "fmt"
    )
    
    type Naming interface {
        SetName(name string)
        GetName() string
    }
    
    type Person struct {
        name string
    }
    
    func (p *Person) SetName(name string) {
        p.name = name
    }
    
    func (p *Person) GetName() string {
        return p.name
    }
    
    type Student struct {
        *Person
    }
    
    func main() {
        var (
            naming Naming
            student Student
        )
        naming = student // (1)
        // naming = &student // (2)
        naming.SetName("some name")
        fmt.Println(naming.GetName())
    }
    ```    
    上述代码中是否存在问题, 如从存在, 是编译问题还是运行时问题? 如将(1)处代码改为(2)处呢?

0. - Go语言中的`slice`类型和`map`类型是并发安全的吗? 如何理解Go语言中的"并发安全"?

0. - Go语言中的`goroutine`是什么, 请简述你的理解?

0. - 协程是什么? 它的原理是什么? 它与线程有什么不同?

0. - 请用Go语言实现一个方法或函数, 对有序`slice`实现原地去重. 可以假设这个`slice`是非严格递增递增的`[]int` .
    ```go
    func TrimRepeat(nums []int) []int {}
    ```

   
## 网络篇
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
    
### 数据结构与算法篇
1. - 什么是哈希表? 哈希表的原理是什么? 如何用Go实现一个哈希表?

0. - 大部分语言中有键值对存取的数据结构, 如Go中的`map`, Python中的`dict`等. 
    但也有语言原生并不支持这种结构, 如果让你设计键值对存取结构, 你会如何设计?
    更进一步的, 你是否能用Go语言实现一个简易版的`map` ?

0. - 请简述二叉树的概念, 性质.

0. - 请简述二叉树的前序, 中序, 后序遍历, 并用Go语言实现之.     
    
    