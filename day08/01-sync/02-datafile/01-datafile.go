package _2_datafile

import (
	"github.com/pkg/errors"
	"io"
	"os"
	"sync"
)

// 数据类型
type Data []byte

// 定义一个并发安全的文件结构
type myDataFile struct {
	f		*os.File		// 文件对象
	fmutex  sync.RWMutex	// 用于文件的读写锁
	woffset	int64			// 写操作要用到的偏移量
	roffset	int64			// 读操作要用到的偏移量
	wmutex	sync.Mutex		// 对写操作的偏移量进行锁定
	rmutex	sync.Mutex		// 对读操作的偏移量进行锁定
	dataLen	uint32			// 数据块大小
}

// interface
// 代表并发安全的数据文件的接口类型
type DataFile interface {
	// 读取一个数据块
	Read()(rsn int64, d Data, err error)
	// 写入一个数据块
	Write(d Data) (wsn int64, err error)
	// 获取最后读取的数据块的序列号
	RSN()	int64
	// 获取最后写入的数据块的序列号
	WSN()	int64
	// 获取数据块长度
	DataLen() uint32
	// 关闭数据文件
	Close() error
}

// 新建数据文件实例函数
func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if nil != err {
		return nil, err
	}
	if dataLen == 0 {
		return nil, errors.New("Invalid data length")
	}
	df := &myDataFile{f:f, dataLen: dataLen}
	return df, nil
}

// 并发安全的读取
func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	// 读取更新偏移量
	var offset int64
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset += int64(df.dataLen)
	df.rmutex.Unlock()
	// 读取一个数据块
	// 获取当前要读取的数据块编号
	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)
	for {
		df.fmutex.RLock()
		_, err = df.f.ReadAt(bytes, offset)
		if nil != err {
			if err == io.EOF {
				df.fmutex.RUnlock()
				continue
			}
			df.fmutex.RUnlock()
			return
		}
		d = bytes
		// 解锁
		df.fmutex.RUnlock()
		return
	}
}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	// 读取并更新写偏移量
	var offset int64
	df.rmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.dataLen)
	df.rmutex.Unlock()

	wsn = offset / int64(df.dataLen)
	var bytes []byte
	if len(d) > int(df.dataLen) {
		// 把多余的数据截断
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}

	// 加锁
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
	return
}

func (df *myDataFile) RSN() int64 {
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.roffset / int64(df.dataLen)
}

func (df *myDataFile) WSN() int64 {
	df.wmutex.Lock()
	defer  df.wmutex.Unlock()
	// 偏移量与数据块大小的关系
	return df.woffset / int64(df.dataLen)
}

// 返回数据块长度
func (df *myDataFile) DataLen() uint32 {
	return df.dataLen
}

// 关闭文件对象
func (df *myDataFile)Close() error {
	if nil == df.f {
		return nil
	}
	return df.f.Close()
}