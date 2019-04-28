package engine

import "fmt"

type SimpleEngine struct {

}
// 单线程抓取

// 总控程序
// 启动函数
func (s *SimpleEngine)Run(seeds ...Request)  {
	// 用一个requests来维护seeds
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	// 判断是否有req
	for len(requests) > 0 {
		// 取出第一个request
		r := requests[0]
		//fmt.Printf("request : %v\n", r)
		requests = requests[1:]
		// 调用worker
		parseResult, err := Worker(r)
		if nil != err {
			continue
		}
		// 将result中的request添加到requests中去
		requests = append(requests, parseResult.Requests...)
		// 打印item
		for _, item := range parseResult.Items {
			fmt.Printf("item : %s\n", item)
		}
	}
}