package data_project_seq

import (
	"fmt"
	"github.com/axgle/mahonia"
)
const (
	SRC_ENCODE = "GBK"
	DST_ENCODE = "UTF-8"
	CACHE_LENGTH = 2
)

// 公共函数
// 编码解码函数
// srcStr:待编码的字符中
// srcEnccoding:原编码格式
// dstEncoding:目标编码格式
func ConvertEncoding(srcStr, srcEncoding, dstEncoding string)(dstString string, err error)  {
	srcDecoder := mahonia.NewDecoder(srcEncoding)
	dstDecoder := mahonia.NewDecoder(dstEncoding)
	// 转UTF-8函数调用
	utfStr := srcDecoder.ConvertString(srcStr)
	_, dstBytes, err := dstDecoder.Translate([]byte(utfStr), true)
	if nil != err {
		return
	}
	dstString = string(dstBytes)
	return
}

// 缓存淘汰
// 到达上限之后，每次淘汰1/3
// 淘汰策略：
// 淘汰加入时间更早的数据
func UpdateCache(cache *map[string]TimeData) {
	// 将缓存中所有key(name)代表的数据加入到slice 中
	var tms []int64
	for _, timeData := range *cache {
		tms = append(tms, timeData.GetCacheTime())
	}
	fmt.Printf("before sorted : %v, len : %d\n", tms, len(tms))
	// 对加入时间进行排序
	tmSort(tms)
	// 获取到要删除的时间的索引,取1/3
	delTmIndex := len(tms)/3
	fmt.Printf("after sorted : %v, len : %d\n", tms, len(tms))
	for key, timeData := range *cache {
		// 小于这个索引对应的时间，就删除掉
		if timeData.GetCacheTime() < tms[delTmIndex] {
			delete(*cache, key)
		}
	}
	fmt.Println("淘汰成功...")
}
// 选择排序
func tmSort(tms []int64)  {
	// 7
	for i := 0; i < len(tms) - 1; i++ {
		// 1
		for j := i+1; j < len(tms); j++ {
			if tms[i] > tms[j] {
				// 交换，保证tms[i]始终是最小的那一个
				tms[i], tms[j] = tms[j], tms[i]
			}
		}
	}
}
