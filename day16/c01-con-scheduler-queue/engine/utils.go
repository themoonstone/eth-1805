package engine

import (
	"1805/day16/c01-con-scheduler-queue/fetcher"
	"github.com/labstack/gommon/log"
)

// 针对请求抓取和方便解析的封装
func Worker(r Request) (ParseResult, error) {

	// 调用fetcher,请求数据
	body, err := fetcher.Fetch(r.Url)
	if nil != err {
		// 当前请求出错，去调用下一个
		log.Printf("the current req failed! %v\n", err)
		return ParseResult{}, err
	}
	// 解析
	parseResult := r.ParseFunc(body)
	return parseResult, nil
}