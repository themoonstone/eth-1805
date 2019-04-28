package main
import (
	"fmt"
	// 不去调用里面的任何API，紧紧执行一下mysql中的init方法
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
	// 只加载，不引用
	import _ "github.com/go-sql-driver/mysql"
/	另外操作
	import sql "github.com/go-sql-driver/mysql"
	golang 包管理工具
	godep
	glide
*/


// 人员结构
// mysql端的person table
//type Person struct {
//	ID		int		`db:"id"`
//	Name	string	`db:"name"`
//	Age		int		`db:"age"`
//}

func main() {
	// 连接数据库
	// TODO tcp : 协议说明
	/*
		username:password@tcp(host:port)/db_name
		username:mysql用户名
		password:mysql密码
		tcp:协议
		host:主机名
		port:端口号
		db_name:数据库名称
	*/
	// func Open(driverName, dataSourceName string) (*DB, error) {}
	//sqlx.Open("mysql","username:password@tcp(host:port)/db_name")
	db, err := sqlx.Open("mysql","root:troytan@itxdl.cn@tcp(192.168.1.15:3306)/eth_1805")
	defer db.Close()
	if nil != err {
		panic(err)
	}
	fmt.Println("打开数据库成功...")
	var ps []Person
	//for _,p := range ps {
	//	result, err := db.Exec("insert into person(name, age) value(?,?);",p.Name,p.Age)
	//
	//}
	// 插入数据
	result, err := db.Exec("insert into person(name, age) value(?,?),(?,?);","kobe",35,"Smith",30)
	if nil != err {
		panic(err)
	}
	rowAffected, _ := result.RowsAffected()
	fmt.Printf("受影响的行数:%v\n", rowAffected)
	lastInsertedId, _ := result.LastInsertId()
	fmt.Printf("最后插入的ID:%v\n", lastInsertedId)
	// 查询数据
	// err := db.Select(&pps, "SELECT * FROM person")
	// 定义切片接收查询结果

	err = db.Select(&ps, "select * from person")
	if nil != err {
		fmt.Printf("查询出错:%v\n", err)
		return
	}
	fmt.Printf("result : %v\n", ps)
}