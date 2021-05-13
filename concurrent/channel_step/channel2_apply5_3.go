package main

import (
	"fmt"
	"time"
)

/**
扇出模式
*/

func fanOut(ch <-chan interface{}, out []chan interface{}, async bool) {
	go func() {
		defer func() { //退出时关闭所有的输出chan
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		for v := range ch { // 从输入chan中读取数据
			fmt.Println("1231324")
			v := v
			for i := 0; i < len(out); i++ {
				i := i
				if async { //异步
					go func() {
						out[i] <- v // 放入到输出chan中,异步方式
					}()
				} else {
					out[i] <- v // 放入到输出chan中，同步方式
				}
			}
		}
	}()
}

func main() {
	in := make(chan interface{})
	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
	}()

	out := make([]chan interface{}, 2)
	out[0] = make(chan interface{})
	out[1] = make(chan interface{})
	go func() {
		for {
			select {
			case a := <-out[0]:
				fmt.Println(a)
			}
		}
	}()

	go func() {
		for {
			select {
			case a := <-out[1]:
				fmt.Println(a)
			}
		}
	}()

	fanOut(in, out, true)
	time.Sleep(time.Second * 30)
}
