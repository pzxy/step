package scheduler

import "godemo/crawler/engine"

type SimplerScheduler struct {
	workerChan chan engine.Request
}

func (s *SimplerScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimplerScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimplerScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

/*func (s *SimplerScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}*/

func (s *SimplerScheduler) Submit(r engine.Request) {
	go func() { s.workerChan <- r }()
}
