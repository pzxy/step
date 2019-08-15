package main

import "fmt"

//同时读写 写完以后关闭管道
func writeData(intChan chan int) {
	defer close(intChan)
	for i := 1; i <= 5000; i++ {
		intChan <- i
	}
}
func readData(intChan chan int, boolChan chan bool) {
	defer close(boolChan)
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读出来的数据: %v\n", v)
	}
	boolChan <- true
}

func main() {
	intChan := make(chan int)
	boolChan := make(chan bool)

	go writeData(intChan)
	go readData(intChan, boolChan)

	//for {
	//	if _, ok := <-boolChan; !ok {
	//		break
	//	}
	//}
	<-boolChan

}
