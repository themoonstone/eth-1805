package parser

import (
	"1805/day16/c03-item-save/fetcher"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err:=fetcher.Fetch(
		"http://www.zhenai.com/zhenghun")
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n", contents)
	// 通过上面的打印，我们可以发现这个前面产生的内容截断
	// 问题出在fetcher/fetcher.go中的determineEncoding上面
	// 所以需要进行修改
}

func TestParseCityList2(t *testing.T) {
	// 因为这里主要是对解析功能进行测试，
	// 所以像请求错误这一类的不应该在里面出现
	// 因此将获得的数据保存在文件中，
	// 然后对文件进行解析
	contents, err := ioutil.ReadFile("cityList.html")
	if nil != err {
		panic(err)
	}

	result := ParseCityList(contents)
	const resultSize = 470
	if len(result. Requests)!=resultSize {
		t.Errorf("result should have %d "+
			"requests; but had %d", resultSize, len(result.Requests))
	}
		if len(result. Items)!=resultSize {
			t.Errorf("result should have %d "+
				"requests; but had %d", resultSize, len(result.Items))
		}
}

