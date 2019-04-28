package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main1() {
	resp, err := http.Get(
		"http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
}

// 添加编码识别工作
func main() {
	resp, err := http.Get(
		"http://album.zhenai.com/u/1456179143")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
}

// 编码识别
func determineEncoding(
	r io.Reader) encoding.Encoding{
		//TODO 需要搞懂这里这个PEEK
		// 这里这个peek是什么意思？
		// 对r读1024个字节，存给bytes
		// 为什么非要加一个peek(1024)，因为只是进行编码判断，所以1024个字节应该是够了
	bytes, err:=bufio.NewReader(r).Peek(1024)
	if err!=nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(
		bytes, "")
	return e
}
