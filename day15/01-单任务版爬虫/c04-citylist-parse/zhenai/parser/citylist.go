package parser

import (
	"1805/备课/day14/实现/01-单任务版爬虫/c04-citylist-parse/engine"
	"regexp"
)

// 城市列表解析器
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func ParseCityList(contents []byte) engine.ParseResult {
	// {linkContent:"淮北",linkURL:"http://m.zhenai.com/zhenghun/huaibei"}
	re := regexp.MustCompile(cityListRe)
	// -1 表示所有内容
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url: string(m[1]),
			// 为什么不直接调用nil，因为调用nil会出错
			ParseFunc: engine.NilParser,
		})
	}
	return result
}