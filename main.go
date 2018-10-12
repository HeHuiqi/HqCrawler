package main

import (
	"HqCrawler/engine"
	"HqCrawler/scheduler"
	"HqCrawler/zhengai/parser"
	"HqCrawler/persist"
)

func main() {

	//1
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun",
	//	ParserFunc:parser.ParserCityList,
	//})



	//2
	//e := hqconcurrent.ConcurrentEngine{
	//	Scheduler:&scheduler.SimpleScheduler{},
	//	WorkerCount:100,
	//}
	//e.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun",
	//	ParserFunc:parser.ParserCityList,
	//})



	//3
	//e := engine.ConcurrentEngine{
	//	Scheduler:&scheduler.QueuedScheduler{},
	//	WorkerCount:100,
	//}
	//e.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun",
	//	ParserFunc:parser.ParserCityList,
	//})



	//4重构Scheduler接口做成通用接口
	//e := engine.ConcurrentEngine{
	//	Scheduler:&scheduler.SimpleScheduler{},
	//	WorkerCount:100,
	//}
	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:100,
		ItemChan:persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParserCityList,
	})

}

