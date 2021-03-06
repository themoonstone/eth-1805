package engine


// 并发抓取逻辑
import (
	"1805/day14/爬虫/c06-con-worker/engine/utils"
	"fmt"
)

type ConcurrentEngine struct {
	// 并发处理逻辑中包含一个调度器
	Scheduler Scheduler
	WorkerCount		int		// 并发调用的worker数量
}

// 调度器
type Scheduler interface {
	// 请求提交方法
	Submit(Request)
}
// 总控程序
// 启动函数
func (e *ConcurrentEngine)Run(seeds ...Request)  {
	for _, r := range seeds {
		// 提交请求
		e.Scheduler.Submit(r)
	}
	// 新建两个 通道
	// 一个负责传递请求
	in := make(chan Request)
	// 一个负责返回解析结果
	out := make(chan ParseResult)

	for i := 0; i < e.WorkerCount; i++ {
		// 每抛出一个goroutine，创建一个worker
		createWorker(in, out)
	}
	// 接收out传递过来的结果
	for {
		result := <- out
		// 打印item
		for _, item := range result.Items {
			fmt.Printf("item : %s\n", item)
		}
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
			result, err := utils.Worker(request)
			if nil != err {
				continue
			}
			out <- result
		}
	}()
}

