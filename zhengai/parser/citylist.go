package parser

import (
	"HqCrawler/engine"
	"regexp"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParserResult  {

	re := regexp.MustCompile(cityListRe)
	//match := re.FindAll(content,-1)
	match := re.FindAllSubmatch(contents,-1)

	result := engine.ParserResult{}

	for _,m := range match{

		result.Items = append(result.Items,string(m[2]))

		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:ParserCity,
		})
		//fmt.Printf("City: %s, URL: %s",m[2],m[1])
		//fmt.Println()

	}
	return result
}

