package scheduler

import "1805/day16/c02-more-user/engine"

// 队列调度器
// 对channel加强控制
type QueueScheduler struct {
	// engine与scheduler交互
	// 只要有人提交一个请求，就传输一个数据
	requestChan		chan engine.Request
	// scheduler与worker交互
	workerChan		chan chan engine.Request
}

// 请求提交
func (s *QueueScheduler) Submit(request engine.Request)  {
	s.requestChan <- request
}

// 添加一个启动函数
func (q * QueueScheduler)Run()  {
	// 初始化两个channel
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)
	// 抛出一个goroutine，对调度相关工作进行维护
	go func() {
		// 添加两个队列进行维护
		var requestQueue	[]engine.Request
		var workerQueue		[]chan engine.Request
		// rabbitmq(第三方的消息队列框架)
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			// 确保两个队列中都元素
			if len(requestQueue) > 0 && len(workerQueue) >0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}

			// 两个channel之间相互独立,接收数据
			// 大多采用select进行监听接收
			select {
			case r := <-q.requestChan:
				requestQueue = append(requestQueue, r)
			case w := <- q.workerChan:
				workerQueue = append(workerQueue,w)
			case activeWorker <- activeRequest:
				// 从队列中丢掉已经使用的数据
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]
				// request<-->channel<-->worker
			}
		}
	}()
}

// 接收worker对应的通道
func (q *QueueScheduler) WorkReady(w chan engine.Request)  {
	q.workerChan <- w
}

// workerChan实现
// 多个goroutine创建多个channel
func (q *QueueScheduler)WorkerChan() chan engine.Request   {
	return make(chan engine.Request)
}