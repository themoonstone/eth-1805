// go error包实例说明
package main

import (
	"fmt"
	"time"
)

// MyError is an error implementation that includes a time and message.
type MyError struct {
	When time.Time
	What string
}

func (err MyError) String() string {
	return "string 接口的实现"
}
// 自己定制的error显示函数
func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	/*
		MyError{
			time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
			"the file system has gone away",
		}
	*/
	err := MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
	fmt.Println(err.String())
	return err
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
