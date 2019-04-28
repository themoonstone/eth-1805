package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"os"
	"bufio"
	"io"
	"strings"
)

/*
·将文本大数据以适当的结构存入MySQL数据库；
·在终端循环输入要查询的姓名，对开放记录进行查询；
·实现精确查询和模糊查询(不加索引)；
·实现精确查询和模糊查询(添加索引)；
·实现内存-数据库的二级缓存策略，并显示每一次查询的时间消耗；
·实现redis-数据库的二级缓存策略，并显示每一次查询的时间消耗
*/

const CACHE_LEN = 2

var (
	kfMap map[string]TimedData
	chanData chan *KfPerson
	db       *sqlx.DB
)

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println("ERROR OCCURED!!!", err, why)
	}
}

//将文本大数据入库
//入库成功后，做一个文件标记，下一次见到标记就不再执行入库操作
func init() {

	//如果数据库已经初始化过了，就直接退出
	exists, _ := CheckIfFileExist("d:/temp/kaifanggood_dbok.mark")
	if exists {
		fmt.Println("信息已入库！")
		return
	}

	//打开数据库
	var err error
	db, err = sqlx.Open("mysql", "root:troytan@itxdl.cn@tcp(192.168.1.15:3306)/kaifang2")
	HandleError(err, "sqlx.Open")
	defer db.Close()
	fmt.Println("数据库已打开")

	//必要时建表
	_, err = db.Exec("create table if not exists kfperson(id int primary key auto_increment,name varchar(20),idcard char(18));")
	HandleError(err, "db.Exec create table")
	fmt.Println("数据表已创建")

	//初始化信号量管道（控制并发数）
	chanData = make(chan *KfPerson, 10000000)

	//开辟协程，源源不断地从数据管道获取信息，插入数据库
	for i := 0; i < 100; i++ {
		go insertKfPerson()
	}

	//打开大数据文件
	file, e := os.Open("d:/temp/kaifanggood.txt")
	HandleError(e, "os.Open")
	defer file.Close()
	reader := bufio.NewReader(file)
	fmt.Println("大数据文本已打开")

	//分批次地读入大数据文本
	for {
		lineBytes, _, err := reader.ReadLine()
		HandleError(err, "reader.ReadLine")
		if err == io.EOF {
			//关闭数据管道
			close(chanData)
			break
		}

		//逐条入库（并发）
		lineStr := string(lineBytes)
		fields := strings.Split(lineStr, ",")
		name, idcard := fields[0], fields[1]

		//抛弃过长的名字
		name = strings.TrimSpace(name)
		if len(strings.Split(name, "")) > 20 {
			fmt.Println("%s名字过长，已经抛弃了")
			continue
		}

		//方案一：开2000万协程，行不通，耗尽了资源，程序崩溃
		//go insertKfPerson(db, &kfPerson)

		//方案二：开有限条协程，从管道中读取数据
		kfPerson := KfPerson{Name: name, Idcard: idcard}
		chanData <- &kfPerson
	}

	//创建一个标记文件，标记数据库已经初始化成功
	fmt.Println("数据初始化成功！")
	_, err = os.Create("d:/temp/kaifanggood_dbok.mark")
	if err == nil {
		fmt.Println("初始化标记文件已创建！")
	}

}

// 信息入库
func insertKfPerson() {
	for kfPerson := range chanData {

		//循环插入至成功为止
		for {
			result, err := db.Exec("insert into kfperson(name,idcard) values(?,?);", kfPerson.Name, kfPerson.Idcard)
			HandleError(err, "db.Exec insert")
			if err != nil {
				//connectex: Only one usage of each socket address is normally permitted
				//db的性能被耗尽，资源已经捉襟见肘了——休眠5秒再来
				<-time.After(5 * time.Second)
			} else {
				if n, e := result.RowsAffected(); e == nil && n > 0 {
					fmt.Printf("插入%s成功！\n", kfPerson.Name)
					break
				}
			}
		}

	}
}

func main() {

	//打开数据库
	db, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/kaifang2")
	HandleError(err, "sqlx.Open")
	defer db.Close()

	//初始化缓存
	kfMap = make(map[string]TimedData, 0)

	//循环查询
	for {
		//循环接收用户想要查询的姓名
		var name string
		fmt.Print("请输入要查询的开房者姓名：")
		fmt.Scanf("%s", &name)

		//用户想退出
		if name == "exit" {
			break
		}

		//查看所有缓存姓名
		if name == "cache" {
			fmt.Printf("共缓存了%d条结果：\n", len(kfMap))
			for key := range kfMap {
				fmt.Println(key)
			}
			continue
		}

		//先查看内存中是否有结果
		//内存中有结果，就直接使用内存中的结果
		if td, ok := kfMap[name]; ok {
			qr := td.(*QueryResult)
			qr.count += 1
			fmt.Printf("查询到%d条结果：\n", len(qr.value))
			fmt.Println(qr.value)
			continue
		}

		//内存中没有，查数据库
		kfpeople := make([]KfPerson, 0)
		e := db.Select(&kfpeople, "select id,name,idcard from kfperson where name like ?;", name)
		HandleError(e, "db.Select")
		fmt.Printf("查询到%d条结果：\n", len(kfpeople))
		fmt.Println(kfpeople)

		//查到的结果丢入内存
		queryResult := QueryResult{value: kfpeople}
		queryResult.cacheTime = time.Now().UnixNano()
		queryResult.count = 1
		kfMap[name] = &queryResult

		//有必要时淘汰一些缓存
		if len(kfMap) > CACHE_LEN {
			delKey := UpdateCache(&kfMap)
			fmt.Printf("%s已经被淘汰出缓存！\n", delKey)
		}
	}

	fmt.Println("GAME OVER!")
}
