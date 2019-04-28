package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// 同步--只执行一次
func main() {
	var count int
	var once sync.Once
	max := rand.Intn(100)
	for i := 0; i < max; i++ {
		once.Do(func() {
			count++
		})
	}
	fmt.Printf("Count : %d\n", count)
}