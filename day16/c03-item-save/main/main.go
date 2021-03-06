package main

import (
	"1805/day16/c03-item-save/database"
	"1805/day16/c03-item-save/engine"
	"1805/day16/c03-item-save/scheduler"
	"1805/day16/c03-item-save/zhenai/parser"
)

// 添加编码识别工作
func main() {
	// 区分调用的是哪一个结构的接口实现
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		// 10 个worker启动goroutine进行数据处理
		WorkerCount: 100,
		//ItemChannel: func() chan interface{}{
		//	out := make(chan interface{})
		//	go func() {
		//		for {
		//			item := <- out
		//			fmt.Printf("Item Server : get item:%v\n", item)
		//		}
		//	}()
		//	return out
		//}(),
		ItemChannel: database.ItemServer(),
		}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
