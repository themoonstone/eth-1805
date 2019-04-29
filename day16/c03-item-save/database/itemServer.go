package database

import (
	"1805/day16/c03-item-save/engine"
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/olivere/elastic"
)

// 有效数据的处理
func ItemServer() (chan interface{}){
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		// Handle error
		return nil
	}
	out := make(chan interface{})
	go func() {
		for {
			item := <- out
			fmt.Printf("Item Server : get item:%v\n", item)
			items := engine.Item{
				Data: item,
				Type: "students",
				Index: "eth1805",
			}
			// 存入数据
			err := Save(client, items)
			if err != nil {
				log.Printf("Item Save: error %v\n", err)
			}
		}
	}()
	return out
}

func Save(client *elastic.Client, item engine.Item) error {
	// http:localhost:9200/index/type/id--(json)
	indexServ := client.Index().Index(item.Index).Type(item.Type).BodyJson(item)
	// TODO context包 ??
	_, err := indexServ.Do(context.Background())
	return err
}
//func ElasticSearch() error {
//	fmt.Println("get info from es")
//	resp, err := http.Get("http://localhost:9200/eth1805/students/1")
//	if nil != err {
//		return err
//	}
//	fmt.Printf("%v\n", resp.Body)
//	return nil
//}
//
//func GoElasticSearch() error {
//	ctx := context.Background()
//	client, err := elastic.NewClient(
//		elastic.SetSniff(false))
//	if err != nil {
//		// Handle error
//		panic(err)
//	}
//	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
//	if err != nil {
//		// Handle error
//		panic(err)
//	}
//	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
//	return nil
//}