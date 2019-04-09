package c01_csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

// CSV 文件读取
func ReadCSV(b io.Reader) ([]Book, error) {
	fmt.Printf("ReadCSV exec! \n")
	r := csv.NewReader(b)
	var movies []Book // 存储电影信息的切片
	// 判断文件本身是否能正常读取
	_, err := r.Read()
	if nil != err && err != io.EOF {
		panic(err)
		return nil, err
	}
	fmt.Printf("the file is ready to get data\n")

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		// title
		title := record[0]
		rs := strings.Split(record[0], ",")
		fmt.Printf("length of rs : %v\n", len(rs))

		// director
		director := record[1]
		m := Book{Title:title, Author:director}
		movies = append(movies, m)
	}
	return movies, nil
}