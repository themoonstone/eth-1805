package database

import "github.com/labstack/gommon/log"

// 有效数据的处理
func ItemServer(in chan interface{}) chan interface{}{
	out := make(chan interface{})
	go func() {
		for {
			item := <- in
			log.Printf("Item Server : get item:%v\n", item)
		}
	}()
	return out
}