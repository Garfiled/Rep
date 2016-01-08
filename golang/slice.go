package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := make([]int, 0, 5)
	fmt.Println(s, len(s), cap(s))
	fmt.Printf("%p %p\n", &s, *(**int)(unsafe.Pointer(&s)))
	for i := 0; i < cap(s); i++ {
		s = append(s, i*100)
	}
	fmt.Println(s, len(s), cap(s))
	fmt.Printf("%p %p\n", &s, *(**int)(unsafe.Pointer(&s)))

	s = append(s, 500) //重新分配底层数组，与原数组无关
	fmt.Println(s, len(s), cap(s))
	fmt.Printf("%p %p\n", &s, *(**int)(unsafe.Pointer(&s)))
}
