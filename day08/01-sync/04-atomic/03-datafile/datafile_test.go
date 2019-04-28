package _2_datafile

import (
	"math/rand"
	"sync"
	"testing"
)

// 新建文件实例的测试函数
// func NewDataFile(path string, dataLen uint32) (DataFile, error)
func testNew(path string, dataLen uint32, t *testing.T)  {
	t.Logf("init a data file....\n")
	dataFile, err := NewDataFile(path, dataLen)
	if nil != err {
		t.Logf("int a data file failed! %v\n", err)
		t.FailNow()
	}
	if dataFile == nil {
		t.Log("data file is nil...")
		t.FailNow()
	}
	defer dataFile.Close()
}

// 测试读写函数
func testRW(path string, dataLen uint32, t *testing.T)  {
	t.Logf("test write and read...\n")
	dataFile, err := NewDataFile(path, dataLen)
	if nil != err {
		t.Logf("init a data file failed! %v\n", err)
		t.FailNow()
	}
	if dataFile == nil {
		t.Log("data file is nil...")
		t.FailNow()
	}
	defer dataFile.Close()
	var wg sync.WaitGroup
	// wg计数器加5
	wg.Add(5)
	// 设置5个goroutine, 3个写入，两个读取
	for i := 0; i < 3; i++ {
		go func() {
			// 每执行一次，wg计数器减1
			defer wg.Done()
			// 上一个数据块编号
			var prevWSN int64 = -1
			data :=  Data{
				byte(rand.Int31n(256)),
				byte(rand.Int31n(256)),
				byte(rand.Int31n(256)),
			}
			wsn, err := dataFile.Write(data)
			if nil != err {
				t.Logf("write file failed! %v\n", err)
				t.FailNow()
			}

			if wsn <= prevWSN {
				t.Fatalf("Incorrect WSN %d!", wsn)
			}
			prevWSN = wsn
		}()
	}

	// 读取
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			var prevRSN int64 = -1
			rsn, _, err := dataFile.Read()
			if nil != err {
				t.Logf("read the data from file failed! %v\n", err)
				t.FailNow()
			}
			if rsn <= prevRSN {
				t.Fatalf("Incorrect WSN %d!", rsn)
			}
			prevRSN = rsn
		}()
	}

	// wg等待，等到计数器为0，才能正常执行下面的程序
	wg.Wait()
}


func TestIDataFile(t *testing.T)  {
	dataLen := uint32(3)
	// 写入
	// 读取
	t.Run("v1/all", func(t *testing.T) {
		// 新建实例
		path := "data_file_test_new.txt"

		t.Run("New", func(t *testing.T) {
			testNew(path, dataLen, t)
		})
		path1 := "data_file_test.txt"
		t.Run("WriteAndRead", func(t *testing.T) {
			testRW(path1, dataLen, t)
		})
	})

}