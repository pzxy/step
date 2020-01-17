package main

import (
	"fmt"
	"time"
)

func main() {
	t := new(ttt)
	c1 := t.GetChan()
	go func() {
		for s := range c1 {
			fmt.Println(s)
		}
	}()
	c2 := make(chan interface{})
	c1 = c2
	c2 <- 2
	time.Sleep(10 * time.Second)
}

type ttt struct {
	c chan interface{}
}

func (t *ttt) GetChan() chan interface{} {

	return t.c
}
