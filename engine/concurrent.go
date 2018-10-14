package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
}
type Scheduler interface {

	ReadyNotifier
	Submit(request Request)
	WorkerChan()chan Request
	Run()
}
type ReadyNotifier interface {
	WorkerReady(w chan Request)
} 
func (e *ConcurrentEngine) Run(seeds ...Request)  {

	out := make(chan ParserResult)
	e.Scheduler.Run()
	for i := 0;i<e.WorkerCount ;i++  {
		createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)
	}

	//创建完worker后提交给Scheduler
	for _,r := range seeds {
		e.Scheduler.Submit(r)
	}

	//itemCount := 0
	for {
		result := <- out
		for _,item := range result.Items {

			//log.Printf("Got Item #%d: %v",itemCount,item)
			//itemCount++
			//存储
			//e.ItemChan <-item
			go func() {e.ItemChan <-item}()
		}
		for _,req := range result.Requests {
			if isDuplicate(req.Url) {
				continue
			}
			e.Scheduler.Submit(req)
		}
	}

}
var visitedUrls = make(map[string]bool)
func isDuplicate(url string) bool  {

	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
func createWorker(in chan Request,out chan ParserResult,ready ReadyNotifier)  {

	go func() {
		for {

			ready.WorkerReady(in)
			request := <- in
			result,err := Worker(request)
			if err != nil {
				continue
			}
			//等待返回结果
			out <- result
		}
	}()
}
