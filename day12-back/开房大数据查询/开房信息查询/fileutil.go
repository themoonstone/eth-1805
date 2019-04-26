package main

import (
	"os"
	"fmt"
)

func CheckIfFileExist(filename string) (exists bool, info string) {
	fileInfo, err := os.Stat(filename)
	fmt.Println(fileInfo, err)
	if fileInfo != nil && err == nil {
		//fmt.Printf("%s文件存在\n", filename)
		exists = true
	} else if os.IsNotExist(err) {
		//fmt.Printf("%s文件不存在\n", filename)
		exists = false
	} else {
		fmt.Println("搞不清存不存在!")
		exists = false
		info = "发生了一些奇奇怪怪的事情..."
	}
	return
}
