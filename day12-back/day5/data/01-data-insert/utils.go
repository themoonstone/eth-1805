package _1_data_insert

import "github.com/axgle/mahonia"
const (
	SRC_ENCODE = "GBK"
	DST_ENCODE = "UTF-8"
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
