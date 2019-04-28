package main

import "fmt"

func main() {
	// 开启3个goroutine,向通道中发送数据
	// main goroutine等待所有的goroutine完成处理逻辑之后，再退出

	sign := make(chan byte, 3)
	go func() {
		sign <- 2
	}()
	go func() {
		sign <- 3
	}()
	go func() {
		sign <- 4
	}()

	for i := 0; i < 3; i++ {
		fmt.Printf("g%d is ended. \n", <-sign)
	}
}
