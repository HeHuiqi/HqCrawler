package persist

import (
	"log"
	"github.com/olivere/elastic"
	"context"
	"reflect"
	"HqCrawler/engine"
	"github.com/pkg/errors"
	"encoding/json"
)

func ItemSaver(index string) (chan engine.Item,error){
	client,err := elastic.NewClient(
		//Must turn off sniff in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil,err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver Got Item #%d: %v",itemCount,item)
			itemCount++
			err := save(client,index,item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v",item,err)
			}
		}
	}()
	return out,nil
}
func save(client *elastic.Client,index string,item engine.Item) error  {

	//这里不存储城市列表名字和某个城市下用户名
	if reflect.TypeOf(item.Payload).Name() == "string" {
		return nil
	}

	if item.Type == "" {
		return errors.New("Must supply Type")
	}

	itemJsonStr,_ := json.Marshal(item)
	//fmt.Println("Item",string(itemJsonStr))

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyString(string(itemJsonStr))

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_,err := indexService.Do(context.Background())

	if err != nil {
		return err
	}
	return nil
}