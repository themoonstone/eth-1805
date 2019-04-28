package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

func main() {
	// 禁用GC(垃圾回收),保证在main函数执行结束前恢复GC
	//debug.SetGCPercent(-1)
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var count int32
	newFunc := func() interface{} {
		return atomic.AddInt32(&count, 1)
	}
	// 新建临时对象池
	pool := sync.Pool{New: newFunc}	// 只赋值，没有去调用newFunc函数
	// 真正的new函数是在get中进行调用的
	v1 := pool.Get()
	fmt.Printf("Value 1 : %v\n", v1)

	// 临时对象池的存取
	// todo put存值规则？？？
	//
	pool.Put(12)
	pool.Put(10)
	pool.Put(11)

	v2 := pool.Get()
	fmt.Printf("Value 2 : %v\n", v2)
	debug.SetGCPercent(100)
	// 手动GC
	// GC只回收put放入池中的内容，不回收new的内容
	runtime.GC()
	v3 := pool.Get()
	fmt.Printf("Value 3: %v\n", v3)
	pool.New = nil
	v4 := pool.Get()
	fmt.Printf("value 4 : %v\n", v4)
}
