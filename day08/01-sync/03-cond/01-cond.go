package main

import (
	"fmt"
	"sync"
	"time"
)

// 新建条件变量实例
var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func test(x int)  {
	cond.L.Lock()
	fmt.Println("a:", x)

	cond.Wait()
	fmt.Println("b:", x)
	time.Sleep(2 * time.Second)
	cond.L.Unlock()
}

func main() {
	for i := 0; i < 5; i++ {
		go test(i)
	}
	fmt.Println("start all goroutine...")
	time.Sleep(1 * time.Second)
	fmt.Println("signal 1")
	cond.Signal()
	time.Sleep(1 * time.Second)
	fmt.Println("signal 2")
	cond.Signal()
	time.Sleep(1 * time.Second)
	fmt.Println("BroadCast")
	cond.Broadcast()
	time.Sleep(10*time.Second)
}
