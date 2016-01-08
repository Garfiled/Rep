package main

import (
	"fmt"
	// "unsafe"
)

func main() {
	s := "abcd你"

	for i, n := 0, len(s); i < n; i++ { //避免多次调用len
		fmt.Println(i, s[i])
	}
	fmt.Println("-----------")

	for i0 := range s {
		fmt.Println(i0, s[i0])
	}
	fmt.Println("-----------")

	for i1, c := range s {
		fmt.Println(i1, c)
	}
	fmt.Println("-----------")
	m := map[string]int{"a": 1, "b": 2}
	for i2, v := range m {
		fmt.Println(i2, v)
	}
	fmt.Println("-----------")
	a := [3]int{0, 1, 2}
	for i3, w := range a { // Range 会复制对象，index、value 都是从复制品中取出
		if i3 == 0 {
			a[1], a[2] = 999, 999
			fmt.Println(a)
		}
		fmt.Println(i3, w)
	}
	fmt.Println("-----------") //建议改用引用

	//使用引用类型，不复制底层数据
	s11 := []int{1, 2, 3, 4, 5}
	for i11, _ := range s11 {
		if i11 == 0 {
			s11[0] = 100
			s11[2] = 300
		}
		fmt.Println(s11)
	}
}

/* Range
类似迭代器操作
				1st value      2nd value
-------------+--------------+-------------+--------------
string			index			s[index]
array/slice		index			s[index]
map				key				m[key]
channel			element
*/
