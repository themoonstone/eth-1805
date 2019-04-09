package main

import (
	"fmt"
	"reflect"
)

func test(b interface{}) {
	val := reflect.ValueOf(b)
	fmt.Printf("val : %v\n", val)
	// 断言
	fmt.Printf("type : %v\n", reflect.TypeOf(b))
	//iv := val.Interface()
	//
	i := interface{}(val)
	fmt.Printf("i type : %v\n", reflect.TypeOf(i))

	//per, ok := i.(*map[string]string)
	//if ok {
	//	m := *per
	//	m["f"] = "S"
	//}else {
	//	fmt.Printf("assert failed!\n")
	//}
	//if ok {
	//	m := *per
	//	m["f"] = "E"
	//} else {
	//	fmt.Printf("assert failed!\n")
	//}
	//fmt.Printf("per : %v\n", per)
}

func main() {
	//var num int = 20
	//test(&num)
	//fmt.Printf("num = %v\n", num )
	mp := make(map[string]string)
	mp["f"] = "F"
	test(&mp)
	fmt.Printf("map : %v\n", mp)
}