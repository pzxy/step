package main

import (
	"fmt"
	"sync"
)

func main() {
	maxCount := 100
	dogC := make(chan token, 1)
	catC := make(chan token, 1)
	pigC := make(chan token, 1)
	wg.Add(3)
	go do("dog", maxCount, dogC, catC)
	go do("cat", maxCount, catC, pigC)
	go do("pig", maxCount, pigC, dogC)
	dogC <- token{}
	wg.Wait()
}

var wg sync.WaitGroup

type token struct{}

func do(str string, maxCount int, curr, next chan token) {
	var count int

	for {
		<-curr
		count++
		if count == maxCount {
			next <- token{}
			wg.Done()
			return
		}
		fmt.Println(str)
		next <- token{}

	}
}
