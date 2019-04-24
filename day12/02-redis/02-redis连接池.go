package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"sync"
	"time"
)

// 连接池的管理
// 基本思想：创建一个连接池，放入连接
// 连接池
// 总数
// 空闲连接数
// 活动连接
var wg sync.WaitGroup

func main() {
	// 返回一个连接对象的指针
	poolStr := &redis.Pool{
		// 最大闲置连接数
		MaxIdle: 20,
		// 最大活动连接数， 0表示没有限制
		MaxActive: 100,
		// 闲置连接的超时时间
		IdleTimeout: 100 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, e := redis.Dial("tcp", "192.168.1.15:6379")
			return conn,e
		},
	}
	wg.Add(10)
	defer poolStr.Close()

	// 多个goroutine访问数据库
	for i := 0; i < 10; i++ {
		go getConnFromPoolAndOperate(poolStr, i)
	}
	wg.Wait()
}

// 获取连接池中的连接，进行redis-golang数据交互
func getConnFromPoolAndOperate(pool *redis.Pool, i int)  {
	// 通过连接池进行获取
	conn := pool.Get()
	defer conn.Close()
	// 数据交互
	reply, _ := conn.Do("set", "redis-pool", "conn-"+strconv.Itoa(i))
	fmt.Printf("type: %T, reply : %v\n", reply, reply)
	wg.Done()
}