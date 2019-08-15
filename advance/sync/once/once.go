package main

import (
	"fmt"
	"sync"
)

func main() {
	var one sync.Once
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(2)

		go func(i int) {
			one.Do(onetest)

			wg.Done()
		}(i)
		go func(i int) {
			one.Do(onetest)
			wg.Done()
		}(i)
		fmt.Println(i)
	}
	wg.Wait()
}

func onetest() {
	fmt.Println("执行一次")
}
