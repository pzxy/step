package mq

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"step/mq/itf"
	"step/mq/msgfrommq"
	"step/utils/log"
	"sync"
)

type RabbitMq struct {
	MqRcvUrl  string
	MqSendUrl string
	MqInfo    *MqInfo

	exchangeName string // exchange的名称
	exchangeType string // exchange的类型
	rcvConn      *amqp.Connection
	rcvChan      *amqp.Channel
	sendConn     *amqp.Connection
	sendChan     *amqp.Channel
	wg           sync.WaitGroup

	receivers []itf.Receiver
	pssChan   chan interface{}
	ifEnded   bool
	On_off    bool //控制发给pss.msgchan的开关
}

// NewMqClient 创建一个新的操作RabbitMq的对象
func NewMqClient(mqInfo *MqInfo, pssChan chan interface{}) (mq *RabbitMq) {
	mq = &RabbitMq{
		MqInfo:  mqInfo,
		pssChan: pssChan,
		On_off:  false,
	}
	mq.init()

	return
}

func (mq *RabbitMq) init() {
	mq.MqRcvUrl = "amqp://" + mq.MqInfo.UserName + ":" + mq.MqInfo.Password + "@" +
		mq.MqInfo.Ip + ":5672/" + mq.MqInfo.VHostRcv
	mq.MqSendUrl = "amqp://" + mq.MqInfo.UserName + ":" + mq.MqInfo.Password + "@" +
		mq.MqInfo.Ip + ":5672/" + mq.MqInfo.VHostSend

	mq.exchangeName = ""

	mq.RegisterReceiver(msgfrommq.NewTaskReqReceiver(mq.pssChan))
	mq.RegisterReceiver(msgfrommq.NewDtRobotsReqReceiver(mq.pssChan))
}

// doExec 启动Rabbitmq的客户端
func (mq *RabbitMq) Start() {
	defer mq.Close()
	for {
		if mq.ifNeedExit() {
			return
		}

		connectErr := mq.connect()
		if connectErr != nil {
			glog.Errorf("连接mq错误，错误信息：%v", connectErr)
			//mytime.Sleep("RabbitMq", reconnectDuration*time.Second)
			continue
		}

		mq.run()

		if mq.ifNeedExit() {
			return
		}

		// 一旦连接断开，那么需要隔一段时间去重连
		// 这里最好有一个时间间隔
		//mytime.Sleep("RabbitMq", reconnectDuration*time.Second)
	}
}

func (mq *RabbitMq) sendQueueMsg2Pss(daemonMsg interface{}) {
	mq.pssChan <- daemonMsg
}

// 发送msg给Mq
func (mq *RabbitMq) SendMsg2MqSrv(queueName string, msg interface{}) {
	if mq.sendConn.IsClosed() {
		sConn, err := amqp.Dial(mq.MqSendUrl)
		if err != nil {
			log.ErrLog(errors.WithStack(err))
			return
		}
		mq.sendConn = sConn
		sChan, err := mq.sendConn.Channel()
		if err != nil {
			log.ErrLog(errors.WithStack(err))
		}
		mq.sendChan = sChan
	}
	data, err := json.Marshal(msg)
	if err != nil {
		log.ErrLog(errors.WithStack(err))
		return
	}

	glog.V(log.Unc).Infof("====>msg to mq: %s\n", string(data))

	if mq.sendChan == nil {
		log.ErrLog(errors.WithStack(err))
		return
	}
	if err := mq.sendChan.Publish(mq.exchangeName, queueName, false,
		false, amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         data,
		}); err != nil {
		log.ErrLog(errors.WithStack(err))
	}
}

