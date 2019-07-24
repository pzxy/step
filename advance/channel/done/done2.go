package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("test waitGroup : go func %d\n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
