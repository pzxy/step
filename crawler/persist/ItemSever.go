package persist

import "log"

func ItemSaver() chan interface{} {
	/**
	创建了一个通道  把通道返回就获得了 这个通道
	下面开启了匿名函数协成时刻监听这个通道获得数据  真是太巧秒了
	*/
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver : got item "+"#%d: %v", itemCount, item)
			itemCount++
		}
	}()
	return out
}
