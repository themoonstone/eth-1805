package main

import "fmt"

// 解析结果
type ParseResult struct {
	//Requests 	[]Request
	Items		[]interface{}
}

// 添加一个空的解析器
func NilParser([]byte) ParseResult  {
	return ParseResult{}
}

func main() {
	fmt.Printf("NilParser : %v--%p\n", NilParser(nil), NilParser(nil))
}
