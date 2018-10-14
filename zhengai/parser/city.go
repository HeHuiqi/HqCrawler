package parser

import (
	"HqCrawler/engine"
	"regexp"
)
var  (
	profile1Re = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/chengdu/[^"]+)"`)

)
func ParserCity(contents []byte,_ string) engine.ParserResult  {

	matchs := profile1Re.FindAllSubmatch(contents,-1)

	result := engine.ParserResult{}
	for _,m := range matchs{

		url := string(m[1])
		name := string(m[2])
		item := engine.Item{
			Url: url,
			Type:"zhenai",
			Id:extractString([]byte(url),idUrlRe),
			Payload:name,

		}
		result.Items = append(result.Items,item)

		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc: ProfileParser(name),
		})
	}

	matchs = cityUrlRe.FindAllSubmatch(contents,-1)
	for _,m := range matchs {
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc: ParserCity,
		})
	}



	return result
}
