package main

import (
	"fmt"
	"sync"
)

func main() {
	// 开启3个goroutine,向通道中发送数据
	// main goroutine等待所有的goroutine完成处理逻辑之后，再退出
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		//wg.Done()
		wg.Add(-1)
	}()
	go func() {
		//wg.Done()
		wg.Add(-1)
	}()
	go func() {
		//wg.Done()
		wg.Add(-1)
	}()
	wg.Wait()
	fmt.Printf("all the sub goroutine is ended. \n")


}
