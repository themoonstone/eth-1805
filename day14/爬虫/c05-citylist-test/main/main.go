package main

import (
	"eth-1805/day14/爬虫/c05-citylist-test/engine"
	"eth-1805/day14/爬虫/c05-citylist-test/zhenai/parser"
)

// 1. 调用http.get获取指定网址的数据
// 2. 获取目标网站的首页数据

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}