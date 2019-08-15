package main

import (
	"context"
	"fmt"
	"time"
)

//同时读写
func writeData3(intChan chan int) {
	defer close(intChan)
	for i := 1; i <= 20; i++ {
		intChan <- i
		time.Sleep(300 * time.Millisecond)
	}
}
func readData3(intChan chan int, cancel func()) {
	defer cancel()
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("intChan 读出来的数据: %v\n", v)

	}

}

func main() {
	//context是个上下文树
	intChan := make(chan int)

	ctx, cancel2 := context.WithCancel(context.Background())
	ctx2, _ := context.WithCancel(ctx)
	ctx3, _ := context.WithDeadline(ctx2, time.Now().Add(1*time.Second))
	ctx4 := context.WithValue(ctx3, "key1", "value1")
	//ctx4, _ := context.WithCancel(ctx3)
	fmt.Print(ctx4)
	go writeData3(intChan)
	go readData3(intChan, cancel2)
	<-ctx4.Done()
	fmt.Println(ctx4.Value("key1"))
}
