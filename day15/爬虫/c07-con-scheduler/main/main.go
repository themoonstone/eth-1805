package main

import (
	"1805/day14/爬虫/c07-con-scheduler/engine"
	"1805/day14/爬虫/c07-con-scheduler/scheduler"
	"1805/day14/爬虫/c07-con-scheduler/zhenai/parser"
)

// 添加编码识别工作
func main() {
	// 区分调用的是哪一个结构的接口实现
	//engine.SimpleEngine{}.Run()

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		// 10 个worker启动goroutine进行数据处理
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
