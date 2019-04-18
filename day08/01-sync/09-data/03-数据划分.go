package main

import (
	"bufio"
	"github.com/axgle/mahonia"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)
const (
	SRC_ENCODE = "GBK"
	DST_ENCODE = "UTF-8"
)

var wg sync.WaitGroup

type Age struct {
	rules string		// 190x, 191x... 200x, 201x
	file *os.File
	chanData chan string
}

// 筛选出包含身份证号信息，存入文件
func main() {
	// 年代对象
	agerMap := make(map[string]*Age)
	for i := 190; i < 202; i++ {
		age := Age{rules:strconv.Itoa(i)}
		file, _ := os.OpenFile("age/"+ age.rules + "x.txt",
			os.O_WRONLY|os.O_APPEND|os.O_CREATE,0644)
		age.file = file
		defer age.file.Close()
		age.chanData = make(chan string, 0)
		agerMap[age.rules] = &age

	}
	// 开辟goroutine
	for _, age := range agerMap {
		wg.Add(1)
		go writeFile(age)
	}

	file, _ := os.Open("./useful.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		lineString,  err := reader.ReadString('\n')
		if err == io.EOF {
			for _, ager := range agerMap {
				close(ager.chanData)
			}
			break
		}

		result := strings.Split(lineString, ",")[1][6:9]
		if age := agerMap[result]; age != nil {
			agerMap[result].chanData <- lineString + "\n"
		}
	}
	wg.Wait()
}

// 写入指定文件
func writeFile(age *Age)  {
	for content := range age.chanData {
		age.file.WriteString(content)
	}
	wg.Done()
}

// 公共函数
// 编码解码函数
// srcStr:待编码的字符中
// srcEnccoding:原编码格式
// dstEncoding:目标编码格式
func ConvertEncoding(srcStr, srcEncoding, dstEncoding string)(dstString string, err error)  {
	srcDecoder := mahonia.NewDecoder(srcEncoding)
	dstDecoder := mahonia.NewDecoder(dstEncoding)
	// 转UTF-8函数调用
	utfStr := srcDecoder.ConvertString(srcStr)
	_, dstBytes, err := dstDecoder.Translate([]byte(utfStr), true)
	if nil != err {
		return
	}
	dstString = string(dstBytes)
	return
}