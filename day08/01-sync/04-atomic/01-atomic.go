package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	// 原子变量
	var countVal atomic.Value
	countVal.Store([]int{1,3,5,7})
	fmt.Printf("the old value : %v\n", countVal)
	anotherStore(&countVal)
	fmt.Printf("the anotherStore value : %v\n", countVal)
}

func anotherStore(countVal *atomic.Value)  {
	countVal.Store([]int{2,4,6,8})
}