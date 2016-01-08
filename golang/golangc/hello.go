package main

// #include "hello.h"
import "C"
import "fmt"
import "unsafe"

func main() {

	C.sayhello()
	i := 100
	ip := unsafe.Pointer(&i)
	C.add((*C.int)(ip))
	fmt.Println(i)
	C.say((*C.int)(ip))
	fmt.Println(i)
}
