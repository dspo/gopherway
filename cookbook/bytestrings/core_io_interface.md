# I/O 操作的核心接口 io.Writer, io.Reader

io.Writer和io.Reader是Go语言中数据操作的核心接口,
 许多常见的类型都通过实现io.Writer, io.Reader进行数据交换.
 
实现io.Writer的有如http.ResponseWriter, os.Stdout等.
 http.ResponseWriter本身也是一个interface, 它包含了io.Writer接口,
 通过.Write([]byte)方法, 可以将数据写入, 从而响应给HTTP的请求端.
 os.Stdout是Go语言中的标准输出流, 通过.Write([]byte)写入数据,
 从而将数据"打印"到终端界面.
 如:
```go
func TestStdout(t *testing.T){
	os.Stdout.Write([]byte("test Stdout print"))
}
```
 fmt.Print()簇函数也是调用os.Stdout.Write()方法进行打印.
 
*http.Request, *os.File等都实现了io.Reader. 通过.Read(p []byte) (n int, err error)
方法, 可以读取HTTP请求端传入的数据, 可以读取文件中的数据.

## io.Writer
type io.Writer interface封装了基本的Write方法.
 io.Writer通过Write(p []byte)(n int, err error)方法, 向某个潜在的对象写入长度为len(p)的数据p.
 返回的n是写入的数据长度, n <= len(p).
 返回的error是造成写入失败的错误. 如果 n < len(p), 那么error必定不为nil.
 方法不能修改p, 哪怕临时修改也不行.
 
 io.Writer的实现不能保留p.

io.Writer的源码如下:
```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

## io.Reader
type io.Reader interface封装了基本的Read方法.
 Read(p []byte) (n int, err error)方法从读取len(p)的bytes到p中.
 返回的n是读取的bytes的长度, 0 <= n <= len(p).
 返回的err返回读取中发生的错误.
 即使返回的n < len(p), 在方法的调用过程中, 也可能使用了p的所有空间.
 如果可以读取到数据, 但数据长度不是len(p), Read方法也会返回所能取得的数据.
 
如果Read方法成功读取到了n>0 的字节数, 而此时到达了文末(end-of-file, EOF),
 它会返回读取到的字节数量. 返回的err可能是nil, 也可能是EOF. 但下一次读取时, 返回的肯定是(0, EOF).
 调用者应当处理n>0的bytes, 再处理err. 这样做可以正确地处理读取部分数据后发生的错误,
 也能处理EOF事件.
 
 一般来说, 当n==0, err要么为EOF, 要么为non-nil.
 如果返回值为(0, nil), 则认为操作没有被执行(可能是读取数据时发生了一些阻碍).
 
io.Reader的源码:
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

## 实现一个io.Reader
> cookbook/bytestrings/core_io_interface.md
```go
package bytestrings

import (
	"io"
)

type IOReaderImp struct {
	Data []byte
}

func (i *IOReaderImp) Read(p []byte) (n int, err error) {
	lenData := len(i.Data)
	if lenData == 0 {
		return 0, io.EOF
	}

	lenp := len(p)
	if lenData <= lenp {
		n := copy(p, i.Data)
		i.Data = []byte{}
		return n, io.EOF
	}

	p, i.Data = i.Data[:len(p)], i.Data[len(p):]
	return lenp, nil
}
```