package scheduler

import "1805/day16/c02-more-user/engine"

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

// workerChan实现
func (s *SimpleScheduler)WorkerChan() chan engine.Request   {
	return s.WorkerChannel
}

// Run函数实现
func (s *SimpleScheduler) Run()  {
	s.WorkerChannel = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkReady(chan engine.Request)  {

}