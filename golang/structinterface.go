package main

import (
	"fmt"
	"reflect"
	"time"
)

type Stu struct {
	Sid  int
	Name string
}

func main() {

	t1 := time.Now()
	for i := 0; i < 100000; i++ {
		v := &Stu{100, "golang"}
		m := &Stu{}
		vv := reflect.ValueOf(v).Elem()
		mv := reflect.ValueOf(m).Elem()
		for i := 0; i < vv.NumField(); i++ {
			switch vv.Field(i).Kind() {
			case reflect.Int:
				mv.Field(i).SetInt(vv.Field(i).Int())
			case reflect.String:
				mv.Field(i).SetString(vv.Field(i).String())
			}
		}
		//		fmt.Println(m)
	}
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))

	var iv interface{}
	for i := 0; i < 100000; i++ {
		v := &Stu{100, "golang"}
		iv = v
		m := &Stu{}
		*m = *(iv.(*Stu))
		//		fmt.Println(m)
	}
	t3 := time.Now()
	fmt.Println(t3.Sub(t2))
}
