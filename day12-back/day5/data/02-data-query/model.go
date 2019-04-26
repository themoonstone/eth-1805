package data_project_seq

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

// 加入缓存的时间接口定义
type TimeData interface {
	GetCacheTime()		int64
}

// 查询结果
type QueryResult struct {
	// 结果集
	Result		[]Person
	// 加入缓存时间
	// 对于结构体的内部变量，尽量用小写
	JoinTime	int64
	// 访问次数
	Counts		int
}

func (qr *QueryResult) GetCacheTime() int64 {
	return qr.JoinTime
}