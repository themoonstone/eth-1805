package main

// goroutine和channel 案例
// 1-10000000的数字中，哪些是素数
// 素数判断 ，使用并发的方式进行任务的处理
// 基本思路：将统计素数的任务分给多个goroutine去完成
// 用channel接收数据，标志退出...
// 代码实现

// 素数判断功能函数
// 从intChan中取出数据，判断是否是一个素数
// 将判断出的素数存入primeChan中
/*
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool)  {
	for {
		var flag bool = false
		// ok-dom
		num, ok := <- intChan
		if !ok {
			break
		}
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
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	// 考虑关闭通道
	//close(intChan)
}
func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
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
		res, ok := <- primeChan
		if !ok {
			break
		}
		fmt.Printf("prime number is %d\n", res)
	}
	fmt.Println("main 线程退出")
}*/