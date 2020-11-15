package main

import (
	"fmt"
	"time"
)

/**
场景4：
锁

要想使用 chan 实现互斥锁，至少有两种方式。
一种方式是先初始化一个 capacity 等于 1 的 Channel，
然后再放入一个元素。这个元素就代表锁，谁取得了这个元素，就相当于获取了这把锁。
另一种方式是，先初始化一个 capacity 等于 1 的 Channel，
它的“空槽”代表锁，谁能成功地把元素发送到这个 Channel，谁就获取了这把锁。


你可以用 buffer 等于 1 的 chan 实现互斥锁，在初始化这个锁的时候往 Channel 中先塞入一个元素，
谁把这个元素取走，谁就获取了这把锁，把元素放回去，就是释放了锁。
元素在放回到 chan 之前，不会有 goroutine 能从 chan 中取出元素的，这就保证了互斥性。

在这段代码中，还有一点需要我们注意下：利用 select+chan 的方式，很容易实现 TryLock、Timeout 的功能。
具体来说就是，在 select 语句中，我们可以使用 default 实现 TryLock，使用一个 Timer 来实现 Timeout 的功能。
*/

// 使用chan实现互斥锁
type Mutex struct {
	ch chan struct{}
}

// 使用锁需要初始化
func NewMutex() *Mutex {
	mu := &Mutex{make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

// 请求锁，直到获取到
func (m *Mutex) Lock() {
	<-m.ch
}

// 解锁
func (m *Mutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

// 尝试获取锁
func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}

// 加入一个超时的设置
func (m *Mutex) LockTimeout(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-m.ch:
		timer.Stop()
		return true
	case <-timer.C:
	}
	return false
}

// 锁是否已被持有
func (m *Mutex) IsLocked() bool {
	return len(m.ch) == 0
}

func main() {
	m := NewMutex()
	ok := m.TryLock()
	fmt.Printf("locked v %v\n", ok)
	ok = m.TryLock()
	fmt.Printf("locked v %v\n", ok)
}
