package main

import (
	"fmt"
	"os"
	"time"
)

func main_io_channel() {
	a, b := make(chan int, 3), make(chan int)

	go func() {
		v, ok := 0, false
		s := ""
		for {
			select {
			// ok-dom
			case v, ok = <-a:
				s = "a"
			case v, ok = <-b:
				s = "b"
			}
			if ok {
				fmt.Println(s, v)
			} else {
				os.Exit(0)
			}
		}
	}()

	for i := 0; i < 4; i++ {
		select {
		// 随机选择可用的channel，发送数据
		case a <- i:
		case b <- i:
		}
	}
	close(a)
	//close(b)
	// 阻塞main goroutine
	select {}
}

// 判断超时
func main_timeout() {
	// 如果在某些情况下，任务处理的goroutine陷入长期阻塞，为了服务程序的正常运行
	// 需要引入类似于心跳机制的超时处理
	w := make(chan bool)
	c := make(chan int, 2 )
	go func() {
		select {
		case v := <- c:
			fmt.Println(v)
		case <- time.After(3 *time.Second):
			fmt.Println("timeout")
		}
		w <- true
	}()
	//c <- 1
	<- w
}

func Add(a, b int)int  {
	return a + b
}

func Sub(a, b int) int {
	return a * b
}
// 退出quit
func main()  {
	// 如果在某些情况下，任务处理的goroutine陷入长期阻塞，为了服务程序的正常运行
	// 需要引入类似于心跳机制的超时处理
	w := make(chan bool)
	c := make(chan int, 2 )
	quit := make(chan bool)	// 退出标识
	go func() {
		select {
		case v := <- c:
			fmt.Println(v)
		case v:= <-quit:
			fmt.Println(v)
			Sub(2,3)
			os.Exit(0)
		case <- time.After(3 *time.Second):
			fmt.Println("timeout")
		}
		w <- true
	}()
	go func() {
		Add(1, 2)
		quit <- true
	}()
	//c <- 1
	<- w

	// TODO 实现一个select企业需求处理
}

