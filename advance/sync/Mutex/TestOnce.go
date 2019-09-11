package main

import (
	"step/advance/sync/once"
	"time"
)

func main() {
	for i:=0;i<10;i++{
		once.Test3()
	}
	time.Sleep(time.Second*3)
}
