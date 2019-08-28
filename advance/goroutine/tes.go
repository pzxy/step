package main

import (
	"fmt"
	"time"
)

func main() {

}
func test02() {

}

func test01() {
	//var num int
	intchan := make(chan interface{})
	go func() {
		for i := 1; i < 10; i++ {
			v, ok := <-intchan
			if !ok {
				break
			}
			fmt.Printf("输出值为 : %v\n", v)
		}
	}()

	for i := 1; i < 10; i++ {
		sli := []int{1, 2, 3, 4}
		intchan <- sli
	}
	close(intchan)
	time.Sleep(time.Second)
}
