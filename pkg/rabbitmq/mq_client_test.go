package rabbitmq

import (
	"step/middleware/mq/mqparams"
	"step/utils/mygo"
	"testing"
	"time"
)

/**
  "Ip": "120.79.177.199",
   "Port": 5762,
   "UserName": "dispatch",
   "Password": "U674u1Qk&Xf!tD",
   "VHostRcv": "/pss_test_aa",
   "VHostSend": "/pss_test_aa"
*/
func TestMq(t *testing.T) {
	var c chan interface{}
	mqInfo := &MqInfo{
		Ip:        "120.79.177.199",
		Port:      5762,
		UserName:  "dispatch",
		Password:  "U674u1Qk&Xf!tD",
		VHostRcv:  "/pss_test_aa",
		VHostSend: "/pss_test_aa",
	}
	//mConfig := params.LoadManualConfigs()
	mq := NewMqClient(mqInfo, c)
	mygo.MyGo(mq.Start)
	time.Sleep(time.Second * 3)
	mq.SendMsg2MqSrv(mqparams.NtfToBackstageDspDown, "这是要传递的数据，interface{}，传什么都可以")
	time.Sleep(time.Second * 10)
}
