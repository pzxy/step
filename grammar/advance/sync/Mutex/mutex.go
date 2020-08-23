package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//互斥锁(全局锁)不确定读写时

func main() {
	lock := new(sync.Mutex)

	testMap(lock)
}
func testMap(lock *sync.Mutex) {
	var a map[int]int
	a = make(map[int]int, 5)

	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[8] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}

	lock.Lock()
	fmt.Println(a)
	lock.Unlock()

	time.Sleep(time.Second)
	fmt.Println(a)
}
