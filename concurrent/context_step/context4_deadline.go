package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))

	go func() {
		defer func() {
			fmt.Println("WithDeadline goroutine exit")
			cancel() //和cancelCtx一样，WithDeadline 返回的cancel()是一定要调用的。
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				//time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
}
