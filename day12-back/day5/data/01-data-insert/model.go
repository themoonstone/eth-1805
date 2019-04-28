package _1_data_insert

import "github.com/jmoiron/sqlx"

// 将要存放数据的信息表结构
type Person struct {
	ID		int			`db:"id"`
	Name	string		`db:"name"`
	IdCard	string		`db:"idcard"`
}

// 数据库实例结构
type DB struct {
	Db 		*sqlx.DB
	ChanData		chan Person
}