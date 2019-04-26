package main

import (
	"os"
	"bufio"
	"io"
	"strings"
	"fmt"
)

func main() {

	//打开原始数据
	file, _ := os.Open("d:/temp/kaifangX.txt")
	defer file.Close()

	//准备一个优质数据文件
	goodFile, _ := os.OpenFile("d:/temp2/kaifang_good.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer goodFile.Close()

	//准备一个优质数据文件
	badFile, _ := os.OpenFile("d:/temp2/kaifang_bad.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer badFile.Close()

	reader := bufio.NewReader(file)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		gbkStr := string(lineBytes)
		lineStr, _ := ConvertEncoding(gbkStr, "GBK", "UTF-8")
		//fmt.Println(lineStr)

		fields := strings.Split(lineStr, ",")
		if len(fields) > 1 && len(fields[1])==18{
			//摘取到另一个优质的数据文件中
			goodFile.WriteString(lineStr+"\n")
			fmt.Println("Good:",lineStr)
		}else{
			badFile.WriteString(lineStr+"\n")
			fmt.Println("Bad:",lineStr)
		}

	}

}
