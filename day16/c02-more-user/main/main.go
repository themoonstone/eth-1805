package main

import (
	"1805/day16/c02-more-user/engine"
	"1805/day16/c02-more-user/scheduler"
	"1805/day16/c02-more-user/zhenai/parser"
)

// 添加编码识别工作
func main() {
	// 区分调用的是哪一个结构的接口实现
	//engine.SimpleEngine{}.Run()

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		// 10 个worker启动goroutine进行数据处理
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
