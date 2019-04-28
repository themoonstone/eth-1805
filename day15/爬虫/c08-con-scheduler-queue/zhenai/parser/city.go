package parser

import (
	"1805/day14/爬虫/c08-con-scheduler-queue/engine"
	"regexp"
)

// 城市解析器
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
// <th><a href="http://album.zhenai.com/u/1358992404" target="_blank">白雪王子</a></th>
func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	// -1 表示所有内容
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items,"User "+ name)
		result.Requests = append(result.Requests,engine.Request{
			Url: string(m[1]),
			// 此处改为闭包调用ParseUser(...)
			ParseFunc: func(bytes []byte) engine.ParseResult {
				return ParseUser(bytes, name)
			},
		})
	}
	return result
}

// 城市：470+1
// parseCity:用户列表打印用户名470*20=9400
// userInfo : 9400
// 19271左右