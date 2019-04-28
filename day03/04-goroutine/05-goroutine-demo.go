package main

import "fmt"

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool)  {
	fmt.Println("primeNum")
	for {
		var flag bool = false
		// ok-dom
		num, ok := <- intChan
		if !ok {
			break
		}
		fmt.Printf("num : %v\n", num)
		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = true
				break
			}
		}
		if !flag {
			// 是一个素数
			primeChan <- num
		}

	}
	exitChan <- true
	fmt.Println("退出")
}

// 向intChan中存入指定数量的数值
func storeToIntChan(intChan chan int)  {
	for i := 1; i <= 50; i++ {
		intChan <- i
	}
	// 考虑关闭通道
	//close(intChan)
	fmt.Println("storeToIntChan")
}

func main() {
	intChan := make(chan int, 10)
	primeChan := make(chan int, 20)
	// 添加一个标志退出的channel
	exitChan := make(chan bool, 4)
	// 另起一个goroutine进行数据存储
	go storeToIntChan(intChan)

	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<- exitChan
		}
		close(primeChan)
	}()
	for {
		_, ok := <- primeChan
		if !ok {
			break
		}
		//fmt.Printf("prime number is %d\n", res)
	}
	fmt.Println("main 线程退出")
}