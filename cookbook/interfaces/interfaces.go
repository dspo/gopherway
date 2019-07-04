package interfaces

import (
	"io"
	"os"
	"fmt"
)

//从in复制数据到out；
//再使用一个buffer。
//数据会写到stdout
func Copy(in io.ReadSeeker, out io.Writer) error {
	w := io.MultiWriter(out, os.Stdout)
	//一种不安全的复制方式，当数据量过大时，会引发内存问题
	if _, err := io.Copy(w, in); err != nil {
		return err
	}
	in.Seek(0, 0)
	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buf); err != nil{
		return err
	}
	fmt.Println()
	return nil
}