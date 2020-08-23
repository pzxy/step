package tmp

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
	"step/utils/mygo"
	"step/utils/mytime"
	"testing"
	"time"
)

var gT *testing.T

var conn *amqp.Connection
var channel *amqp.Channel
var count = 0

const (
	queueName = "backStageToDispatchingTaskQueue"
	//exchange  = "park_ex"
	exchange = ""
	mqurl    = "amqp://guest:guest@127.0.0.1:5672/"
)

func TestRabbitMq(t *testing.T) {
	gT = t

	mygo.MyGo(myGoFunc4TestRabbitMq)

	var daemon chan struct{}
	<-daemon
	receive()
	fmt.Println("end")
	closeRabbitMq()
}

func myGoFunc4TestRabbitMq() {
	for {
		push()
		mytime.Sleep("Tester ", 1*time.Second)
	}
}

func failOnErr(err error, msg string) {
	if err != nil {
		gT.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

func mqConnect() {
	var err error
	conn, err = amqp.Dial(mqurl)
	failOnErr(err, "failed to connect tp rabbitmq")

	channel, err = conn.Channel()
	failOnErr(err, "failed to open a channel")

	fmt.Println("mqConnect() success.")
}

func closeRabbitMq() {
	channel.Close()
	conn.Close()
}

//连接rabbitmq server
func push() {

	if channel == nil {
		mqConnect()
	}
	fmt.Println("pushing...")
	msgContent := "hello world!"

	channel.Publish(exchange, queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})
	fmt.Println("publish success")
}

func receive() {
	if channel == nil {
		mqConnect()
	}

	msgs, err := channel.Consume(queueName, "", true, false, false, false, nil)
	failOnErr(err, "")

	forever := make(chan bool)

	mygo.MyGo(myGoFunc4receive, msgs)

	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C\n")
	<-forever
}

func myGoFunc4receive(msgs chan amqp.Delivery) {
	for d := range msgs {
		s := BytesToString(&(d.Body))
		count++
		fmt.Printf("receve msg is :%s -- %d\n", *s, count)
	}
}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}
