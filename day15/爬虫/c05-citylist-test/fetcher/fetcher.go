package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 根据指定的URL进行数据抓取
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 状态码判断
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error : status code :%v\n", resp.StatusCode)
		return nil, fmt.Errorf("get incorrect status code!", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}