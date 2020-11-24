package main

import (
	"fmt"
	"time"
)

func main() {
	s1 := make([]int, 5, 10)
	s2 := s1[:5]
	s3 := s1[:5:5]
	fmt.Printf("s1 :len:%v,cap:%v,%p,%v \n", len(s1), cap(s1), s1, s1)
	fmt.Printf("s2 :len:%v,cap:%v,%p,%v \n", len(s2), cap(s2), s2, s2)
	fmt.Printf("s3 :len:%v,cap:%v,%p,%v \n", len(s3), len(s3), s3, s3)
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Printf("1")
			//break
		}
	}()
	select {} //只要有其他协程没有退出，这里就不会有问题，否则会发生死锁。
}
