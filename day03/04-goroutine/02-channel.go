package main

import (
	"fmt"
	"time"
)

// 计算指定范围的阶乘
// 放到一个map中，进行显示

var mMap = make(map[int]int, 10)

func alg(n int)  {
	res := 1
	for i := 1; i <= n; i++ {
		res = res * i
	}
	mMap[n] = res
}

func main() {
	for i:=1; i <= 9; i++ {
		go alg(i)
	}
	time.Sleep(10 * time.Second)
	// 显示结果
	for i, v := range mMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}