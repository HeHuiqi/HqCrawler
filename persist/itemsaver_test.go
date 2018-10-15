package persist

import (
	"testing"
	"HqCrawler/model"
	"github.com/olivere/elastic"
	"context"
	"encoding/json"
	"HqCrawler/engine"
)

func TestItemSaver(t *testing.T) {

	expected := engine.Item{
		Url:"http://album.zhenai.com/u/108906739",
		Type:"zhenai",
		Id:"108906739",
		Payload:model.Profile{
			Age:34,
			Height:162,
			Weight:162,
			Income:"3001-5000元",
			Gender:"女",
			Name:"安静的雪",
			Xinzuo:"牡羊座",
			Occupation:"人事/行政",
			Marriage:"离异",
			House:"已购房",
			Hukou:"山东菏泽",
			Education:"大学本科",
			Car:"未购车",
		},

	}
	// TODO: try to start up elastic search
	// here using docker go client.
	client,err := elastic.NewClient(
		//Must turn off sniff in docker
		elastic.SetSniff(false),
	)

	//save item
	const index  = "dating_test"
	err = save(client,index,expected)
	if err !=nil {
		panic(err)
	}

	//get item
	resp,err := client.Get().Index(index).
		Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s",resp.Source)

	var actual engine.Item
	err = json.Unmarshal([]byte(*resp.Source),&actual)

	actualProfile,_ := model.FromJsonObj(actual.Payload)
	actual.Payload=  actualProfile

	//verify item
	if actual != expected {
		t.Errorf("got %v; expected %v",actual,expected)
	}

}
