package parser

import (
	"1805/day14/爬虫/c06-con-worker/engine"
	"regexp"
)

// 城市列表解析器
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	// -1 表示所有内容
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 10
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url: string(m[1]),
			// 为什么不直接调用nil，因为调用nil会出错
			ParseFunc: ParseCity,
		})
		limit--
		if limit == 0{
			break
		}
	}
	return result
}