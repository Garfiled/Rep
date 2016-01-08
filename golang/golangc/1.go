package main

// #include <stdio.h>
// #include <stdlib.h>
/*
void print() {
    printf("hello");
}
*/
import "C"

func main() {
	//	s := "Hello Cgo"
	//	cs := C.CString(s)
	C.print()
}
