package main

import (
	"fmt"
	"sync"
)

func main() {
	var one sync.Once
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			one.Do(onetest)
			fmt.Println(i)
			wg.Done()
		}(i)

	}
	wg.Wait()
}

func onetest() {
	fmt.Println("执行一次")
}
