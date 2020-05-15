package itf

// Receiver 观察者模式需要的接口
// 观察者用于接收指定的queue到来的数据
type Receiver interface {
	QueueName1() string     // 获取接收者需要监听的队列
	OnError1(error)         // 处理遇到的错误，当RabbitMQ对象发生了错误，他需要告诉接收者处理错误
	OnReceive1([]byte) bool // 处理收到的消息, 这里需要告知RabbitMQ对象消息是否处理成功
	MustReceive1() bool     // 初始化能接受消息
}
