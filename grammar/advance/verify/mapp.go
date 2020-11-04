package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.Range(func(key, value interface{}) bool {
		fmt.Println("123")
		return true
	})

}
