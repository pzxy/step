package persist

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

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

func save(item interface{}) {
	client, err := elastic.NewClient(elastic.SetSniff(false)) //must turn off sniff in docker
	if err != nil {
		panic(err)
	}
	//determinr the location by index & type & id
	resp, err := client.Index(). //插入 or 更改
					Index("dating_profile").
					Type("zhenai").
					BodyJson(item).
					Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
