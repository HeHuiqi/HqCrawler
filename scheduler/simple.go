package scheduler

import "HqCrawler/engine"

/*
type SimpleScheduler struct {

	workerChan chan engine.Request
}

//hqconcurrent.ConcurrentEngine.Scheduler接口实现
func ( s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// send request down to worker chan

	//这里是为了防止 收发循环等待，重新起一个goroutine会立马返回,
	//每一个request都会起一个goroutine
	go func() {
		s.workerChan <- request
	}()
}
*/
type SimpleScheduler struct {

	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(w chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
/*
Submit()即是engine.ConcurrentEngine.Scheduler的接口又是
hqconcurrent.ConcurrentEngine.Scheduler接口
*/
//hqconcurrent.ConcurrentEngine.Scheduler接口实现
func (s *SimpleScheduler) Submit(request engine.Request) {
	// send request down to worker chan

	//这里是为了防止 收发循环等待，重新起一个goroutine会立马返回,
	//每一个request都会起一个goroutine
	go func() {
		s.workerChan <- request
	}()
}

//hqconcurrent.ConcurrentEngine.Scheduler接口实现
func ( s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}


