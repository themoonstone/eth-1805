package main

import (
	"fmt"
	"reflect"
)

func SetElem(a interface{})  {
	// 获取类型
	tp := reflect.ValueOf(a)
	fmt.Printf("v.kind() is %s\n", tp.Kind())
	tp.Elem().SetString("the new moon stone")
}

func main() {
	//var num int = 20
	//SetElem(&num)
	//fmt.Printf("num = %v\n", num )

	var s string = "themoonstone"
	SetElem(&s)
	fmt.Printf("s : %v\n", s)
}