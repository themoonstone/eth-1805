package engine


// 调度器
type Scheduler interface {
	// 请求提交方法
	Submit(request Request)
	// 内部通道配置
	ConfigWorkerChannel(chan Request)
}