//https://github.com/chai2010/advanced-go-programming-book/blob/master/ch2-cgo/ch2-01-hello-cgo.md
package main

/*
#include <stdio.h>
 */

import "C"

func main() {
	C.puts(C.CString("hello cgo")) //没有在程序退出前释放C.CString创建的C语言字符串,会导致内存泄漏
}

