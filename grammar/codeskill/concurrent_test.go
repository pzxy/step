package codeskill

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestConcurrentFromChan(t *testing.T) {
	// test11()
	f1 := func(msg *SendMsg) {
		fmt.Println("f1, msg:", *msg)
	}
	in := make(chan *SendMsg, 0)
	ConcurrentFromChan(in, 10, f1)

	for i := 0; i < 100; i++ {
		msg := SendMsg(strconv.Itoa(i))
		in <- &msg
	}

	<-time.After(3 * time.Second)
}
