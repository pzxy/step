package main

import (
	"fmt"
	"time"
)

/**
常见的错误
1. 未初始化
2. 并发读写
*/

func main() {
	demo2()
}

/**
1. 未初始化
*/

func demo1() {
	//从一个nil的map中获取值不会panic，但是赋值操作会直接panic
	var m map[int]int
	fmt.Println(m[100])
}

//结构体中的map常常忘记初始化
func demo11() {
	var c Counter
	c.Website = "baidu.com"

	c.PageCounters["/"]++
}

type Counter struct {
	Website      string
	Start        time.Time
	PageCounters map[string]int
}

/**
2. 并发读写
此时必须加读写锁
*/

func demo2() {
	var m = make(map[int]int, 10) // 初始化一个map
	go func() {
		for {
			m[1] = 1 //设置key
		}
	}()

	go func() {
		for {
			_ = m[2] //访问这个map
		}
	}()
	select {}
}
