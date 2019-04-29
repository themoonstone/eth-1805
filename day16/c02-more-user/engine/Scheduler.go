package engine


// 调度器
type Scheduler interface {
	// 请求提交方法
	Submit(request Request)
	// 调度器总控函数
	Run()
	// 接收worker准备好的通道
	WorkReady(worker chan Request)

	// 返回对应的chan request
	WorkerChan() chan Request
}