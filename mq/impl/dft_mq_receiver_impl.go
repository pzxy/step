package impl

type DftReceiverImpl struct {
	queueName string

	PssChan chan interface{}
}

// 获取接收者需要监听的队列
func (r *DftReceiverImpl) QueueName1() string {
	return r.queueName
}

func (r *DftReceiverImpl) SetQueueName(queueName string) {
	r.queueName = queueName
}

func (r *DftReceiverImpl) SendQueueMsg2Pss(daemonMsg interface{}) {
	r.PssChan <- daemonMsg
}
