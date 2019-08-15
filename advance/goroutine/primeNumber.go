package main

import (
	"fmt"
	"time"
)

//求素数
//总结:创建channel后向里面发的时候必须考虑到有人在接受这个channel
//这个时候不分我是先发还是 先接受呢,如果都写了话,系统会知道怎么去调度的.
//并发的取发送和接受,都在一个协程中肯定会造成死锁

func numChanFunc(numChan chan int, primeChan chan int, boolChan chan bool) {
	var flag bool
	for {
		num, ok := <-numChan
		if !ok {
			break
		}
		if num > 2 {
			flag = true
			for i := 2; i < num; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				primeChan <- num
			}
		}
	}
	boolChan <- true
}

func main() {
	var num int
	numChan := make(chan int)
	primeChan := make(chan int)
	boolChan := make(chan bool)
	t := time.Now()

	go func() {
		for i := 0; i < 100000; i++ {
			num = i
			numChan <- num
		}
		defer close(numChan)
	}()
	//开四个携程来做
	for i := 0; i < 4; i++ {
		go numChanFunc(numChan, primeChan, boolChan)
	}
	//这个协程来接受值
	go func() {
		defer close(primeChan)
		for {
			val, ok := <-primeChan
			if !ok {
				break
			}
			fmt.Printf("素数是: %v\n", val)
		}
	}()
	//通知四个携程
	for i := 1; i <= 4; i++ {
		<-boolChan
	}
	fmt.Println(time.Now().Sub(t))

}
