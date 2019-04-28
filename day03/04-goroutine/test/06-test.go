package main

import "fmt"

func primeNum(intChan chan int, primeChan chan int)  {
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
	fmt.Println("退出")
}

func storeToIntChan(intChan chan int)  {
	for i := 1; i <= 50; i++ {
		intChan <- i
	}
	// 考虑关闭通道
	close(intChan)
	fmt.Println("storeToIntChan")
}

func main() {
	intChan := make(chan int)
	primeChan := make(chan int)

	go storeToIntChan(intChan)

	for i := 0; i < 4; i ++ {
		go primeNum(intChan, primeChan)
	}
	for {
		_, ok := <- primeChan
		if !ok {
			break
		}
		//fmt.Printf("prime number is %d\n", res)
	}
}