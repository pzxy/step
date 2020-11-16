package main

import (
	"fmt"
	"sync"
)

/**
死锁1
*/

func main() {
	var once sync.Once
	var f = func() {
		once.Do(func() {
			fmt.Println("初始化")
		})
	}
	//这里会发生死锁是因为，Do方法中有锁，执行完以后才会解锁。再未解锁时执行到了f方法中的Do就会再次上锁，
	//因为go中的互斥锁是不可重入的锁，所以会报错。
	once.Do(f)
}
