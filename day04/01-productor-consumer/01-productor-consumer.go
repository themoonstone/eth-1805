package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// 生产者实现
func Producer(f int, out chan int)  {
	for i := 0; ; i++ {
		out <- i * f
	}
}
// 消费者实现
func Consumer(in chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main_sleep() {
	ch := make(chan int, 64)
	go Producer(3, ch)
	go Consumer(ch)
	time.Sleep(3*time.Second)
}

// 让用户自行控制退出
func main() {
	ch := make(chan int, 64)
	go Producer(3, ch)
	go Consumer(ch)

	// ctrl+c退出
	c := make(chan os.Signal, 1)
	//signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	//fmt.Printf("quit (%v)\n", <-c)

	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
}
