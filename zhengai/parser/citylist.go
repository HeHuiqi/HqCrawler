package parser

import (
	"HqCrawler/engine"
	"regexp"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`


func ParserCityList(contents []byte,_ string) engine.ParserResult  {

	re := regexp.MustCompile(cityListRe)
	//match := re.FindAll(content,-1)
	match := re.FindAllSubmatch(contents,-1)

	result := engine.ParserResult{}



	for _,m := range match{

		url := string(m[1])
		name := string(m[2])

		cpinyin := url[len("http://www.zhenai.com/zhenghun/"):]
		item := engine.Item{
			Url: url,
			Type:"zhenai",
			Id:cpinyin,
			Payload:name,

		}
		result.Items = append(result.Items,item)

		result.Requests = append(result.Requests,engine.Request{
			Url:url,
			ParserFunc:ParserCity,
		})
		//fmt.Printf("City: %s, URL: %s",m[2],m[1])
		//fmt.Println()

	}
	return result
}

