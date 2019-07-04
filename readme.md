# Go Notes - Gopher Way

## [Go Syntax](go_syntax)
* var
* value
* types
    - int
    - float
    - byte
    - string
    - array
    - slice
    - []bytes
    - pointer
    - map
    - struct
    - interface
* conditions and branches
* loop
    - for [index]
    - for [condition]
    - switch-case
* function
* struct
    - method
* error
* goroutine
* channel
    - select

## [Data structures in Go](data_structures)
* basic types
    - number
    - bool
    - byte & string
* array & slice
* map
* pointer
* struct
* interface
* list
* tree
* graph 

## [Algorithms in Go](algorithms)
* [Binary Search](algorithms/binary_search.go)  
* [Sort Algorithms](algorithms/sort_algorithms.go) Based on sort.Interface  
    - Bubble Sort
    - Short Bubble Sort
    - Selection Sort
    - Insertion Sort (see official implementation here)
    - Heap Sort (see official implementation here)

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
## Practice
* go env management
* go build
* go project management
## go in action
### stream video server web application
### distributed storage web application
### cron tab by go