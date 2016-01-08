package main

import (
	"fmt"
	"unsafe"
)

/*
默认值nil,没有NULL常量
操作符 & 取变量地址，  * 透过指针访问目标对象
不支持指针运算，不支持 -> 运算符  直接用 . 访问目标成员
*/
func main() {
	type data struct {
		a string
		b int
	}
	var d = data{"abcd", 5678}
	var p *data

	p = &d
	fmt.Printf("%p,%d\n", p, p.b)                                            //直接用指针访问目标对象成员
	fmt.Println((*string)(unsafe.Pointer(p)), *(*string)(unsafe.Pointer(p))) //指针类型转换
	fmt.Println(unsafe.Sizeof(d), unsafe.Sizeof(d.a), unsafe.Sizeof(d.b))
	fmt.Println(unsafe.Offsetof(d.a), unsafe.Offsetof(d.b))

	p0 := uintptr(unsafe.Pointer(&d))
	p0 += unsafe.Offsetof(d.b) //指针运算

	p1 := (*int)(unsafe.Pointer(p0))
	fmt.Println(p0, *p1)
	*p1 = 1234 //指针操作
	fmt.Println(d)

}
