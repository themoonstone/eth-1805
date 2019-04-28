package engine

import (
	"1805/备课/day14/实现/01-单任务版爬虫/c05-cityList-test/fetcher"
	"fmt"
)
// 总控程序
// 启动函数
func Run(seeds ...Request)  {
	// 用一个requests来维护seeds
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	// 判断是否有req
	for len(requests) > 0 {
		// 取出第一个request
		r := requests[0]
		fmt.Printf("request : %v\n", r)
		requests = requests[1:]
		// 调用fetcher,请求数据
		body, err := fetcher.Fetch(r.Url)
		if nil != err {
			// 当前请求出错，去调用下一个
			continue
		}
		// 解析
		parseResult := r.ParseFunc(body)
		// 将result中的request添加到requests中去
		requests = append(requests, parseResult.Requests...)
		// 打印item
		for _, item := range parseResult.Items {
			fmt.Printf("item : %s\n", item)
		}
	}
}