package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
	"os"

	// 不去调用里面的任何API，紧紧执行一下mysql中的init方法
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)

/*
	1. mysql中存放大量数据
	2. 直接能mysql中查询数据
	3. 添加redis作为缓存层
	4. 在第进行查询的时候，先判断redis中是否有对应的结果，
		1. 如果redis中有，直接从redis中查询
		2. 如果redis没有包含对应的数据，进入mysql获取对应的结果，存入redis缓存中
*/

// 人员结构
// mysql端的person table
type Person struct {
	ID		int		`db:"id"`
	Name	string	`db:"name"`
	Age		int		`db:"age"`
}

func main() {
	var inputName string
	// 从用户端接收查询条件
	for {
		fmt.Println("input your query condition please:")
		fmt.Scan(&inputName)
		switch inputName {
		case "exit":
			// 用户主动结束
			os.Exit(1)
		}
		// 通过inputName获取所有与之相关的记录
		getData(inputName)

	}

}

// 获取数据
func getData(input_name string) []Person {
	// 获取数据
	// 判断redis
	strs := getInfoFromRedis(input_name)
	fmt.Printf("strs : %v, length : %v\n",strs[0], len(strs))

	if len(strs) == 0 {
		// redis中不存在相关的数据
		// 从mysql中获取
		fmt.Println("从mysql中获取...")
		ps := getInfoFromMysql(input_name)
		// 存入redis
		fmt.Println("缓存到redis中...")
		cacheInfoToRedis(input_name, ps)
	}
	fmt.Printf("strs : %v\n", strs)
	return nil
}

// 从redis中获取数据操作
func getInfoFromRedis(input_name string) []string {
	// 连接redis数据库
	conn, err := redis.Dial("tcp", "192.168.1.15:6379")
	if nil != err {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("连接成功...")
	reply, err := conn.Do("hmget", "emp", input_name)
	if nil != err {
		panic(err)
	}
	fmt.Printf("%T: %v\n", reply, reply)
	rs, err := redis.Strings(reply, err)
	if nil != err {
		panic(err)
	}
	return rs
}

// 从mysql中获取数据操作
func getInfoFromMysql(input_name string) []Person {
	db, err := sqlx.Open("mysql","root:troytan@itxdl.cn@tcp(192.168.1.15:3306)/eth_1805")
	defer db.Close()
	if nil != err {
		log.Panicf("open the database failed! %v\n", err)
	}
	fmt.Println("打开数据库成功...")
	var ps []Person
	err = db.Select(&ps, "select * from person where  name = ?", input_name)
	if nil != err {
		fmt.Printf("查询出错:%v\n", err)
		return nil
	}
	fmt.Printf("data : %v\n", ps)
	return ps
}

// 将mysql中的数据缓存到redis中
func cacheInfoToRedis(input_name string, persons []Person)  {
	/*
		hmget person james
		1) "[{0 James 35} {0 James 35}]"
	*/
	// 连接redis数据库
	conn, err := redis.Dial("tcp", "192.168.1.15:6379")
	if nil != err {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("连接成功...")
	// 使用HashMap
	_, err = conn.Do("hmset", "emp", input_name, persons)
	if nil != err {
		panic(err)
	}
}