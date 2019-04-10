package main

import (
	"fmt"
	"reflect"
	"strings"
)

// 实现一个通用的加法函数
//  不光是能够进行单一类型的加法运算，可以实现任意其它类型的加法运算
func add(args []reflect.Value) (result []reflect.Value) {
	var ret reflect.Value
	// 判断类型
	switch args[0].Kind() {
	case reflect.Int:
		n :=0
		for _, a := range args {
			n += int(a.Int())
		}
		ret = reflect.ValueOf(n)
	case reflect.String:
		// 分配切片内存
		ss := make([]string, 0, len(args))
		for _, s := range args {
			ss = append(ss, s.String())
		}
		ret = reflect.ValueOf(strings.Join(ss, ""))
	}
	// 最终结果
	result = append(result, ret)
	return
}

// 将函数指针与通用算法函数关联
func makeAdd(fptr interface{})  {
	fn := reflect.ValueOf(fptr).Elem()
	// reflect函数创建
	v := reflect.MakeFunc(fn.Type(), add)

	// 指向通用算法
	fn.Set(v)

}

func main() {
	// 整形相加
	// 字符串相加
	//var s1 interface{}
	//s1 = 10
	//
	//add([]reflect.Value{reflect.ValueOf(s1)})
	//var w string = "sss"
	//s := unsafe.Pointer(&w)
	////s := &w
	//fmt.Printf("s : %v\n", s)
	//// todo 回顾， 为什么要用两个指针符号
	// ss := int64(*(*int)(s))
	// fmt.Println("ss : ", ss)

	//var a int = 100
	//ar := reflect.ValueOf(a)
	//var b int = 100
	//br := reflect.ValueOf(b)
	// fmt.Println(add([]reflect.Value{ar, br}))

	var intAdd func(x, y int) int
	var stringAdd func(x, y string) string
	makeAdd(&intAdd)
	makeAdd(&stringAdd)
	fmt.Println(intAdd(10,20))
	fmt.Println(stringAdd("hello", "world"))
}