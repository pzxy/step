package main

import (
	"fmt"
	"github.com/ouqiang/timewheel"
	"time"
)

func main() {
	// 初始化时间轮
	// 第一个参数为tick刻度, 即时间轮多久转动一次
	// 第二个参数为时间轮槽slot数量
	// 第三个参数为回调函数
	var tw *timewheel.TimeWheel
	tw = timewheel.New(1*time.Second, 1, func(data interface{}) {
		m := data.(map[string]int)
		v := m["uid"]
		if v == 2 {
			tw.AddTimer(2*time.Second, nil, map[string]int{"uid": 2})
		}
		fmt.Println("123")
	})

	// 启动时间轮
	tw.Start()

	// 添加定时器
	// 第一个参数为延迟时间
	// 第二个参数为定时器唯一标识, 删除定时器需传递此参数
	// 第三个参数为用户自定义数据, 此参数将会传递给回调函数, 类型为interface{}
	tw.AddTimer(1*time.Second, nil, map[string]int{"uid": 2})
	tw.AddTimer(1*time.Second, nil, map[string]int{"uid": 2})
	tw.AddTimer(1*time.Second, nil, map[string]int{"uid": 2})

	// 删除定时器, 参数为添加定时器传递的唯一标识
	//tw.RemoveTimer(conn)

	// 停止时间轮
	//tw.Stop()
	time.Sleep(time.Second * 10000)
	select {}
}
