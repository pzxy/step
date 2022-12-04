package main

import (
	"fmt"
	"sync"
	"time"
)

type A45 struct {
	b int64
}

func main() {
	a := &A45{
		b: int64(0),
	}
	go func() {
		for {
			fmt.Println(a.b)
		}
	}()
	go func() {
		for i := 0; i < 10000000; i++ {
			a = &A45{
				b: int64(i),
			}
		}
	}()
	time.Sleep(100 * time.Second)
}

func SyncDelete() {
	a := &sync.Map{}
	ReadMap := &sync.Map{}

	a.Store("foo", "bar")
	a.Store("foo1", "baz")
	a.Store("foo2", "2")
	a.Store("kobe3", "3")
	fmt.Println("----->start sync")
	a.Range(func(key, value any) bool {
		ReadMap.Store(key, value)
		return true
	})
	fmt.Println("!!!!end sync")
	PrintSync(ReadMap)

	fmt.Println("=========>add key print")
	a.Store("foo", "bar--------------------")
	fmt.Println("----->start sync")
	a.Range(func(key, value any) bool {
		ReadMap.Store(key, value)
		return true
	})
	fmt.Println("!!!!end sync")
	PrintSync(ReadMap)

	fmt.Println("=========>delete key print")
	a.Delete("kobe3")
	fmt.Println("----->start sync")
	a.Range(func(key, value any) bool {
		fmt.Printf("key: %v, value: %v \n", key, value)
		ReadMap.Store(key, value)
		return true
	})
	fmt.Println("!!!!end sync")
	PrintSync(ReadMap)
}

func PrintSync(read *sync.Map) {
	fmt.Println("@@@----->start print")
	read.Range(func(key, value any) bool {
		fmt.Printf("key: %v, value: %v\n", key, value)
		return true
	})
	fmt.Println("@@!!!!end print")
}
