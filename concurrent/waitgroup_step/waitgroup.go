package main

/**
Wait 方法的实现逻辑是：不断检查 state 的值。
如果其中的计数值变为了 0，那么说明所有的任务已完成，调用者不必再等待，直接返回。
如果计数值大于 0，说明此时还有任务没完成，那么调用者就变成了等待者，需要加入 waiter 队列，并且阻塞住自己。
*/

/**
func (wg *WaitGroup) Add(delta int) {
	statep, semap := wg.state()
	// 高32bit是计数值v，所以把delta左移32，增加到计数上
	state := atomic.AddUint64(statep, uint64(delta)<<32)
	v := int32(state >> 32) // 当前计数值
	w := uint32(state) // waiter count

	if v > 0 || w == 0 {
		return
	}

	// 如果计数值v为0并且waiter的数量w不为0，那么state的值就是waiter的数量
	// 将waiter的数量设置为0，因为计数值v也是0,所以它们俩的组合*statep直接设置为0即可。此时需要并唤醒所有的waiter
	*statep = 0
	for ; w != 0; w-- {
		runtime_Semrelease(semap, false, 0)
	}
}


// Done方法实际就是计数器减1
func (wg *WaitGroup) Done() {
	wg.Add(-1)
}
*/
