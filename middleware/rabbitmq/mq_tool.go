package rabbitmq

import (
	"sync"
)

var (
	//是否完成充电
	chargeCompleteRobotIds []string

	//TaskID Add
	ChargeTaskID_Add int

	ChargeTaskID_Add_Lock sync.RWMutex
)

const reconnectDuration = 3
const (
	ChargeStop     int32 = 2
	MacChargeCount int   = 100
)

type MqInfo struct {
	Ip        string `json:"Ip"`
	Port      int32  `json:"Port"`
	UserName  string `json:"UserName"`
	Password  string `json:"Password"`
	VHostRcv  string `json:"VHostRcv"`
	VHostSend string `json:"VHostSend"`
}
