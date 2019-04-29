package parser

import (
	"1805/day16/c02-more-user/engine"
	"1805/day16/c02-more-user/model"
	"regexp"
	"strings"
)

// 用户信息解析
// 因为每次正则匹配都需要compile，比较耗时间，所以在这里通过全局变量采用预编译的方式来进行
// <div class="des f-cl" data-v-3c42fade>白山 | 39岁 | 大学本科 | 离异 | 160cm | 3000元以下</div>
var userInfoRe = regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>([^<]+)</div>`)
//// 选取名称
//var nameRe = regexp.MustCompile(`<h1 class="nickName"[^>]+>([^<]+)</h1>`)
// 用户信息解析函数
func ParseUser(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	infoMatches := userInfoRe.FindSubmatch(contents)
	// 替换生成切片
	userInfo := strings.Split(string(infoMatches[1])," | ")
	//fmt.Printf("userInfo : %v\n", userInfo)
	if (len(userInfo) == 6) {
		profile.WorkPosition = userInfo[0]
		profile.Age = userInfo[1]
		profile.Education = userInfo[2]
		profile.Marriage = userInfo[3]
		profile.Height = userInfo[4]
		profile.Income = userInfo[5]
	}
	// 名称赋值
	profile.Name = name
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}