// run 开始获取连接并初始化相关操作
func (mq *RabbitMq) run() {
	mq.sendChan.Qos(1, 0, true) // 确保rabbitmq会一个一个发消息
	mq.wg.Add(1)
	//todo p600放开了   s300 注释了 by pzxy 2020年4月26日16:52:39
	//for _, receiver := range mq.receivers {
	//	mq.wg.Add(1)
	//	mygo.MyGo(mq.listen, receiver)
	//}

	glog.V(log.Nc).Infoln("Mq running...")
	mq.wg.Wait()

	glog.V(log.Nc).Infoln("---- All RabbitMq receivers have exit")
}

// RegisterReceiver 注册一个用于接收指定队列指定路由的数据接收者
func (mq *RabbitMq) RegisterReceiver(receiver itf.Receiver) {
	mq.receivers = append(mq.receivers, receiver)
}

// Listen 监听指定路由发来的消息
// 这里需要针对每一个接收者启动一个goroutine来执行listen
// 该方法负责从每一个接收者监听的队列中获取数据，并负责重试
func (mq *RabbitMq) listen(receiver itf.Receiver) {
	defer mq.wg.Done()

	queueName := receiver.QueueName1()

	msgS, err := mq.rcvChan.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.ErrLog(errors.WithStack(err))
		return
	}

	// 使用callback消费数据
	for msg := range msgS {
		// 暂时不考虑处理失败的情况
		// 当接收者消息处理失败的时候，
		// 比如网络问题导致的数据库连接失败，redis连接失败等等这种
		// 通过重试可以成功的操作，那么这个时候是需要重试的
		// 直到数据处理成功后再返回，然后才会回复rabbitmq ack
		if !mq.On_off && !receiver.MustReceive1() {
			glog.V(log.Nc).Infoln("RabbitMQ is off ...")
			continue
		}
		if !receiver.OnReceive1(msg.Body) {
			log.ErrLog(errors.New("receiver dealing with msg failed"))
		}
	}
}
func (mq *RabbitMq) connect() (err error) {
	//mq.rcvConn, err = amqp.Dial(mq.MqRcvUrl)
	//if err != nil {
	//	log.ErrLog(errors.WithStack(err))
	//	return
	//}
	//
	//mq.rcvChan, err = mq.rcvConn.Channel()
	//if err != nil {
	//	log.ErrLog(errors.WithStack(err))
	//	return
	//}

	//if mq.MqSendUrl != mq.MqRcvUrl {
	glog.V(log.Unc).Infof("rcv mq vHost(%v) is not equal to send mq VHost(%v)", mq.MqInfo.VHostRcv, mq.MqInfo.VHostSend)
	mq.sendConn, err = amqp.Dial(mq.MqSendUrl)
	if err != nil {
		log.ErrLog(errors.WithStack(err))
		return
	}

	mq.sendChan, err = mq.sendConn.Channel()
	if err != nil {
		log.ErrLog(errors.WithStack(err))
		return
	}
	//} else {
	//mq.sendConn = mq.rcvConn
	//mq.sendChan = mq.rcvChan

	//mq.wg = &sync.WaitGroup{}
	glog.V(log.Nc).Infoln("mqConnect(), Success")
	return nil
}

func (mq *RabbitMq) ifNeedExit() bool {
	return mq.ifEnded
}
func (mq *RabbitMq) prepare2Exit() {
	mq.ifEnded = true
}

func (mq *RabbitMq) Close() {
	mq.prepare2Exit()
	//if mq.rcvChan != nil {
	//	mq.rcvChan.Close()
	//	mq.rcvChan = nil
	//}
	//
	//if mq.rcvConn != nil {
	//	mq.rcvConn.Close()
	//	mq.rcvConn = nil
	//}

	if mq.MqInfo.VHostRcv != mq.MqInfo.VHostSend {
		if mq.sendChan != nil {
			mq.sendChan.Close()
			mq.sendChan = nil
		}

		if mq.sendConn != nil {
			mq.sendConn.Close()
			mq.sendConn = nil
		}
	}

}

func (mq *RabbitMq) ControlOn_off(Opt bool) {
	mq.On_off = Opt
}
