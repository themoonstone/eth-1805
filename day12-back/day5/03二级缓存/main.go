package main

import (
	//导入SDK包（）
	//格式化的输入输出包
	"fmt"
	//操作系统包
	"os"

	//导入mysql数据库驱动，只需要执行其包的初始化方法，不需要使用其API
	_ "github.com/go-sql-driver/mysql"

	//导入第三方包，存储路径在GOPATH下
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
)

/*定义结构体，用于接收数据库的数据*/
type Person struct {
	Name string `db:"name"`
	IdCard  int    `db:"idcard"`
}

/*
错误处理函数
参数：传入错误、出错的场景
只要有错误，就打印错误并暴力退出程序
*/
func HandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)

		//暴力结束程序
		os.Exit(1)
	}
}

/*程序入口*/
func main() {

	//定义字符串变量userInput，用于稍后接收用户输入
	var userInput string

	//源源不断地接收用户输入
	for {
		//提示用户输入命令
		fmt.Println("请输入命令:")

		//阻塞等待用户输入，并存入userInput的内存地址
		fmt.Scan(&userInput)

		//根据用户的具体输入，做不同的响应
		switch userInput {

		//显示所有人员信息
		case "getall":
			GetAllPeople()

		//退出程序
		case "exit":
			//直接跳转到GAMEOVER处
			goto GAMEOVER

		default:
			fmt.Println("什么破命令，fuckoff！")
		}

	}

GAMEOVER:
	//打印并换行
	fmt.Println("GAME OVER")
}

/*获取并打印所有人员信息*/
func GetAllPeople() {

	//先尝试拿缓存
	strs := GetPeopleFromRedis()

	if len(strs) == 0{

		//如果没有拿到数据从MySQL拿取数据
		people := GetPeopleFromMysql()

		//缓存查询结果到Redis
		CachePeople2Redis(people)
	}else{

		//拿到了缓存就直接打印
		fmt.Println(strs)
	}

}

/*
从MySQL核心数据库拿取人员信息
返回值：
people []Person 预定义的Person切片容器
*/
func GetPeopleFromMysql()(people []Person) {
	fmt.Println("GetPeopleFromMysql")

	//连接数据库:mysql=驱动，root=用户名，123456=密码,localhost:3306=地址，mydb=数据库名称，db=数据库对象
	//db, err := sqlx.Open("mysql", "root:123456@tcp(localhost:3306)/mydb")
	db, err := sqlx.Connect("mysql", "root:troytan@itxdl.cn@tcp(192.168.1.15:3306)/eth_1805")
	//处理错误
	HandleError(err,"sqlx.Connect")
	//延时关闭数据库（当前函数结束前）
	defer db.Close()

	//执行SQL语句查询，查询结果转化为[]Person，丢入返回值people的地址
	err = db.Select(&people, "select name,idcard from person")
	//处理错误
	HandleError(err,"db.Select")

	//打印结果
	fmt.Println(people)
	return
}

/*
从Redis缓存获取获取人员信息
返回值：【很多的】人员信息字符串，使用【切片容器[]string】返回
*/
func GetPeopleFromRedis() (strs []string) {
	fmt.Println("GetPeopleFromRedis")

	//连接Redis数据库,通信协议=tcp，访问地址=localhost:6379
	//获得连接或错误
	conn, err := redis.Dial("tcp", "192.168.1.15:6379")
	HandleError(err,"redis.Dial")
	//延时关闭连接
	defer conn.Close()

	//执行redis命令：lrange people 0 -1
	// -1表示列表的最后一个元素
	reply, err := conn.Do("lrange", "people", "0", "-1")
	//尝试将结果转换为[]string
	strs, err = redis.Strings(reply, err)
	fmt.Println("缓存拿取结果", strs, err)

	//返回结果
	return
}

/*缓存查询结果到Redis*/
func CachePeople2Redis(people []Person) {

	//连接数据库并延时关闭
	conn, _ := redis.Dial("tcp", "192.168.1.15:6379")
	defer conn.Close()

	//清除原有缓存：del people
	conn.Do("del","people")

	//遍历people中的每一个Person对象
	//得到每一个Person对象的序号和值，序号不想用，使用下划线除掉
	for _, person := range people {
		//获得待打印的字符串
		personStr := fmt.Sprint(person)

		//执行redis命令：rpush people xxx，把每一个personStr存入列表people
		_, err := conn.Do("rpush", "people", personStr)
		HandleError(err, "@ rpush people "+personStr)
	}

	//执行redis命令：expire people 20，使people在20秒后过期
	_, err := conn.Do("expire", "people", 20)
	HandleError(err, "@ expire people 60")
	fmt.Println("缓存people成功！")
}
