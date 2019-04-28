package main

import (
	"eth-1805/day01/c04-testing/01-testing"
	"fmt"
	"os"
)

func main() {
	var s int = 10
	if s == 10 {
		fmt.Println("退出")
		os.Exit(1)
	}
	res := _1_testing.Add(10,20)
	if res != 12 {
		//fmt.Printf("add error : %v\n", errors.New("the result is not equal expected value"))
		//log.Printf("add error : %v\n", errors.New("the result is not equal expected value"))
	}else {
		fmt.Println("correct...")
	}


}
