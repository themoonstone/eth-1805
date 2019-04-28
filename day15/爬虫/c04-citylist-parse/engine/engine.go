package engine

import (
	"eth-1805/day14/爬虫/c03-fetcher/fetcher"
	"fmt"
)

// 总控模块
func Run(seeds ...Request)  {
	// 用一个Requests列表来维护seeds
	var requests []Request
	// 遍历seeds获取每个request
	for _, request := range  seeds {
		requests = append(requests, request)
	}

	// 从任务中获取req
	for len(requests) > 0{
		r := requests[0]
		requests = requests[1:]
		// 调用fetcher，请求数据
		body, err := fetcher.Fetch(r.Url)
		if nil != err {
			continue
		}
		// 解析
		parseResult := r.ParseFunc(body)
		// 将result中的requst添加到requests队列中
		// parseResult.Requests...
		// 将parseResult.Requests中的所有元素追加到requests中，
		// 相当range parseResult.Requests，将每个元素进行追加
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			fmt.Printf("item : %s\n", item)
		}
	}
}