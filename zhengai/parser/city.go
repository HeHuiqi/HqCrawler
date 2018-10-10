package parser

import (
	"HqCrawler/engine"
	"regexp"
)
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
func ParserCity(contents []byte) engine.ParserResult  {
	re := regexp.MustCompile(cityRe)
	//match := re.FindAll(content,-1)
	match := re.FindAllSubmatch(contents,-1)

	result := engine.ParserResult{}
	for _,m := range match{

		name := string(m[2])
		result.Items = append(result.Items,"User:"+name)

		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParserProfile(c,name)
			},
		})
	}
	return result
}
