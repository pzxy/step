package main

import (
	"fmt"
	"sync"
)

//读写锁
func main() {
	var rw sync.RWMutex
	rw.RLocker()
	var wg sync.WaitGroup
	wg.Add(2)

	go read(&rw, 1, &wg)
	go read(&rw, 2, &wg)
	wg.Wait()
}

//wg *sync.WaitGroup必须是地址传递,保证是同一个对象
func read(rw *sync.RWMutex, i int, wg *sync.WaitGroup) {
	fmt.Println(i, "开始读...")
	rw.RLock()
	fmt.Println(i, "正在读...")
	rw.RUnlock()
	fmt.Println(i, "结束读...")
	wg.Done()
}
