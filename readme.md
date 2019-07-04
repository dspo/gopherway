# Go Notes - Gopher Way

## [Go Syntax](go_syntax)
以下链接指向菜鸟教程，但不包含教程的全部内容
* [基础语法](https://www.runoob.com/go/go-basic-syntax.html)
* [数据类型](https://www.runoob.com/go/go-data-types.html)
* [变量](https://www.runoob.com/go/go-variables.html)
* [条件语句](https://www.runoob.com/go/go-decision-making.html)
* [循环语句](https://www.runoob.com/go/go-loops.html)
* [函数](https://www.runoob.com/go/go-functions.html)
* [数组](https://www.runoob.com/go/go-arrays.html) [切片](https://www.runoob.com/go/go-slice.html)
* [指针](https://www.runoob.com/go/go-pointers.html)
* [结构体](https://www.runoob.com/go/go-structures.html)
* [map](https://www.runoob.com/go/go-map.html)
* [接口](https://www.runoob.com/go/go-interfaces.html)
* [错误处理](https://www.runoob.com/go/go-error-handling.html)

## [Data structures in Go](data_structures)
![](https://img.shields.io/badge/download-89-brightgreen.svg)  
[PDF eBook - en](data_structures/go-data-structures-and-algorithms.pdf)   <- 如无英文阅读障碍，可直接阅读电子书; 如文件过大不能在线阅读可打开页面后下载到本地阅读   
* [ ] basic types
    - [ ] number
    - [ ] bool
    - [ ] byte & string
* [ ] array & slice
* [ ] map
* [ ] pointer
* [ ] struct
* [ ] interface
* [ ] list
* [ ] tree
* [ ] graph 

## [Algorithms in Go](algorithms)
* [Binary Search](algorithms/binary_search.go)  
* [Sort Algorithms](algorithms/sort_algorithms.go) Based on sort.Interface  
    - [x] Bubble Sort
    - [x] Short Bubble Sort
    - [x] Selection Sort
    - [x] Insertion Sort (see official implementation here)
    - [x] Heap Sort (see official implementation here)

## Classic Problems
* 八皇后问题
* Josephus问题
* 汉诺塔问题

## [Cookbook](cookbook)
* [string & bytes](cookbook/bytestrings)
    - [bytes.Buffer](cookbook/bytestrings/buffers.go) -> a better practice to solve stream than []bytes 
    - [ioutil.ReadAll](cookbook/bytestrings/buffers.go) -> a shortcut to read io.Reader content
    - [string()](cookbook/bytestrings/buffers.go) -> convert []byte to string
    - conversion
    - regex
    - format
* date and time
    - time.Sleep 一定时间后再执行后面的程序
    - time.After 要求某段程序至少要执行一定时间
    - time.Since 计算时间差， time.Now().Sub(t)的快捷方式
    - time.Sub 计算时间差
    - time.Now 计算当前时间
* type conversion
    - string & bytes
    - standard lib: strconv
    - number
* core I/O interfaces
    - [io.Reader & io.Writer](notes/io.Reader-and-io.Writer.md)
    - io.Closer
* file & file system
    - directories & files
    - working with CSV
    - working with JSON
    - working with temporary files
    - working with text template & HTML templates
* context in Go
* concurrency in Go
* databases & storage
    - db/sql
    - mySQL
    - postgreSQL
    - SQLite
    - Redis
    - MongoDB
* http client programing
* http server programing
* gRPC
* micro-services for applications
* distributed
* testing
    * testing
    * benchmark
* data streams
* reflex

## Practice & Advices
* go env management
* go build
* go project management

## Go in action & Learning source
### 入门
* 《Go程序设计语言》 [京东](https://item.jd.com/12187988.html)
### 进阶
* 幕课网 [Go语言实战流媒体视频网站](https://coding.imooc.com/learn/list/227.html)
* 幕课网 [Go实战仿百度云盘 实现企业级分布式云存储系统](https://coding.imooc.com/learn/list/323.html) 
* 幕课网 [Go语言开发分布式任务调度 轻松搞定高性能Crontab](https://coding.imooc.com/learn/list/281.html)
* astaxie [go-best-practice](https://github.com/astaxie/go-best-practice)
* astaxie [build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang)
* 《Go高级编程》 [在线](https://github.com/chai2010/advanced-go-programming-book) [购买](https://www.epubit.com/book/detail/40090) 

## Golang Articles
* 知乎专栏文章 [知乎社区核心业务 Golang 化实践](https://zhuanlan.zhihu.com/p/48039838) xlzd  

## Go Frameworks & Libs
* 最流行第三方HTTP库 [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
* bilibili 开源微服务框架 [bilibili/Kratos](https://github.com/bilibili/kratos)
* Go开发工程师招聘启示提及最多的Web及微服务开发框架 [gin-gonic/gin](https://github.com/gin-gonic/gin)
* 一站式Web开发框架 [astaxie/beego](github.com/astaxie/beego) <- Go语言中的Django  
* Go语言UI库及桌面程序开发框架 [sciter-sdk/go-sciter](https://github.com/sciter-sdk/go-sciter)

## Who use Go a ton in China
百度BFE 七牛云 知乎 bilibili 360 今日头条 饿了么 滴滴  
积梦 阿里巴巴 网易  
