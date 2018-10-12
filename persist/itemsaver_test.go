package persist

import (
	"testing"
	"HqCrawler/model"
	"github.com/olivere/elastic"
	"context"
	"encoding/json"
)

func TestItemSaver(t *testing.T) {

	profile := model.Profile{
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
	}
	id,err := save(profile)
	if err !=nil {
		panic(err)
	}
	// TODO: try to start up elastic search
	// here using docker go client.
	client,err := elastic.NewClient(
		//Must turn off sniff in docker
		elastic.SetSniff(false),
	)
	resp,err := client.Get().Index("dating_profile").
		Type("zhenai").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s",resp.Source)

	var actual model.Profile
	err = json.Unmarshal([]byte(*resp.Source),&actual)
	if err != nil {
		panic(err)
	}
	if actual != profile {
		t.Errorf("got %v; expected %v",actual,profile)
	}

}
