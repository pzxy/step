package mytime

import (
	"github.com/golang/glog"
	"step/utils/log"
	"sync"
	"time"
)

func SleepCalcProcess(taskId string, duration time.Duration, calcMutex *sync.Mutex) {
	calcMutex.Unlock()
	glog.V(log.Unc).Infof("task(%v), is now sleeping for %v with calcProcess unlocked...", taskId, duration)

	time.Sleep(duration)

	calcMutex.Lock()
	glog.V(log.Unc).Infof("task(%v), waked up and locked calcProcess after sleeping for %v", taskId, duration)
}

func SleepCalcProcessWithoutPrint(duration time.Duration, calcMutex *sync.Mutex) {
	calcMutex.Unlock()
	time.Sleep(duration)
	calcMutex.Lock()
}

func Sleep(runner string, duration time.Duration) {
	glog.V(log.Unc).Infof(runner+" is now sleeping for %v...", duration)

	time.Sleep(duration)

	glog.V(log.Unc).Infof(runner+" waked up after sleeping for %v", duration)
}
