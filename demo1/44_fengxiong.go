package main

import (
	"fmt"
	"sync"
)

var b = &sync.Map{}

func main() {
	a := &sync.Map{}
	c := &sync.Map{}
	c.Store(1, "1")
	c.Store(2, "2")
	a.Store("a", c)

	b := &sync.Map{}
	a.Range(func(key, value any) bool {
		val := value.(*sync.Map)
		m := &sync.Map{}
		val.Range(func(key, value any) bool {
			m.Store(key, value)
			return true
		})
		b.Store(key, m)
		return true
	})

	c.Delete(1)

	val, ok := b.Load("a")
	if !ok {
		fmt.Println("fail")
	}

	val.(*sync.Map).Range(func(key, value any) bool {
		fmt.Println("key:, val:", key, value)
		return true
	})

	return
}
