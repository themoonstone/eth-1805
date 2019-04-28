package main

import (
	"eth-1805/day14/爬虫/c03-fetcher/regexp"
	"fmt"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
)

// 1. 调用http.get获取指定网址的数据
// 2. 获取目标网站的首页数据

func main_write_file() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		// handle error
		fmt.Errorf("访问首页失败: %v\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile("cityList.html", body, 0666)
	if nil != err {
		log.Errorf("write to file failed! %v\n", err)
	}
	fmt.Printf("body : %v\n", string(body))
}

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		// handle error
		fmt.Errorf("访问首页失败: %v\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		panic(err)
	}
	regexp.PrintCityList(body)
}