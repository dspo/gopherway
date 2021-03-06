# Gopher熟练工千千问 - 语法篇

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
   
0. - 阅读如下代码, 回答问题
    ```go
    package main
    
    import (
        "fmt"
    )
    
    type ReaderA interface {
        Read([]byte)(int, error)
    }
    
    type ReaderB = interface {
        Read([]byte)(int, error)
    }
    
    type ReaderC = interface {
        Read([]byte)(int, error)
    }
    
    type AMaker interface {
        MakeReader() ReaderA
    }
    
    type BMaker interface {
        MakeReader() ReaderB
    }
    
    type CMaker interface {
        MakeReader() ReaderC
    }
    
    type readerC struct {
    
    }
    
    func (r *readerC) Read([]byte) (int, error) {
        return 0, nil
    }
    
    type cMaker struct {
    
    }
    
    func (r *cMaker) MakeReader() ReaderC {
        return new(readerC)
    }
    
    func f1() {
        var b BMaker
        b = new(cMaker)
        _, okA := b.(AMaker)
        _, okC := b.(CMaker)
        fmt.Println(okA, okC)
    }
    
    func f2() {
        var c CMaker
        c = new(cMaker)
        _, okA := c.(AMaker)
        _, okB := c.(BMaker)
        fmt.Println(okA, okB)
    }
    
    func main() {
        f1()
        f2()
    }
    
    ```
    如上代码中是否有错误, 如有错误, 请问是编译错误还是运行时错误?
    如无错误, 请问打印的结果是?
    注: 请独立解答, 不要借助智能IDE等工具检查代码.