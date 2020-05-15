package msgfrommq

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"step/mq/impl"
	"step/mq/mqparams"
	"step/utils/log"
)

//Dt (data transfer 数据传输)
type DtStartOrStop int32

const (
	DtStart DtStartOrStop = 1
	DtStop  DtStartOrStop = 2
)

type DtRobotsReqReceiver struct {
	impl.DftReceiverImpl
	DtRobotsReq *DtRobotsReq
}

type DtRobotsReq struct {
	//接受来自后台的数据
	OldDownloadUrl     string `json:"olddownloadUrl"`
	DownloadUrl        string `json:"downloadUrl"`
	Content            string `json:"content"`
	Version            uint16 `json:"version"`
	MD5                string `json:"md5"`
	UpdateSerialNumber uint32 `json:"updateSerialNumber"`
	DeviceIds          []int  `json:"deviceIds"`
	PointCode          string `json:"pointCode"`
}

//创建一个接受者从后台接受除命令信息外的数据,比如要升级的机器人等信息
//mq客户端中注册
func NewDtRobotsReqReceiver(pssChan chan interface{}) (receiver *DtRobotsReqReceiver) {
	receiver = &DtRobotsReqReceiver{}
	//队列管道
	receiver.SetQueueName(mqparams.ReqDtRobotsMq)
	receiver.PssChan = pssChan

	return
}

// 固定写法  处理遇到的错误，当RabbitMQ对象发生了错误，他需要告诉接收者处理错误
func (r *DtRobotsReqReceiver) OnError1(err error) {
	log.ErrLog(errors.WithStack(err))
}

// 固定写法 处理收到的消息, 这里需要告知RabbitMQ对象消息是否处理成功
//这里将数据发送过去
func (r *DtRobotsReqReceiver) OnReceive1(msg []byte) bool {
	glog.V(log.Unc).Infoln("DtRobotsReqReceiver received a msg : ", string(msg))

	r.DtRobotsReq = &DtRobotsReq{}

	/*testJson := make([]ChargeRobotOrder, 0)*/

	if err := json.Unmarshal(msg, r.DtRobotsReq); err != nil {
		log.ErrLog(errors.Errorf("msg = %+v", r.DtRobotsReq))
		return false
	}

	/*r.chargeRobotsReq = testjson*/
	glog.V(log.Unc).Infof("receiver info from %+v", r.DtRobotsReq)
	//dtRobotOrder := &DtRobotOrder{
	//	UpdateSerialNumber: r.DtRobotsReq.UpdateSerialNumber,
	//	DeviceIds:          r.DtRobotsReq.DeviceIds,
	//}

	//发送消息给调度
	//r.SendQueueMsg2Pss(dtRobotOrder)
	return true
}

func (r *DtRobotsReqReceiver) MustReceive1() bool {
	return false
}
