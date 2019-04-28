
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
type Person struct {
	ID		int		`db:"id"`
	Name	string	`db:"name"`
	Age		int		`db:"age"`
}

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
	// 开启事务
	tx, _ := db.Begin()
	// 插入数据
	r1, e1 := tx.Exec("insert into person(name, age) value(?,?);","kobe",35)
	_, e2 := tx.Exec("delete from person where name = ?;","jordan")
	_, e3 := tx.Exec("update  person set name= ? where name = ?;","James","Kobe")

	// 任何一个执行语句出错，都要回滚
	if e1 != nil || e2 != nil || e3 != nil {
		fmt.Printf("事务执行失败:e1/e2/e3", e1,e2, e3)
		// 回滚
		tx.Rollback()
	} else {
		// 提交事务
		tx.Commit()
		rowAffected, _ := r1.RowsAffected()
		fmt.Printf("受影响的行数:%v\n", rowAffected)

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


}