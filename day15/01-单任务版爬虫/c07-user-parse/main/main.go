package main

import  (
	"1805/day14/01-单任务版爬虫/c07-user-parse/engine"
	"1805/day14/01-单任务版爬虫/c07-user-parse/zhenai/parser"
)

// 添加编码识别工作
func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})

	//chiReg := regexp.MustCompile(`[^\u4e00-\u9fa5]`)
	//chiReg.FindAllSubmatch([]byte("haha 哈哈东方科技dfjkdj2323"), -1)
	//fmt.Println(chiReg.ReplaceAllString("haha 哈哈东方科技dfjkdj2323", ""))
}
