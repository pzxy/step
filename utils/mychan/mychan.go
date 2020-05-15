package mychan

import (
	"github.com/golang/glog"
	"step/utils/log"
)

func Send(key string, dstChan chan interface{}, data interface{}) {
	if dstChan == nil {
		glog.V(log.Unc).Infof("%v, exit for dstChan is nil, discard(%#v)", key, data)
		return
	}

	select {
	case dstChan <- data:
	default:
		glog.V(log.Unc).Infof("%v, exit for dstChan is full, discard(%#v)", key, data)
	}
}

func SendBytes(key string, dstChan chan []byte, data []byte) {
	if dstChan == nil {
		glog.V(log.Unc).Infof("%v, exit for dstChan is nil, discard(%#v)", key, data)
		return
	}

	select {
	case dstChan <- data:
	default:
		glog.V(log.Unc).Infof("%v, exit for dstChan is full, discard(%#v)", key, data)
	}
}
