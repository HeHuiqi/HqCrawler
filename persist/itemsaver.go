package persist

import (
	"log"
	"github.com/olivere/elastic"
	"context"
	"reflect"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver Got Item #%d: %v",itemCount,item)
			itemCount++
			_,err := save(item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v",item,err)
			}
		}
	}()
	return out
}
func save(item interface{}) (id string,err error)  {
	if reflect.TypeOf(item).Name() == "string" {
		return "",nil
	}
	client,err := elastic.NewClient(
		//Must turn off sniff in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		return "",err
	}
	resp,err := client.Index().
		Index("dating_profile").
		Type("zhenai").
		BodyJson(item).Do(context.Background())

	if err != nil {
		return "",err
	}
	return resp.Id,nil
}