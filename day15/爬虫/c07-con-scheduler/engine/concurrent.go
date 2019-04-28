package engine




// 并发抓取逻辑
import (
	"fmt"
)

type ConcurrentEngine struct {
	// 并发处理逻辑中包含一个调度器
	Scheduler Scheduler
	WorkerCount		int		// 并发调用的worker数量
	InChan			chan Request
}

// 总控程序
// 启动函数
func (e *ConcurrentEngine)Run(seeds ...Request)  {

	// 新建两个 通道
	// 一个负责传递请求
	in := make(chan Request)
	// 一个负责返回解析结果
	out := make(chan ParseResult)
	// 将in送给workerChannel
	e.Scheduler.ConfigWorkerChannel(in)

	for i := 0; i < e.WorkerCount; i++ {
		// 每抛出一个goroutine，创建一个worker
		createWorker(in, out)
	}
	// 必须在in传递并传递之后再提交
	// 需要等到goroutine抛出之后才能正常传递
	for _, r := range seeds {
		// 提交请求
		e.Scheduler.Submit(r)
	}
	var itemAccount int = 0
	// 接收out传递过来的结果
	for {
		result := <- out
		// 打印item
		for _, item := range result.Items {
			itemAccount++
			fmt.Printf("id:%#d, item : %s\n",itemAccount,item)
			//if _, ok := item.(model.Profile);ok {
			//	fmt.Printf("id:%#d, item : %s\n",itemAccount,item)
			//}
		}
		// 470
		// 将获取到的result中解析出的请求进行提交
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

// 创建worker
// 调用worker处理进行请求抓取，文本解析
func createWorker(in chan Request, out chan ParseResult)  {
	go func() {
		for {
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

