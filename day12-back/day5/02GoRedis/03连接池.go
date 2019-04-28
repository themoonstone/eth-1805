package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"strconv"
	"time"
)

func main() {

	//配置并获得一个连接对象的指针
	poolPtr := &redis.Pool{

		//最大活动连接数，0=无限
		MaxActive: 100,

		//最大闲置连接数
		MaxIdle: 20,

		//闲置连接的超时时间
		IdleTimeout: time.Second * 100,

		//定义拨号获得连接的函数
		Dial: func() (redis.Conn, error){
			conn, e := redis.Dial("tcp", "192.168.1.15:6379")
			return conn,e
		}}

	//延时关闭连接池
	defer poolPtr.Close()

	//10个人访问数据库
	for i := 0; i < 10; i++ {
		go getConnFromPool_And_Happy(poolPtr,i)
	}

	//保持主协程存活
	time.Sleep(3 * time.Second)

}

func getConnFromPool_And_Happy(pool *redis.Pool, i int) {
	//通过连接池获得连接
	conn := pool.Get()
	//延时关闭连接
	defer conn.Close()

	//使用连接操作数据
	reply, err := conn.Do("set", "conn"+strconv.Itoa(i), i)
	s, _ := redis.String(reply, err)
	fmt.Println("result:", s)
}
