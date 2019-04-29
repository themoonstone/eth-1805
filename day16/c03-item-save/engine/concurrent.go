package engine


// 并发抓取逻辑
import (
	"1805/day16/c03-item-save/model"
)

type ConcurrentEngine struct {
	// 并发处理逻辑中包含一个调度器
	Scheduler Scheduler
	WorkerCount		int		// 并发调用的worker数量
	// item数据传输通道
	ItemChannel		chan interface{}
}

// 总控程序
// 启动函数
func (e *ConcurrentEngine)Run(seeds ...Request)  {
	// 一个负责返回解析结果
	out := make(chan ParseResult)
	// 启动调度器总控函数
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		// 每抛出一个goroutine，创建一个worker
		in := e.Scheduler.WorkerChan()
		createWorker(in, out, e.Scheduler)
	}
	// 必须在in传递并传递之后再提交
	// 需要等到goroutine抛出之后才能正常传递
	for _, r := range seeds {
		// 提交请求
		e.Scheduler.Submit(r)
	}
	// 接收out传递过来的结果
	for {
		result := <- out
		// 打印item
		for _, item := range result.Items {
			if _, ok := item.(model.Profile); ok {
				// 传输有效数据Item
				go func(item interface{}) {
					// 将item传入e.ItemChannel(对应ItemServer()-->out)
					e.ItemChannel <- item
				}(item)
			}
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

// 创建worker
// 调用worker处理进行请求抓取，文本解析
func createWorker(in chan Request, out chan ParseResult, s Scheduler)  {
	go func() {
		for {
			// 告诉scheduler,准备好接收数据了
			s.WorkReady(in)
			// 循环获取请求
			request := <-in
			//fmt.Printf("request : %v\n", request)
			result, err := Worker(request)
			if nil != err {
				continue
			}
			// 向通道回传解析结果
			out <- result
		}
	}()
}

