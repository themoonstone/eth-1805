package main
//
//import (
//	"bufio"
//	"github.com/axgle/mahonia"
//	"io"
//	"os"
//	"strings"
//)
//const (
//	SRC_ENCODE = "GBK"
//	DST_ENCODE = "UTF-8"
//)
//// 筛选出包含身份证号信息，存入文件
//func main() {
//	file, _ := os.Open("./kf.txt")
//	defer file.Close()
//
//	useful_f, _ := os.OpenFile("./useful.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
//	defer useful_f.Close()
//	unUseful_f, _ := os.OpenFile("./unuseful.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
//	defer unUseful_f.Close()
//	reader := bufio.NewReader(file)
//	for {
//		lineBytes, _, err := reader.ReadLine()
//		if err == io.EOF {
//			break
//		}
//		srcStr := string(lineBytes)
//		utfStr, _ := ConvertEncoding(srcStr, SRC_ENCODE, DST_ENCODE)
//		infos := strings.Split(utfStr, ",")
//		// 查找是否包含身份证号
//		if len(infos) > 1 && len(infos[1]) == 18 {
//			useful_f.WriteString( utfStr+ "\n")
//		} else {
//			unUseful_f.WriteString(utfStr + "\n")
//		}
//
//	}
//}
//
//// 公共函数
//// 编码解码函数
//// srcStr:待编码的字符中
//// srcEnccoding:原编码格式
//// dstEncoding:目标编码格式
//func ConvertEncoding(srcStr, srcEncoding, dstEncoding string)(dstString string, err error)  {
//	srcDecoder := mahonia.NewDecoder(srcEncoding)
//	dstDecoder := mahonia.NewDecoder(dstEncoding)
//	// 转UTF-8函数调用
//	utfStr := srcDecoder.ConvertString(srcStr)
//	_, dstBytes, err := dstDecoder.Translate([]byte(utfStr), true)
//	if nil != err {
//		return
//	}
//	dstString = string(dstBytes)
//	return
//}