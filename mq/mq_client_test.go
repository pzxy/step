package mq

import "testing"

func TestMq(t *testing.T) {
	var c chan interface{}
	NewMqClient(&MqInfo{"guest", "guest"}, c)
}
