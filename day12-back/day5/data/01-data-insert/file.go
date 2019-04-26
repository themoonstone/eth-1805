package _1_data_insert

import "os"

// 判断标记文件是否存在
func IfFileExist(filename string) (exist bool, info string) {
	fileInfo, err := os.Stat(filename)
	if fileInfo != nil && err == nil {
		exist = true
		info = "文件存在，不需要再次存入"
		return
	} else if os.IsNotExist(err) {
		exist = false
		info = "文件不存在，存入mysql"
		return
	} else {
		exist = false
		info = "其它错误"
	}
	return
}