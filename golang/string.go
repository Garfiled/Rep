//包名
package main

import (
	"fmt"
	"unsafe"
)

// 字符串是不可变值类型，内部用指针指向UTF-8字节数组
// runtime.h
// struct String
// {
// byte* str;
// intgo len;
// };
// 索引号访问字符byte
func main() {
	str := "Hello, world!"
	str1 := str[:5] //支持用两个索引号返回字串，字串依然指向原字节数组，仅修改了指针和长度属性
	str2 := str[7:]
	str3 := str[1:5]
	fmt.Println(str1, len(str1), str2, str3)
	s := "abc汉字"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c,", s[i])
	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%c,", r)
	}
	//修改字符串
	s0 := "abcd"
	bs := []byte(s0) //重新分配内存，并复制字节数组
	bs[1] = 'B'
	fmt.Println("\n-----")
	fmt.Println(bs, string(bs), s0)
	fmt.Println(&s0, *(*(*byte))(unsafe.Pointer(&bs)))
	u := "电脑"
	us := []rune(u) //copy走了
	us[1] = '话'
	fmt.Println(string(us), u)
}
