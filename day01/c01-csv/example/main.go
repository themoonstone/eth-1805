package main

import (
	"eth-1805/day01/c01-csv"
	"fmt"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	file_name := "test.csv"
	// 创建文件，返回文件对象
	f , err := os.Create(file_name)
	if nil !=err {
		panic(err)	// 在默认情况下，直接中断程序的运行，打印出错误的原因
	}
	// 写入文件
	if err := c01_csv.WriteCSVToFile(f); nil != err {
		panic(err)
	}

	// CSV文件名
	csv_file_name := "test.csv"
	err = ReadCsv(csv_file_name)
	if nil != err {
		log.Errorf("open the csv file failed! %v\n", err)
	}
}

// 读取CSV文件
func ReadCsv(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if nil != err {
		fmt.Printf("error : %v\n", err)
		panic(err)
		return err
	}

	// []bytes --> io.reader
	movies, err := c01_csv.ReadCSV(strings.NewReader(string(content)))
	if nil != err {
		return err
	}
	fmt.Printf("movies : %v\n", movies)
	return nil
}