package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
)

// 下标越界
func main_array()  {
	//panic("抛出一个异常")
	var s [2]int
	for i := 0; i < 3; i++ {
		s[i] = i
	}
}
// 向未分配内存的map中传值
func main_map_slice() {
	// 没有分配内存
	var Map map[string]string = make(map[string]string)
	Map["first"] = "panic"
	fmt.Println(Map)

	// 向未分配内存的切片中插入数据
	var slice []string
	slice[0] = "first"
	fmt.Println(slice)
}

// interface {} is string, not int
// 类型断言采用的类型不匹配
func main_assert() {
	mSlice := make([]interface{}, 0)
	mSlice = append(mSlice, "the interface")
	fmt.Println(mSlice[0].(int))
}

// 向空指针赋值
func main()  {
	// 客户信息
	type NameStruct struct {
		Name string
		Age 	int
	}
	// 银行
	type Bank struct {
		nameS *NameStruct
	}

	b := Bank{}
	// 通过new对结构体进行内存分配
	//b.nameS = new(NameStruct)
	fmt.Printf("before %p\n", b.nameS)
	b.nameS = &NameStruct{}
	fmt.Printf("after %p\n", b.nameS)
	b.nameS.Age = 100
}

// 自己设计panic抛出逻辑
func mainNumber()  {
	a := 100
	if a > 99 {
		log.Panicf("the number is not valid! expected is 99, get is %d\n", a)
	}
}