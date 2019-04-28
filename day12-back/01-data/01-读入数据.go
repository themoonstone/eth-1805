package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//const (
//	SRC_ENCODE = "GBK"
//	DST_ENCODE = "UTF-8"
//)

// 一次性全部
func mainReadAllData() {
	content, err := ioutil.ReadFile("./kfc.txt")
	if nil != err {
		fmt.Printf("读入失败", err)
	}
	contentStr := string(content)

	// 单行打印
	lineStr := strings.Split(contentStr, "\n\r")
	for _, line := range lineStr {
		newStr, _ := ConvertEncoding(line, "GBK", "UTF-8")
		fmt.Println(newStr)
	}
}


// 按行读取
func main() {
	file, _ := os.Open("./kf.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		srcStr := string(lineBytes)
		utfStr, _ := ConvertEncoding(srcStr, SRC_ENCODE, DST_ENCODE)
		fmt.Println(utfStr)
	}
}
// 公共函数
// 编码解码函数
// srcStr:待编码的字符中
// srcEnccoding:原编码格式
// dstEncoding:目标编码格式
//func ConvertEncoding(srcStr, srcEncoding, dstEncoding string)(dstString string, err error)  {
//	srcDecoder := mahonia.NewDecoder(srcEncoding)
//	dstDecoder := mahonia.NewDecoder(dstEncoding)
//	// 转UTF-8函数调用
//	utfStr := srcDecoder.ConvertString(srcStr)
//	_, dstBytes, err := dstDecoder.Translate([]byte(utfStr), true)
//	if nil != err {
//		return
//	}
//	dstString = string(dstBytes)
//	return
//}