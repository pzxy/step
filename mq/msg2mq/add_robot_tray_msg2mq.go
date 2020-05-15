package msg2mq

import "step/mq/msgfrommq"

type FindExitResult int8

const (
	Found4NewRobotOrTray    FindExitResult = 1
	NotFound4NewRobotOrTray FindExitResult = -1
	AddNewRobotOrTrayTask   FindExitResult = 2
	AddRobotOrTraySuccess   FindExitResult = 3
)

// 转化为json时，字段的名字需要改变(开头均改为小写）
type Ack4AddRobotOrTrayAsk struct {
	UniqueIdentification string                       `json:"taskId"`
	Code                 msgfrommq.AddRobotOrTrayType `json:"code"`
	Status               FindExitResult               `json:"status"`
	MapDetailsId         string                       `json:"mapDetailsId"`
}

func NewAck4AddRobotOrTrayAsk(taskId string, code msgfrommq.AddRobotOrTrayType, result FindExitResult, mapDetailId string) (ack *Ack4AddRobotOrTrayAsk) {
	ack = &Ack4AddRobotOrTrayAsk{
		UniqueIdentification: taskId,
		Code:                 code,
		Status:               result,
		MapDetailsId:         mapDetailId,
	}

	return
}
