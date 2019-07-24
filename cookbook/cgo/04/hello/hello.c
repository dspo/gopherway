#include "hello.h"
#include "<stdio.h>"

// by include hello.h, we make sure the functions implemented satisfy the module opening api
// hello.h is the engagement between implementor and user of the hello module
// the engagement does not require to implement in C or C++ or others
void SayHello(const char* s) {
    puts(s);
}