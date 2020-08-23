package msgfrommq

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"step/middleware/mq/impl"
	"step/middleware/mq/mqparams"
	"step/utils/log"
)

type AddRobotOrTrayType int

// fixme
const (
	AddTrayType      AddRobotOrTrayType = -1
	AddRobotType     AddRobotOrTrayType = 1
	AddRobotTrayType AddRobotOrTrayType = 2
	DelRobotType     AddRobotOrTrayType = 3
)

type Ask4AddRobotOrTrayReqReceiver struct {
	impl.DftReceiverImpl
	ask4AddRobotOrTrayReq *Ask4AddRobotOrTrayReq
}

type Ask4AddRobotOrTrayReq struct {
	TaskId   string             `json:"taskId"`
	Code     AddRobotOrTrayType `json:"code"`
	DeviceId int32              `json:"deviceId"`
}

func NewAsk4AddRobotOrTrayReqReceiver(pssChan chan interface{}) (receiver *Ask4AddRobotOrTrayReqReceiver) {
	receiver = &Ask4AddRobotOrTrayReqReceiver{}
	receiver.SetQueueName(mqparams.ReqAsk4AddingRobotOrTrayMq)
	receiver.PssChan = pssChan

	return
}

// 处理遇到的错误，当RabbitMQ对象发生了错误，他需要告诉接收者处理错误
func (r *Ask4AddRobotOrTrayReqReceiver) OnError(err error) {
	log.ErrLog(errors.WithStack(err))
}

// 处理收到的消息, 这里需要告知RabbitMQ对象消息是否处理成功
func (r *Ask4AddRobotOrTrayReqReceiver) OnReceive(msg []byte) bool {
	glog.V(log.Unc).Infoln("Ask4AddRobotOrTrayReqReceiver received a msg : ", string(msg))

	r.ask4AddRobotOrTrayReq = &Ask4AddRobotOrTrayReq{}
	if err := json.Unmarshal(msg, r.ask4AddRobotOrTrayReq); err != nil {
		log.ErrLog(errors.Errorf("msg = %v", msg))
		return false
	}

	r.SendQueueMsg2Pss(r.ask4AddRobotOrTrayReq)
	return true
}

func (r *Ask4AddRobotOrTrayReqReceiver) MustReceive() bool {
	return false
}
