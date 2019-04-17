package main

import (
	"fmt"
	"sync"
)

func main() {
	defer func() {
		fmt.Println("Try to recover the panic")
		if p := recover(); p!= nil {
			fmt.Printf("Recovered the panic(%#v). \n", p)
		}
	}()

	var mutex sync.Mutex
	fmt.Println("Lock the lock. (main)")
	mutex.Lock()
	fmt.Println("The lock is locked. (main)")
	mutex.Unlock()
	fmt.Println("the lock is unlocked")
	// 在go1.8以后，类似于重复解锁这一类运行时的panic不可再通过
	// recover进行捕获
	mutex.Unlock()
	fmt.Println("continue...")
}