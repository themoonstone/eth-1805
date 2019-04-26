package main

import (
	"bufio"
	// 加载mysql驱动
	_ "github.com/go-sql-driver/mysql"
	seq "eth-1805/day12-back/day5/data/01-data-insert"
	"fmt"
	"github.com/labstack/gommon/log"
	"io"
	"os"
	"strings"
)

// 通过init调用数据存储函数
// 只执行一次(考虑用一下once.do)
// 考虑添加一个判断，通过一个标记文件来判断是否已经插入过
// 因为有可能会多次启动，但数据库已经存入数据了，不能多次存入
var (
	db seq.DB
)

const (
	// 标记文件名称
	markFile 	= "./kaifang.mark"
)

// 将数据从文件中读取出来，插入到mysql数据中
func init() {
	// 读取文件
	// 先判断一下标记文件是否存在
	exists, _ := seq.IfFileExist(markFile)
	// 标记文件已存在
	if exists {
		log.Info("数据已插入!")
		return
	}

	// 打开数据库
	db.OpenDb()
	defer db.Db.Close()
	// 启动channel管道
	db.ChanData = make(chan seq.Person, 100000)
	// 建表
	_, err := db.Db.Exec("create table if not exists person(id int primary key auto_increment," +
		" name varchar(20), idcard char(18));")
	if nil != err {
		seq.HandleError(err, "create table!")
	}
	log.Info("数据表创建成功\n")
	// 读取文件数据
	f, err := os.Open("./kf.txt")
	if err != nil {
		seq.HandleError(err, "open source file!")
	}

	// 开启多个goroutine
	for i:=0; i < 100; i++ {
		// 存入数据库
		go db.SaveDataToMysql()
	}

	reader := bufio.NewReader(f)
	for {
		// 按行读取
		line, _, err := reader.ReadLine()
		if nil != err {
			if err == io.EOF {
				// 关闭通道
				close(db.ChanData)
				break
			}
			seq.HandleError(err, "read data from file")
		}
		// utf8编码转换
		utfStr, _ := seq.ConvertEncoding(string(line), seq.SRC_ENCODE, seq.DST_ENCODE)
		// 解析数据，将[]byte解析为Person结构体
		datas := strings.Split(utfStr, ",")
		// 筛选掉名称太长的信息(data[0]小于20)
		if len(datas) >= 2 && len(datas[0]) <= 20 && len(datas[1]) == 18 {
			p := seq.Person{}
			p.Name = datas[0]
			p.IdCard = datas[1]
			db.ChanData <- p
		}
	}
	// 数据存储完成之后，生成标记文件
	_, err = os.Create(markFile)
	if nil != err {
		seq.HandleError(err, "open mark file!")
	}
}

func main() {
	fmt.Printf("exec...\n")
}