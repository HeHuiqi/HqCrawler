package main

import (
	"HqCrawler/engine"
	"HqCrawler/zhengai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParserCityList,
	})


}

