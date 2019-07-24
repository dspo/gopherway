package main

// declare C program in comments
/*
#includ <stdio.h>

static void SayHello(const char* s) {
	puts(s);
}
*/

import "C"

func main() {
	C.SayHello(C.CString("Hello, World\n"))  // we implement the SayHello in comments
}
