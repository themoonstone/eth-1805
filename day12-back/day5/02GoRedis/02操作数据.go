package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"os"
)

func HandleError02(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

/*
rpush mlist 3 // 从右侧向列表mlist追加元素3
lrange mlist 0 -1 //从头看到尾
*/

func main() {

	conn, _ := redis.Dial("tcp", "192.168.1.15:6379")
	defer conn.Close()

	//reply, err := conn.Do("SET", "myname", "你妹")
	//reply, err := conn.Do("setex", "myname", "60","bill")
	//reply, err := conn.Do("persist", "myname")
	//reply, err := conn.Do("mset", "age", "60","gender","male")
	reply, err := conn.Do("hmset", "bangzhu", "name", "jobs", "age",10000)
	reply, err = conn.Do("hgetall", "bangzhu")
	//reply, err := conn.Do("rpush", "mydearlist",11,22,33)
	//reply, err := conn.Do("lrange", "fucklist",0,-1)
	//reply, err := conn.Do("sadd", "mydearset", 1, 2, 3)
	//reply, err := conn.Do("smembers", "mydearset")
	//reply, err := conn.Do("zadd", "mz",10,"bill",9,"bangzhu",8,"zuckberg")
	//reply, err := conn.Do("zrange", "mz",0,-1)
	//reply, err := conn.Do("zrange", "mz",0,-1)

	fmt.Println("原始结果", reply, err)

	//具体转为何种数据类型，根据业务的实际需要
	ret, _ := redis.Strings(reply, err)
	//ret, _ := redis.String(reply, err)
	//ret, _ := redis.Int(reply, err)
	fmt.Println(ret)
}
