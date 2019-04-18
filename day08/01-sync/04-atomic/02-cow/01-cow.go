package main

import (
	"errors"
	"fmt"
	"sync/atomic"
)

// 代表并发安全的整数数组接口
type ConcurrentArray interface {
	// Set 用于设置指定下标的元素值
	Set(index int, elem int) (err error)
	// Get 用于获取指定下标的元素值
	Get(index int) (elem int, err error)
	// Len 用于获取长度
	Len() int
}

// 并发安全整数数组结构
type intArray struct {
	length 	int		// 数组长度
	val 	atomic.Value
}

// 新建一个并发安全数组的实例
func NewConcurrentArray(length int) ConcurrentArray {
	array := intArray{}
	array.length = length
	array.val.Store(make([]int, array.length))
	return &array
}

// 向指定下标写入数据
func (ia *intArray) Set(index int, elem int) (err error)  {
	if err = ia.checkIndex(index); err != nil {
		return
	}
	// COW(copy-on-write)写时复制算法
	// 竞态条件
	newArray := make([]int, ia.length)
	copy(newArray, ia.val.Load().([]int))
	newArray[index] = elem
	ia.val.Store(newArray)
	return
}

func (ia *intArray) Get(index int) (elem int, err error) {
	if err = ia.checkIndex(index); err != nil {
		return
	}
	if err := ia.checkValue(); nil != err {
		return
	}

	elem = ia.val.Load().([]int)[index]
	return 0, nil
}

func (ia *intArray) Len() int {
	return ia.length
}

// 检查index合法性
func (ia *intArray) checkIndex(index int) error  {
	if index < 0 || index >= ia.length {
		return fmt.Errorf("Index out of range [0, %d)!d\n", ia.length)
	}
	return nil
}

// 检查原子值
func (ia *intArray) checkValue() error {
	// 原子加载，
	v := ia.val.Load()
	if nil == v {
		return errors.New("Invalid int array")
	}
	return nil
}