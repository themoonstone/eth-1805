package data_project_seq

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"os"
	"time"
)

// sql操作

// 初始化数据库连接
// driverName, dataSourceName string
func (db *DB) OpenDb() {
	driverName := "mysql"
	dataSourceName := "root:troytan@itxdl.cn@tcp(192.168.1.15:3306)/kaifang2"
	sqlDb, err := sqlx.Open(driverName, dataSourceName)

	if nil != err {
		HandleError(err, "open database")
		// 数据库打开失败，直接退出程序
		os.Exit(1)
	}
	db.Db = sqlDb
	log.Info("数据库打开成功")
}

// 将从文本中获取到的数据存入mysql中
// 考虑的问题
	// 1. 文件传输方式
	// 2. 并发操作
func (db *DB)SaveDataToMysql() {
	for  person := range db.ChanData {
		// 循环插入直到成功
		for {
			res, err := db.Db.Exec("insert into person(name, idcard) values(?,?)", person.Name, person.IdCard)
			if nil != err {
				//  Only one usage of each socket address (protocol/network address/port) is normally permitted., insert table failed!\n"
				// db性能被耗尽，休眠5秒接着继续
				time.After(5 * time.Second)
				HandleError(err, "insert table failed!")
			} else {
				n, err := res.RowsAffected()
				if nil != err {
					HandleError(err, "affected is not equal expected!")
				}
				if n > 0 && err == nil {
					break
				}
			}
		}
	}
}