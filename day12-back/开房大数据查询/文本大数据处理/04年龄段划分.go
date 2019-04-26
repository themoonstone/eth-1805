package main

import (
	"os"
	"strconv"
	"sync"
	"bufio"
	"strings"
	"io"
	"fmt"
)

type Ager struct {
	//年代
	decade string //190x...197x,198x,199x,200x,201x,
	//
	file     *os.File
	chanData chan string
}

func main() {
	//创建一大堆年代对象
	agersMap := make(map[string]*Ager)
	for i := 190; i < 202; i++ {
		ager := Ager{decade: strconv.Itoa(i)}
		file, _ := os.OpenFile("d:/temp2/"+ager.decade+"x.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		ager.file = file
		defer ager.file.Close()
		ager.chanData = make(chan string,0)
		agersMap[ager.decade] = &ager
	}

	//为每一个年代开辟一个写入协程
	for _,ager := range agersMap{
		dWg.Add(1)
		go write2File(ager)
	}

	//读入未分类数据
	file, _ := os.Open("d:/temp2/kaifang_good.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		//断行-判断年代-丢入响应的管道
		lineStr, err := reader.ReadString('\n')
		if err==io.EOF{
			for _,ager := range agersMap{
				close(ager.chanData)
			}
			break
		}
		decade := strings.Split(lineStr, ",")[1][6:9]
		if ager:=agersMap[decade];ager!=nil{
			agersMap[decade].chanData <- lineStr + "\n"
		}else{
			fmt.Println("\n\n\n\n\n\n\n\n",lineStr,"\n\n\n\n\n\n\n\n")
		}

	}

	//阻塞等待结束
	dWg.Wait()
}

var dWg sync.WaitGroup
func write2File(ager *Ager) {
	for contentStr := range ager.chanData{
		ager.file.WriteString(contentStr)
		fmt.Print(ager.decade,"+=",contentStr)
	}
	dWg.Done()
}
