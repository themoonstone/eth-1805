package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
type PersonRedis struct {
	ID		int		`db:"id"`
	Name	string	`db:"name"`
	Age		int		`db:"age"`
}
func main() {
	// 连接redis数据库
	conn, err := redis.Dial("tcp", "192.168.1.15:6379")
	if nil != err {
		panic(err)
	}
	defer conn.Close()
	fmt.Print("连接成功...")
	// Do sends a command to the server and returns the received reply.
	var p1 PersonRedis
	p1.ID = 1
	p1.Name = "James"
	p1.Age = 35
	var p2 PersonRedis
	p2.ID = 2
	p2.Name = "James"
	p2.Age = 35
	pers := []PersonRedis{p1,p2}
	reply, err :=conn.Do("hmset", "person", "james", pers)
	//reply, err := conn.Do("get", "go-redis")
	if nil != err {
		panic(err)
	}
	fmt.Printf("type=%T, value=%v\n", reply, reply)
	// 此处要传入的error指的是上面的conn.Do执行结果的error
	res, _ := redis.String(reply, err)
	fmt.Printf("res : %v\n", res)

	reply1, err :=conn.Do("hmget", "person", "james")
	fmt.Printf("type=%T, value=%v\n", reply1, reply1)
	rs, err := redis.Strings(reply1, err)
	if nil != err {
		panic(err)
	}
	fmt.Printf("rs : %v\n", rs)
}
