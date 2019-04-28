package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"github.com/axgle/mahonia"
	"bufio"
	"io"
	"os"
)

//一次性将全部数据载入内存（不可取）
func main1() {
	//一次性将全部数据载入内存（不可取）
	contentBytes, err := ioutil.ReadFile("d:/temp/kaifangX.txt")
	if err != nil {
		fmt.Println("读入失败", err)
	}
	contentStr := string(contentBytes)

	//逐行打印
	lineStrs := strings.Split(contentStr, "\n\r")
	for _, lineStr := range lineStrs {
		//fmt.Println(lineStr)
		newStr, _ := ConvertEncoding(lineStr, "GBK", "UTF-8")
		fmt.Println(newStr)
	}

}

//基于磁盘和缓存的读取
func main2() {

	file, _ := os.Open("d:/temp/kaifangX.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		gbkStr := string(lineBytes)
		utfStr, _ := ConvertEncoding(gbkStr, "GBK", "UTF-8")
		fmt.Println(utfStr)
	}

}

//srcStr = 待转码的原始字符串
//srcEncoding = 原始字符串的编码（字符集）
//dstEncoding = 目标编码（字符集）
func ConvertEncoding(srcStr string, srcEncoding string, dstEncoding string) (dstStr string, err error) {
	srcDecoder := mahonia.NewDecoder(srcEncoding)
	dstDecoder := mahonia.NewDecoder(dstEncoding)

	utfStr := srcDecoder.ConvertString(srcStr)
	_, dstBytes, err := dstDecoder.Translate([]byte(utfStr), true)
	if err != nil {
		return
	}
	dstStr = string(dstBytes)
	return
}
