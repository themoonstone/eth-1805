package main

import (
	"fmt"
	"math"
)

// 服务器错误码
// 异常的捕获

// 求面积
func GetCircle(r float64) float64 {
	if r < 0 {
		panic("半径非负")
	}
	return math.Pi * r * r
}

// 调用
func main() {

	// 捕获异常
	defer func() {
		// 捕获panice产生的异常
		err := recover()
		if nil != err {
			fmt.Printf("有异常出现 : %v\n", err)
		}else {
			fmt.Printf("normal\n")
		}
	}()

	fmt.Println("area of circle is : ", GetCircle(-12.5))
}