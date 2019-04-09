package _2_monster

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"io/ioutil"
)

// 妖怪
type Monster struct {
	Name		string
	Age			int
	Skill		string
}

// 实现Monster结构体存储到文件中
// TODO 考虑什么时候用函数，什么时候用方法
func (m *Monster)Store() bool {
	// 序列化
	b, err := json.Marshal(m)
	if nil != err {
		log.Errorf("marshal struct to []byte failed! %v\n", err)
		return false
	}
	// 存储
	file_name := "monster.ser"
	// 1. 打开
	//f , err := os.Open(file_name)
	//if nil != err {
	//	log.Errorf("open file %s failed! %v\n",file_name,  err)
	//	return false
	//}
	//_, err = f.Write(b)
	//if nil != err {
	//	log.Errorf("write the struct to  file %s failed! %v\n",file_name,  err)
	//	return false
	//}
	err = ioutil.WriteFile(file_name, b, 0666)
	if nil != err {
		//log.Errorf("write file failed")
		log.Printf("%v\n", 	errors.New("write file failed"))
	}
	return true
}

// 读取内容
func (m *Monster)ReadFromFile() bool {
	// 1. 读取文件内容
	file_name := "monster.ser"
	data, err := ioutil.ReadFile(file_name)
	if nil != err {
		log.Printf("%v\n", 	errors.New(fmt.Sprintf("read data from file %s failed! %v\n",file_name, err )))
		return false
	}

	// 2. 反序列化
	if err := json.Unmarshal(data, &m); err != nil {
		log.Errorf("UnMarshal err : %v\n", err)
		return false
	}
	log.Printf("m : %v\n", m)
	return true
}