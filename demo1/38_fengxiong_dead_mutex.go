package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	fmt.Println("start")
	// 模拟写入神策的文件管道
	//revCh := make(chan int, 50)
	revCh := make(chan int)
	defer close(revCh)

	ch := make(chan struct{}, 5)
	g := new(errgroup.Group)

	for i := 1; i <= 100; i++ {
		ch <- struct{}{}
		n := i
		// Launch a goroutine to fetch the URL.
		g.Go(func() error {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("err:", err)
				}
			}()
			<-ch
			revCh <- i
			fmt.Println("n:--->", n)
			return nil
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}

	fmt.Println("success")
}

func ConsumerMsg(revCh chan int) {
	for {
		select {
		case v, _ := <-revCh:
			fmt.Println("v:---->", v)
		default:
			fmt.Println("none")
		}
	}
}
