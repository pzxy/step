package main

/**
刚开始的锁，实现很简单，但是现在很复杂了。
第一版mutex就是一个标志，一个信号量。标志用cas控制，来判断获取或者释放。信号量来告诉协程休眠还是唤醒。
*/
// CAS操作，当时还没有抽象出atomic包
func cas(val *int32, old, new int32) bool {
	//....
	return true
}

func semacquire(*int32) {
	//...
}
func semrelease(*int32) {
	//...
}

// 互斥锁的结构，包含两个字段
type Mutex3 struct {
	key  int32 // 锁是否被持有的标识
	sema int32 // 信号量专用，用以阻塞/唤醒goroutine
}

// 保证成功在val上增加delta的值
func xadd(val *int32, delta int32) (new int32) {
	for {
		v := *val
		if cas(val, v, v+delta) {
			return v + delta
		}
	}
	panic("unreached")
}

// 请求锁
func (m *Mutex3) Lock() {
	if xadd(&m.key, 1) == 1 { //标识加1，如果等于1，成功获取到锁
		return
	}
	semacquire(&m.sema) // 否则阻塞等待
}

func (m *Mutex3) Unlock() {
	if xadd(&m.key, -1) == 0 { // 将标识减去1，如果等于0，则没有其它等待者
		return
	}
	semrelease(&m.sema) // 唤醒其它阻塞的goroutine
}
