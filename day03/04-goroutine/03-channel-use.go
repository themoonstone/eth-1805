package main

import (
	"fmt"
	"time"
)

// channel基本操作

//func main()  {
//	intChan := make(chan int, 100)
//	go func() {
//		for i := 0;i < 200; i++ {
//			intChan <- i
//		}
//	}()
//	for v := range intChan {
//		fmt.Printf("v : %v\n", v)
//	}
//}

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}

func main_basic_opera() {
	// 1. 创建
	// 无缓冲channel (buffer)
	//var intChan chan int = make(chan int)
	//fmt.Printf("intChannel is %v, %p\n", intChan, &intChan)

	//go func() {
	//	<-intChan
	//}()
	//intChan <- 100

	// 缓冲channel
	var intBufferChannel chan int = make(chan int, 3)
	intBufferChannel <- 100
	num := 211
	intBufferChannel <- num
	// 查看一下长度和容量
	fmt.Printf("length of channel is %d, cap of channel is %d\n",
		len(intBufferChannel), cap(intBufferChannel))
	// 从channel中读取数据
	go func() {
		fmt.Printf("the first data of channel is %d\n", <-intBufferChannel)
		fmt.Printf("the second data of channel is %d\n", <-intBufferChannel)
		fmt.Printf("the third data of channel is %d\n", <-intBufferChannel)
	}()

	// ok-dom

}

// channel遍历和关闭
func main_range() {
	//intChan := make(chan int, 10)
	//
	//for i := 0; i < 10; i++ {
	//	intChan <- i * 2
	//}
	//// 先关闭channel
	//close(intChan)
	// 面试高频题
	// 从一个已经关闭的channel中读取数据，能不能正常读取
	//for v := range intChan {
	//	fmt.Printf("v = %d\n", v)
	//}
	// panic: send on closed channel
	//intChan <- 200

	nilChan := make(chan bool)
	close(nilChan)
	fmt.Printf("data is : %v\n", <- nilChan)
	// 从一个关闭的空的channel中读取数据，会直接读取到该类型的初始值
	// 不能重复关闭channel
	//close(nilChan)

	// nil
	var nilChannel chan int
	// panic: close of nil channel
	close(nilChannel)
}

// 单向channel
func main_single_channel() {

	// 1. 直接声明定义
	//sendChan := make(chan <-int)
	//go func() {
	//	sendChan <- 10
	//}()
	//<-sendChan

	c := make(chan int)
	// 向channel中存储数据
	var send chan <- int = c
	// 从channel中存储数据
	var recv <- chan int = c

	go func() {
		for x := range recv {
			fmt.Println("x = ", x)
		}
	}()

	go func() {
		defer close(c)
		for i := 0; i < 3; i++ {
			send <- i
		}
	}()

	time.Sleep(2 * time.Second)
}

// 单向channel非法操作
func main_single_invalid_operation() {
	c := make(chan int)
	// 向channel中存储数据
	// chan <- int代表一个独立的类型
	var send chan <- int = c
	// 从channel中存储数据
	// <- chan int代表一个独立的类型
	//var recv <- chan int = c

	// 不能够对单向channel进行逆向操作
	//<-send
	//recv<-1

	// 关闭通道
	close(send)
	// close不能用于接收端
	//close(recv)

	// 将单向通道转成双向通道
	//var b chan int
	//b = (chan int)(recv)
}