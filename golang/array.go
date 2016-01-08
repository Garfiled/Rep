package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	a := [3]int{1, 2} // 数组
	fmt.Printf("a:value=%v length=%d\n", a, len(a))
	b := [...]int{1, 2, 3, 4} // trick
	fmt.Printf("b:value=%v length=%d\n", b, len(b))
	c := [5]int{2: 100, 4: 200} // 数组初始化方式 支持索引的
	fmt.Printf("c:value=%v length=%d\n", c, len(c))
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10},
		{"user2", 20},
	}
	fmt.Printf("d:%v,%+v,%T %d\n", d, d, d, binary.Size(d))

	//多维数组
	e := [...]int{1: 100, 3: 300}
	f := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	g := [...][2]int{{1, 1}, {2, 2}, {3, 3}}
	fmt.Println(a, b, c, d, e)
	fmt.Println(f, g)
	fmt.Println("----------")

	h := [2]int{}
	fmt.Printf("a: %p\n", &h)
	fmt.Println(h)

}
