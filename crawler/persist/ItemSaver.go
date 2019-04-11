package persist

import (
	"context"
	"github.com/kataras/iris/core/errors"
	"godemo/crawler/engine"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false)) //must turn off sniff in docker

	if err != nil {
		return nil, err
	}

	/**
	创建了一个通道  把通道返回就获得了 这个通道
	下面开启了匿名函数协成时刻监听这个通道获得数据  真是太巧秒了
	*/
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver : got item "+"#%d: %v", itemCount, item)
			itemCount++

			err := save(client, index, item)
			if err != nil {
				log.Printf("Item Saver : error "+"saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, index string, item engine.Item) error {
	client, err := elastic.NewClient(elastic.SetSniff(false)) //must turn off sniff in docker
	if err != nil {
		return nil
	}
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	//determinr the location by index & type & id
	indexService := client.Index(). //插入 or 更改
					Index(index).
					Type(item.Type).
					BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.Do(context.Background())
	if err != nil {
		return nil
	}
	return nil
}
