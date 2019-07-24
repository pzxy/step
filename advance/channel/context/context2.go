package main

import (
	"context"
	"fmt"
)

//同时读写
func writeData3(intChan chan int) {
	defer close(intChan)
	for i := 1; i <= 20; i++ {
		intChan <- i
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
	intChan := make(chan int)

	ctx, cancel2 := context.WithCancel(context.Background())
	ctx2, _ := context.WithCancel(ctx)
	ctx3 := context.WithValue(ctx2, "key1", "value1")
	//ctx4, _ := context.WithCancel(ctx3)

	go writeData3(intChan)
	go readData3(intChan, cancel2)

	<-ctx2.Done()
	fmt.Println(ctx3.Value("key1"))
}
