package main

import (
	"fmt"
	. "unsafe"
)

type Human struct {
	name string
	age  int
}
type Men interface {
	SayHi()
}

func (_ Human) SayHi() {
	fmt.Printf("hello!\n")
}
func main() {
	h := Human{"liu", 25}
	var i Men
	i = h // copy 了一份数据
	fmt.Println(h, Pointer(&h), i)
	i.SayHi()
	fmt.Println(Sizeof(i))
	p1 := Pointer(uintptr(Pointer(&i)) + 8) // 这里是指向data的void＊的指针的地址
	fmt.Println("debug:", p1)
	p := *(*(*Human))(p1) // (Pointer(uintptr(Pointer(&i)) + 8))
	fmt.Println(*p, Pointer(p))
	// 找到interface 的data 了
	value, ok := i.(Human)
	fmt.Println(value, ok)

	var i1 interface{}
	i1 = 1
	switch v := i1.(type) {
	case int:
		fmt.Println("int: ", v)
	case string:
		fmt.Println("string: ", v)
	case Human:
		fmt.Println("Human: ", v)
	default:
		fmt.Println("default: ", v)
	}
}
