package utils

import (
	"1805/day14/爬虫/c06-con-worker/engine"
	"1805/day14/爬虫/c06-con-worker/fetcher"
	"github.com/labstack/gommon/log"
)

// 针对请求抓取和方便解析的封装
func Worker(r engine.Request) (engine.ParseResult, error) {

	// 调用fetcher,请求数据
	body, err := fetcher.Fetch(r.Url)
	if nil != err {
		// 当前请求出错，去调用下一个
		log.Printf("the current req failed! %v\n", err)
		return engine.ParseResult{}, err
	}
	// 解析
	parseResult := r.ParseFunc(body)
	return parseResult, nil
}