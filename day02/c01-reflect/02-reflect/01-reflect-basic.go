package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{})  {
	// 1. 获取类型
	rTyp := reflect.TypeOf(b)
	fmt.Println("rtype = ", rTyp)

	// 2. 获取value
	val := reflect.ValueOf(b)
	fmt.Printf("val : %v\n", val)
	n2 := 10 + val.Int()
	fmt.Printf("n2 : %v\n", n2)

	// 断言
	iv := val.Interface()
	num2 := iv.(int)
	fmt.Printf("num2 : %v\n", num2)

}

// 结构与反射
func reflectTest02(b interface{})  {
	// 1. 获取类型
	rTyp := reflect.TypeOf(b)
	fmt.Println("rtype = ", rTyp)

	// 2. 获取value
	val := reflect.ValueOf(b)
	fmt.Printf("val : %v\n", val)

	// 3. 断言
	iv := val.Interface()
	fmt.Printf("iv=%v\n", iv)
	// ok-dom
	per, ok := iv.(Person)
	if ok {
		fmt.Printf("person : %v\n", per.Name)
	}


}

// Person
type Person struct {
	Name string
	Age		int
}

func main()  {
	var num int = 100
	reflectTest01(num)

	// 定义一个结构体
	per := Person{
		Name: "troy",
		Age: 10,
	}

	reflectTest02(per)
}