package main

import (
	// 此处添加别名是因为不允许有数字开头的包名去调用函数
	seq "eth-1805/day12-back/day5/data/02-data-query"
	"fmt"
	"time"

	// 加载mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// 通过init调用数据存储函数
// 只执行一次(考虑用一下once.do)
// 考虑添加一个判断，通过一个标记文件来判断是否已经插入过
// 因为有可能会多次启动，但数据库已经存入数据了，不能多次存入
var (
	db seq.DB
	// 缓存定义
	// key:name

	CacheData	map[string] seq.TimeData
)

const (
	// 标记文件名称
	markFile 	= "./kaifang.mark"
)
//func init() {
//	// 读取文件
//	// 先判断一下标记文件是否存在
//	exists, _ := seq.IfFileExist(markFile)
//	if exists {
//		log.Info("数据已插入!")
//		return
//	}
//
//	// 打开数据库
//	db.OpenDb()
//	defer db.Db.Close()
//	// 启动channel管道
//	db.ChanData = make(chan seq.Person, 100000)
//	// 建表
//	_, err := db.Db.Exec("create table if not exists person(id int primary key auto_increment, name varchar(20), idcard char(18));")
//	if nil != err {
//		seq.HandleError(err, "create table!")
//	}
//	log.Info("数据表创建成功\n")
//	// 读取文件数据
//	f, err := os.Open("./kf.txt")
//	if err != nil {
//		seq.HandleError(err, "open source file!")
//	}
//
//	// 开启多个goroutine
//	for i:=0; i < 100; i++ {
//		// 存入数据库
//		go db.SaveDataToMysql()
//	}
//
//	reader := bufio.NewReader(f)
//	for {
//		line, _, err := reader.ReadLine()
//		if nil != err {
//			if err == io.EOF {
//				// 关闭通道
//				close(db.ChanData)
//				break
//			}
//			seq.HandleError(err, "read data from file")
//		}
//		utfStr, _ := seq.ConvertEncoding(string(line), seq.SRC_ENCODE, seq.DST_ENCODE)
//		// 解析数据，将[]byte解析为Person结构体
//		datas := strings.Split(utfStr, ",")
//		// 筛选掉名称太长的信息(data[0]小于20)
//		if len(datas) >= 2 && len(datas[0]) <= 20 && len(datas[1]) == 18 {
//			p := seq.Person{}
//			p.Name = datas[0]
//			p.IdCard = datas[1]
//			db.ChanData <- p
//		}
//	}
//	// 数据存储完成之后，生成标记文件
//	_, err = os.Create(markFile)
//	if nil != err {
//		seq.HandleError(err, "open mark file!")
//	}
//}

//func main() {
//	fmt.Printf("exec...\n")
//	db.OpenDb()
//	defer db.Db.Close()
//
//	// 循环查询
//	for {
//		var name string
//		fmt.Printf("你想找谁? \n")
//		// 获取用户输入的名称
//		fmt.Scanf("%s\n", &name)
//		// 退出
//		if name == "exit" {
//			break
//		}
//		// 直接在数据库中查询
//		// 因为是通过name查询，所以可能有重复
//		var result []seq.Person = make([]seq.Person, 0)
//		err := db.Db.Select(&result,"select id,name,idcard from person where name like ?;", name)
//		if nil != err {
//			seq.HandleError(err, "sql.query!")
//		}
//		for i := 0; i < len(result); i++ {
//			fmt.Printf("person : %v\n", result[i])
//		}
//	}
//}

func main() {
	fmt.Printf("exec...\n")
	db.OpenDb()
	defer db.Db.Close()
	// 初始化缓存结构
	CacheData = make(map[string] seq.TimeData)
	// 循环查询
	for {
		var name string
		fmt.Printf("你想找谁? \n")
		// 获取用户输入的名称
		fmt.Scanf("%s\n", &name)
		// 退出
		if name == "exit" {
			break
		}
		// 因为是通过name查询，所以可能有重复
		var result []seq.Person = make([]seq.Person, 0)
		// 先检查缓存中是否有数据，如果有，直接在缓存中查找
		if res, ok := CacheData[name]; ok{
			// 实现了接口
			rs := res.(*seq.QueryResult)
			rs.Counts++
			for i := 0; i < len(rs.Result); i++ {
				fmt.Printf("person : %v\n", rs.Result[i])
			}
			continue
		}
		// 缓存中没有
		// 直接在数据库中查询
		err := db.Db.Select(&result,"select id,name,idcard from person where name like ?;", name)
		if nil != err {
			seq.HandleError(err, "sql.query!")
		}
		for i := 0; i < len(result); i++ {
			fmt.Printf("person : %v\n", result[i])
		}

		// 数据库中查询的结果放入缓存
		qr := seq.QueryResult{
			Result:result,
			JoinTime:time.Now().UnixNano(),
			Counts:1,// 刚刚放入内存，所以只有一次
		}
		// 将qr添加到缓存中
		CacheData[name] = &qr
		// 执行缓存淘汰，因为一个name可能包含多条数据，所以最多放两条
		if len(CacheData) > seq.CACHE_LENGTH {
			// 执行淘汰策略
			seq.UpdateCache(&CacheData)
		}
	}
	fmt.Println("完成...")
}
