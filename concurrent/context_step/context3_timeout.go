package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	go func() {
		defer func() {
			fmt.Println("goroutine exit")
			cancel() //和cancelCtx一样，WithTimeout返回的cancel()是一定要调用的。
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
}
