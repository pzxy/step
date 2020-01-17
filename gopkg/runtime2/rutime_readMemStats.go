package main

import (
	"fmt"
	"runtime"
	"time"
)

//ReadMemStats 用内存分配器统计信息填充 m 。
//返回的内存分配程序统计信息在调用 ReadMemStats 时是最新的。这与堆配置文件形成鲜明对比，堆配置文件是最近完成的垃圾回收周期的快照。
type Garbage struct{ a int }

func notify(f *Garbage) {
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	s := time.Unix(0, int64(stats.LastGC)).Format("2006-01-02 15:04:05")
	fmt.Println("Last GC was:", s)
	fmt.Println("Last GC time:", stats.NumGC)
	fmt.Printf("read trace :%s\n", runtime.ReadTrace())
	//go ProduceFinalizedGarbage()
}

func ProduceFinalizedGarbage() {
	x := &Garbage{}
	runtime.SetFinalizer(x, notify) //垃圾回收指定结构体时执行
}

func main() {
	go ProduceFinalizedGarbage()
	for {
		runtime.GC()
		time.Sleep(5 * time.Second) // Give GC time to run
	}
}
