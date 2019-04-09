package c01_csv

import (
	"encoding/csv"
	"io"
	"os"
)

// csv 文件写入

type Book struct {
	Author string
	Title	string
}
// 图书列表
type Books []Book
// 将csv内容写入文件
func WriteCSVToFile(f *os.File) error  {
	b := Books{
		Book{
			Author:"Astixe",
			Title:"go web 编程",
		},
		Book{
			Author:"许式伟",
			Title:"go 语言编程",
		},
	}
	// UTF-8
	// UTF-8 BOM
	f.WriteString("\xEF\xBB\xBF")
	return b.ToCSV(f)
}

// csv转换函数
func (bs *Books)ToCSV(w io.Writer) error {
	n := csv.NewWriter(w)
	err := n.Write([]string{"Author", "Title"})
	if nil != err {
		return err
	}
	for _, book := range *bs {
		err = n.Write([]string{book.Author, book.Title})
		if nil != err {
			return err
		}
	}
	n.Flush()
	// TODO 解决一下汉字显示乱码的问题
	return n.Error()
}