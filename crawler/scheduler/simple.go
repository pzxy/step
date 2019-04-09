package scheduler

import "godemo/crawler/engine"

type SimplerScheduler struct {
	workerChan chan engine.Request
}

//下面实现了Schduler接口
func (s *SimplerScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimplerScheduler) Submit(r engine.Request) {
	go func() { s.workerChan <- r }()
}
