package main

import (
	"eth-1805/day14/爬虫/c04-citylist-parse/engine"
	"eth-1805/day14/爬虫/c04-citylist-parse/zhenai/parser"
)

// 1. 调用http.get获取指定网址的数据
// 2. 获取目标网站的首页数据

func main() {
	//resp, err := http.Get("http://www.zhenai.com/zhenghun")
	//if err != nil {
	//	// handle error
	//	fmt.Errorf("访问首页失败: %v\n", err)
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if nil != err {
	//	panic(err)
	//}
	//parser.ParseCityList(body)
	// 调用engine.Run开启爬虫程序
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}