package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var m map[int]int
	fmt.Println(&m, unsafe.Sizeof(m))
	m = make(map[int]int)
	m[1] = 100 //注意引用类型 指针所指向的内存单位需要手动开辟出来
	fmt.Println(m)

	var arr []int
	fmt.Println(unsafe.Pointer(&arr), unsafe.Sizeof(arr), len(arr), cap(arr))
	arr = make([]int, 5, 5) //是否开辟空间.是否len 直接决定 [] 这种方式是否允许调用
	// arr = append(arr, 100)
	arr[0] = 100
	fmt.Println(arr)
}
