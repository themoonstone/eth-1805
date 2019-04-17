package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	fmt.Println("Lock the lock. (main)")
	mutex.Lock()
	fmt.Println("The lock is locked. (main)")
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("Lock the lock. (g%d)\n", i)
			mutex.Lock()
			fmt.Printf("the lock is locked. (g%d)\n", i)
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Unlock the lock. (main)")
	mutex.Unlock()
	fmt.Printf("The lock is unlocked (main)\n")
	time.Sleep(1 * time.Second)
}
