package main

import (
	"1805/备课/day14/实现/01-单任务版爬虫/c04-citylist-parse/engine"
	"1805/备课/day14/实现/01-单任务版爬虫/c04-citylist-parse/zhenai/parser"
)

// 添加编码识别工作
func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
