package msgfrommq

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"step/middleware/mq/impl"
	"step/middleware/mq/mqparams"
	"step/utils/log"
)

type TaskReqReceiver struct {
	impl.DftReceiverImpl
	taskReq *TaskReq
}

type TaskReq struct {
	PointCode                      string
	TaskId                         string
	InParkId                       string `json:"inParkId"`
	MapId                          string
	ElevatorId                     string
	ExecutingTaskId                string
	EntranceOrExitOrientationAngle uint16
	CarHeadAngle                   uint16
	EntranceOrExitNodeId           string
}

func NewTaskReqReceiver(pssChan chan interface{}) (receiver *TaskReqReceiver) {
	receiver = &TaskReqReceiver{}
	receiver.SetQueueName(mqparams.ReqTaskReqMq)
	receiver.PssChan = pssChan

	return
}

// 处理遇到的错误，当RabbitMQ对象发生了错误，他需要告诉接收者处理错误
func (r *TaskReqReceiver) OnError1(err error) {
	log.ErrLog(errors.WithStack(err))
}

// 处理收到的消息, 这里需要告知RabbitMQ对象消息是否处理成功
func (r *TaskReqReceiver) OnReceive1(msg []byte) bool {
	glog.V(log.Unc).Infoln("TaskReqReceiver received a msg : ", string(msg))

	r.taskReq = &TaskReq{}
	if err := json.Unmarshal(msg, r.taskReq); err != nil {
		log.ErrLog(errors.Errorf("msg = %v", msg))
		return false
	}

	r.SendQueueMsg2Pss(r.taskReq)
	return true
}

func (r *TaskReqReceiver) MustReceive1() bool {
	return false
}
