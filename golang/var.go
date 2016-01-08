//声明相关
package main

import (
	"fmt"
	"unsafe"
)

//常量值必须是编译期可确定的数字、字符串、布尔值
//还可以是len、cap、unsafe.Sizeof 等编译期可确定结果的函数返回值
const (
	Pi     = 3.1415926
	Sunday = iota //自增类型
	Monday
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(b)
)

//var 定义变量
var x int //自动初始化为零值
var f float32 = 1.6
var s = "abc" //可省略变量类型，由编译器推断

func test() (int, string) {
	return 1, "abc"
}
func main() {
	y := 123           //函数内部可省略var关键字
	var y0, y1, y2 int //一次定义多个变量
	var y3, y4 = "abc", 123
	fmt.Println(x, f, s, y, y0, y1, y2, y3, y4) //编译器会将未使用的局部变量当做错误
	data, i := [3]int{0, 1, 2}, 0               //多变量赋值先计算所有相关值，再从左至右依次赋值
	fmt.Println(data, i)
	_, str := test() // _忽略值占位
	fmt.Println(str)
	w := "abc"
	// q := 1
	fmt.Println(&w)
	w, q := "bcd", 20
	fmt.Println(&w, q)
	fmt.Println(Sunday, Monday)
	var n0 byte = 100
	var n1 int = int(b)
	fmt.Println(n0, n1)
}
