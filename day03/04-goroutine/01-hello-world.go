package main

import (
	"fmt"
	"sync"
	"time"
)

// goroutine
func main() {

	// 抛出一个goroutine
	go func() {
		fmt.Println("hello world")
	}()

	time.Sleep(1 * time.Nanosecond)
	// 主goroutine退出，会直接杀掉其它的goroutine
}

// mutext
func main_mutex() {
	var mu sync.Mutex
	// 抛出一个goroutine

	go func() {
		fmt.Println("hello world")
		mu.Lock()
	}()
	mu.Unlock()
}

// channel使用
func main_channel() {
	done := make(chan string)
	go func() {
		fmt.Println("hello world")
		<-done
	}()
	// 向channel中发送一个数据
	done<-"11"
}

