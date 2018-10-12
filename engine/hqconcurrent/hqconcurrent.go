package hqconcurrent

import (
	"log"
	"HqCrawler/engine"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}
type Scheduler interface {

	Submit(request engine.Request)
	ConfigureMasterWorkerChan (chan engine.Request)
}
func (e *ConcurrentEngine) Run(seeds ...engine.Request)  {

	in := make(chan engine.Request)
	out := make(chan engine.ParserResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)
	for i := 0;i<e.WorkerCount ;i++  {
		createWorker(in,out)
	}

	//创建完worker后提交给Scheduler
	for _,r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <- out
		for _,item := range result.Items {
			log.Printf("Got item #%d: %v",itemCount,item)
			itemCount++
		}
		for _,req := range result.Requests {
			e.Scheduler.Submit(req)
		}
	}

}
func createWorker(in chan engine.Request,out chan engine.ParserResult)  {
	go func() {
		for {
			//这里会造成收发循环等待。处理方式是在提交request时候重新起一个goroutine
			request := <- in
			result,err := engine.Worker(request)
			if err != nil {
				continue
			}
			//等待返回结果
			out <- result
		}
	}()
}
