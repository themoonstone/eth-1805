
package main

import "fmt"

func demo(i int)(){

	var arr[10]int
	defer func(){
		// 第一种捕获方式，执行这个时将第二种捕获方式注释掉
		if recover()!=nil{
			fmt.Println(recover())
		}

		// 第二种捕获方式，执行这个时将第一种捕获方式注释掉
		err := recover()
		if err!=nil{
			fmt.Println(err)
		}

	}()
	arr[i]=100
	fmt.Println(arr)
}
// TODO 两种捕获方式输出的值为什么不一样？
func main() {
	demo(11)
	fmt.Printf("a")
}