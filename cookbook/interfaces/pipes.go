package interfaces

import (
	"io"
	"os"
)

//一个Pipe示例
func PipeExample() error {
	//io.Pipe()返回一个互相关联的PipeReader和PipeWriter
	//(它们分别实现了io.Reader和io.Writer接口),
	//形成一个内存中的管道，
	//对PipeWriter的写操作可以从PipeReader中读取。
	//这样最大的目的在于，在将不同源数据写到同一流中时，可以立即读取到。
	//在链接输入输出流的场景中，这套接口提供了非常有用的线程安全的的操作方法，
	//但许多对Go语言面向接口编程思想和io库不熟悉的开发者，往往浪费了这套设计。
	r, w := io.Pipe()
	go func() {
		w.Write([]byte("test"))
		w.Close()
	}()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}