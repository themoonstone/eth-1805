package scheduler

import "1805/day14/爬虫/c07-con-scheduler/engine"

// 简单调度器
// 多个goroutine共同使用一个channel

// 简单调度器结构
type SimpleScheduler struct {
	WorkerChannel		chan engine.Request
}

// 请求提交
func (s *SimpleScheduler) Submit(request engine.Request)  {
	go func() {s.WorkerChannel <- request}()
}

// 内部配置
func (s *SimpleScheduler) ConfigWorkerChannel(in chan engine.Request)  {
	s.WorkerChannel = in
}