package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 大小判断
func SizeOf()  {
	var p float64 = 99
	// int string bool
	fmt.Println(reflect.TypeOf(unsafe.Sizeof(p)))
	fmt.Println(unsafe.Sizeof(p))
}

func main() {
	//SizeOf()
	type W struct {
		a byte	// 1
		b byte	// 4
		c int64 // 8
	}

	w := W{0,0,0}

	// 将结构体指针转换为通用指针
	s := unsafe.Pointer(&w)
	// int
	// uintptr
	up0 := uintptr(s)
	pb := (*byte)(s)
	*pb = 10
	fmt.Println(w)
	// uintptr 是平台相关
	//var w *W
	//fmt.Println(unsafe.Sizeof(w))
	//
	//fmt.Println(unsafe.Sizeof(w.a))
	//fmt.Println(unsafe.Sizeof(w.b))
	//fmt.Println(unsafe.Sizeof(w.c))

	// func Alignof(x ArbitraryType) uintptr 对齐
	//fmt.Println(unsafe.Sizeof(*w))
	//fmt.Println(unsafe.Alignof(w.a))
	//fmt.Println(unsafe.Alignof(w.b))
	//fmt.Println(unsafe.Alignof(w.c))

	// 偏移到第2个字段
	up := up0 + unsafe.Offsetof(w.b)
	// int int64

	p1 := unsafe.Pointer(up)
	p2 := (*byte)(p1)
	*p2 = 20
	fmt.Println(w)
	// 可以改变结构体的私有变量

}
