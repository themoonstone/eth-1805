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
	"regexp"
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
		"http://www.zhenai.com/zhenghun")
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
	printCityList(all)
}

// 编码识别
func determineEncoding(
	r io.Reader) encoding.Encoding{
	bytes, err:=bufio.NewReader(r).Peek(1024)
	if err!=nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(
		bytes, "")
	return e
}

// 提取城市和URL
func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	// -1 表示所有内容
	//matches := re.FindAll(contents, -1)
	//for _, m := range matches {
	//	fmt.Printf("%s\n", m)
	//}
	//fmt.Printf("匹配数量:%v\n", len(matches))
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("City:%s, URL:%s\n", m[2],m[1])
	}
}
