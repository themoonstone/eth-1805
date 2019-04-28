package parser

import (
	"1805/备课/day14/实现/01-单任务版爬虫/c06-city-parse/engine"
	"fmt"
	"regexp"
)

// 城市解析器
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
// <th><a href="http://album.zhenai.com/u/1379327440" target="_blank">随缘</a></th>
// <th><a href="http://album.zhenai.com/u/1358992404" target="_blank">白雪王子</a></th>
func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	// -1 表示所有内容
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		for i := 0; i < len(m); i++ {
			fmt.Printf("m[%d] : %s ",i, m[i])
		}
		fmt.Println()
		result.Items = append(result.Items,"User "+ string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url: string(m[1]),
			// 为什么不直接调用nil，因为调用nil会出错
			ParseFunc: engine.NilParser,
		})
	}
	return result
}