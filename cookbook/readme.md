# Go语言简明Cookbook

## 目录
* [string & bytes](bytestrings)
    - [x] [bytes.Buffer](bytestrings/buffer.go) -> 处理流数据时，用bytes.Buffer而不要用[]bytes 
    - [x] [ioutil.ReadAll](bytestrings/buffer.go) -> 读取io.Reader内容的一个快捷方式
    - regex
    - format
* date and time
    - time.Sleep 一定时间后再执行后面的程序
    - time.After 要求某段程序至少要执行一定时间
    - time.Since 计算时间差， time.Now().Sub(t)的快捷方式
    - (t time.time).Sub 计算时间差
    - time.Now 计算当前时间
* type conversion
    - string & bytes
    - standard lib: strconv
    - number
* [core I/O interfaces](/cookbook/io-interface)
* file & file system
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

* 模式
    - [x] [基本生产者消费者模型](design_patterns/production_and_consumer/production_and_consumer.go)  
    示例了一个利用channel构造的基本的生产者消费者模型。  
* [环境管理](/cookbook/env-management)    
