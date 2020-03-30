# Gopher Way

## [Go Syntax](go_syntax)
> 入门编程和Go语言

Go语言基本语法
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

## [Briefer Tutorial of Data Structures & Algorithms in Go](data_structures/go-D&A-briefer-tutorial.md)
> 数据结构与算法简明教程 - Go语言描述

## [Cookbook and Practice](cookbook/readme.md)
> Go语言烹饪书 Go语言的操作细节

* [数据操作的核心: string & bytes](bytestrings)
    - 核心接口: io.Writer & io.Reader
    - [x] [bytes.Buffer](bytestrings/buffer.go) -> 处理流数据时，用bytes.Buffer而不要用[]bytes 
    - [x] [ioutil.ReadAll](bytestrings/buffer.go) -> 读取io.Reader内容的一个快捷方式
    - regex
    - format
* 时间日期操作: date and time
    - time.Sleep 一定时间后再执行后面的程序
    - time.After 要求某段程序至少要执行一定时间
    - time.Since 计算时间差， time.Now().Sub(t)的快捷方式
    - (t time.time).Sub 计算时间差
    - time.Now 计算当前时间
* type conversion
    - string & bytes
    - standard lib: strconv
    - number
* [I/O操作的核心: io.Writer & io.Reader interfaces](/cookbook/io-interface)
* 文件和文件系统 file & file system
    - directories & files
    - [x] [操作CSV文件](filesystem/csv.go)
    - [x] [解析JSON文本](filesystem/json.go)
    - working with text template & HTML templates
    - [x] [使用临时文件](/cookbook/io-interface#使用临时文件)
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
* RPC & gRPC
    - 基于TCP的[RPC服务端](rpc/jsonrpc_tcp/server)与[端客户端](rpc/jsonrpc_tcp/server)  
    - 基于HTTP的[RPC服务端](rpc/jsonrpc_http/server)与[端客户端](rpc/jsonrpc_http/server)  
* socket & web-socket
* [micro-services for applications](rpc/protobuf)
    - [基于go-micro编写一个微服务](../go-micro-by-example/greeterservice/readme.md)  
* distributed
* 测试  
    参见[algorithms中的测试文件](../algorithms)，可以了解到testing的使用。  
    - [x] testing  
    用testing.M管理测试中的依赖顺序，用testing.T进行单元测试。  
    - [x] benchmark  
    用testing.B进行性能测试，但要注意被测函数的时间开销应该是收敛的，否则无法进行性能测试。  
* data streams  
* reflex 
涉及重要接口使用；正则表达式、常用字符串处理技巧；时间日期处理；文件操作与文件系统；
channel, goroutine, context, reflex等Go语言核心特性；http客户端与服务端编程；gRPC；
微服务；分布式；测试；数据流操作等。  

#### 实践主题
* [go-micro by example](go-micro-by-example)

## [Concurrency in Go](concurrency-in-go)
### Go 语言并发编程 提要
* Go语言并发组件
* Go语言并发模型

## [network-programming with go]()

## Learning source & Action
### 入门
* 《Go程序设计语言》 [京东](https://item.jd.com/12187988.html)
* The Way to Go [在线](https://github.com/Unknwon/the-way-to-go_ZH_CN)   

### 进阶
* 幕课网 [Go语言实战流媒体视频网站](https://coding.imooc.com/learn/list/227.html)
* 幕课网 [Go实战仿百度云盘 实现企业级分布式云存储系统](https://coding.imooc.com/learn/list/323.html) 
* 幕课网 [Go语言开发分布式任务调度 轻松搞定高性能Crontab](https://coding.imooc.com/learn/list/281.html)
* astaxie [go-best-practice](https://github.com/astaxie/go-best-practice)
* astaxie [build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang)
* 《Go高级编程》 [在线](https://github.com/chai2010/advanced-go-programming-book)|[购买](https://www.epubit.com/book/detail/40090) 

## Good Articles
- [知乎社区核心业务Golang化实践](https://zhuanlan.zhihu.com/p/48039838) xlzd - 知乎  
- [知乎推荐系统的实践及重构之路](https://zhuanlan.zhihu.com/p/53130925) 孙付伟 - 知乎
- [如何掌握所有的程序语言](http://www.yinwang.org/blog-cn/2017/07/06/master-pl) 王垠   
- [对于递归有没有什么好的理解方法？](https://www.zhihu.com/question/31412436/answer/683820765) - 帅地的回答 - 知乎  
- [知乎千万级高性能长连接网关揭秘](https://zhuanlan.zhihu.com/p/66807833) [faceair](https://www.zhihu.com/people/faceair) - 知乎

## Third Libs & Frameworks & Awesome Projects
- 官方站点
    *  https://go.dev/ Go开发者 
    * https://pkg.go.dev/ Go pkg搜索 
- HTTP, Web, 微服务    
    * [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter) 流行的第三方HTTP库 
    * [bilibili/Kratos](https://github.com/bilibili/kratos) bilibili 开源微服务工具集
    * [gin-gonic/gin](https://github.com/gin-gonic/gin) Go开发工程师招聘启示提及最多的Web开发框架 
    * [astaxie/beego](github.com/astaxie/beego) <- Go语言中的Django, 一站式Web开发框架  
    * [gorilla/websocket](https://github.com/gorilla/websocket) 流行的第三方websocket库
    - [golang.org/x/net/websocket](http://golang.org/x/net/websocket) 官方维护的websocket库
    * [caddyserver/caddy](https://github.com/caddyserver/caddy) 支持HTTP/2、原生支持HTTPS的**Web服务器**  
    * [micro/go-micro](https://github.com/micro/go-micro) 微服务开发框架 | 资料: [go-micro 最佳实践](https://github.com/micro-in-cn/all-in-one), [go-micro 特性全合一教程](https://github.com/micro-in-cn/tutorials)
    * [prometheus/prometheus](https://github.com/prometheus/prometheus) 系统与时间序列数据库监控管理工具 
    * [hashicorp/consul](https://github.com/hashicorp/consul) 分布式高可用数据中心
- DevOpts
    * [CodisLabs/codis](https://github.com/CodisLabs/codis) 分布式Redis解决方案
    * [etcd-io/etcd](https://github.com/etcd-io/etcd) 高可用的分布式键值数据库
    * [kubernetes/kubernetes（k8s）](https://github.com/kubernetes/kubernetes) 基于容器技术的分布式架构领先方案
* [sciter-sdk/go-sciter](https://github.com/sciter-sdk/go-sciter) Go语言**UI库及桌面程序开发框架**

## Who is heavily using Go ?
[百度](https://talent.baidu.com/external/baidu/index.html) 
[七牛云](https://career.qiniu.com/) 
[知乎](https://app.mokahr.com/apply/zhihu/3819) 
[bilibili](https://www.bilibili.com/blackboard/join.html) 
[360](http://hr.360.cn/) 
[今日头条(字节跳动)](https://job.bytedance.com) 
饿了么 
滴滴  
积梦 
阿里巴巴 
网易  

## 面试
- [初级程序员Linux题](https://developer.aliyun.com/ask/274614?utm_content=g_1000109865)
