package parser

import (
	"testing"
	"io/ioutil"
	"net/http"
	"fmt"
)
func hqCatchData(url string)  {
	req,_:=http.NewRequest(http.MethodGet,url,nil)
	req.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	//resp,err :=http.DefaultClient.Do(req)
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {

			fmt.Println("Redict:",req.URL)
			return nil
		},
	}
	resp,_ :=client.Do(req)
	defer resp.Body.Close()
	contents,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(contents))
}

func TestData(t *testing.T)  {
	hqCatchData("http://album.zhenai.com/u/1995815593")
}
func TestParserCityList(t *testing.T) {
	/*
	contents,err := hqfetcher.HqFetch("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	*/
	contents,err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParserCityList(contents,"")
	//fmt.Printf("contents == %s",contents)

	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}


	if len(result.Requests) != resultSize  {
		t.Errorf("result should have %d requests;but had %d",resultSize,len(result.Requests))
	}

	for i,url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s;but was %s",i,url,result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize  {
		t.Errorf("result should have %d requests;but had %d",resultSize,len(result.Requests))
	}

}