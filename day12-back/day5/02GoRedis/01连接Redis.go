package main

import (
	"github.com/garyburd/redigo/redis"
	"os"
	"fmt"
)

func HandleError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)

		//暴力结束程序
		os.Exit(1)
	}
}

func main() {
	//连接Redis数据库库
	conn, err := redis.Dial("tcp", "192.168.1.15:6379")
	//HandleError(e, "redis.Dial")
	if nil != err {
		panic(err)
	}
	//延时关闭连接
	defer conn.Close()

	//执行Redis命令，获得结果
	//reply, err := conn.Do("Get", "name")
	reply, err := conn.Do("llen", "mlist")
	HandleError(err, "conn.Do Get")

	//结果的原始类型是[]byte
	fmt.Printf("type=%T,value=%v\n", reply, reply)

	//根据具体的业务类型进行数据类型转换
	//ret, _ := redis.String(reply, err)
	ret, _ := redis.Int(reply, err)
	fmt.Println(ret)
}
