package main

import (
	"flag"
	"github.com/golang/glog"
	"reflect"
)

type Stu struct {
	Id   int32
	Name string
}

func main() {
	//	var i interface{}
	flag.Parse()
	s := "this is a string"
	sType := reflect.TypeOf(s)
	glog.Infoln(int(sType.Kind()), sType.Name(), sType.PkgPath(), sType.Size())

	integer := int32(100)
	integerType := reflect.TypeOf(integer)
	glog.Infoln(int(integerType.Kind()), integerType.Name(), integerType.PkgPath(), integerType.Size())

	float := float64(3.14)
	floatType := reflect.TypeOf(float)
	glog.Infoln(int(floatType.Kind()), floatType.Name(), floatType.PkgPath(), floatType.Size())

	stu := Stu{100, "liu"}
	stuType := reflect.TypeOf(stu)
	glog.Infoln(int(stuType.Kind()), stuType.Name(), stuType.PkgPath(), stuType.Size())
	glog.Infoln(stuType.Field(0), stuType.Field(1))
	fieldId, ok1 := stuType.FieldByName("Id")
	fieldIdName, ok2 := stuType.FieldByName("Name")
	glog.Infoln(fieldId, fieldIdName, ok1, ok2)

	m := make(map[int32]string)
	m[100] = "liu"
	mType := reflect.TypeOf(m)
	glog.Infoln(int(mType.Kind()), mType.Name(), mType.PkgPath(), mType.Size())
	glog.Infoln(mType.Key().Name(), mType.Elem().Name())

	arr := [5]int{1, 2, 3, 4, 5}
	arrType := reflect.TypeOf(arr)
	glog.Infoln(int(arrType.Kind()), arrType.Name(), arrType.PkgPath(), arrType.Size())
	glog.Infoln(arrType.Elem().Kind(), arrType.Elem().Name(), arrType.Elem().Size())
	glog.Infoln(arrType.Len())
}
