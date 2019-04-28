package regexp

import (
	"fmt"
	"regexp"
)

// 通过正则提取城市和URL
// <a href="http://www.zhenai.com/zhenghun/binhaixin" data-v-5e16505f>滨海新</a>

// 打印城市和URL
/*
	args :
		contents:http请求返回的响应内容
*/
func PrintCityList(contents []byte)  {
	// 	re := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[a-z0-9]+" data-v-5e16505f>滨海新</a>`) //匹配规则
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`) //匹配规则

	//matches := re.FindAll(contents, -1)
	matches := re.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		fmt.Printf("match : %s, URL: %s, cityName:%s\n",match, match[1], match[2])
	}
}
