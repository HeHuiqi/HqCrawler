package parser

import (
	"HqCrawler/engine"
	"regexp"
)
var  (
	profile1Re = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/chengdu/[^"]+)"`)
)
func ParserCity(contents []byte) engine.ParserResult  {

	matchs := profile1Re.FindAllSubmatch(contents,-1)

	result := engine.ParserResult{}
	for _,m := range matchs{

		name := string(m[2])
		result.Items = append(result.Items,name)

		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParserProfile(c,name)
			},
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
