package main

import (
	"context"
	"fmt"
	"time"
)

/**
用处：
我们常常在一些需要主动取消长时间的任务时，
创建这种类型的 Context，然后把这个 Context 传给长时间执行任务的 goroutine。
当需要中止任务时，我们就可以 cancel 这个 Context，这样长时间执行任务的 goroutine，
就可以通过检查这个 Context，知道 Context 已经被取消了。

注意：
不是只有你想中途放弃，才去调用 cancel，只要你的任务正常完成了，
就需要调用 cancel，这样，这个 Context 才能释放它的资源（通知它的 children 处理 cancel，
从它的 parent 中把自己移除，甚至释放相关的 goroutine）。
很多同学在使用这个方法的时候，都会忘记调用 cancel，切记切记，而且一定尽早释放。

原理：
代码中调用的 propagateCancel 方法会顺着 parent 路径往上找，
直到找到一个 cancelCtx，或者为 nil。如果不为空，就把自己加入到这个 cancelCtx 的 child，
以便这个 cancelCtx 被取消的时候通知自己。如果为空，会新起一个 goroutine，
由它来监听 parent 的 Done 是否已关闭。

取消的时候是向下传递，和查值不一样，值是向上传递。
*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer func() {
			fmt.Println("goroutine exit")
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

	time.Sleep(time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}
