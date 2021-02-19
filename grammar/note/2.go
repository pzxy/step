package main

import (
	"fmt"
	"time"
)

//就是说这个v不在本goroutinue空间，所以他取值的时候要去主mian.go里面取这个位置的值，所以取值的时候，当前for到的v是哪个他取出来的就是哪个。
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range s {
		time.Sleep(time.Second)
		go func() {
			fmt.Println(v)
		}()
	}

}
