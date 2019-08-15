package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("test waitGroup : go func %d\n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
