package parser
// 城市列表解析器
import (
	"eth-1805/day14/爬虫/c04-citylist-parse/engine"
	"regexp"
)

// 通过正则提取城市和URL
// <a href="http://www.zhenai.com/zhenghun/binhaixin" data-v-5e16505f>滨海新</a>

// 打印城市和URL
/*
	args :
		contents:http请求返回的响应内容
*/
// 城市列表解析规则
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

// 输入:utf-8的文本
// 返回值：解析器结果
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe) //匹配规则

	matches := re.FindAllSubmatch(contents, -1)
			result := engine.ParseResult{}
			for _, m := range matches {
			result.Items = append(result.Items, string(m[2]))
			result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			// parseFunc：解析每个城市的用户列表的解析器
			ParseFunc: engine.NilParser,
		})
	}
	return result
}
