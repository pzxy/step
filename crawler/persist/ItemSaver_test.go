package persist

import (
	"context"
	"encoding/json"
	"godemo/crawler/engine"
	"godemo/crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

/**
测试保存方法
*/
func TestItemSaver(t *testing.T) {
	//TODO :try to start up elastic search
	//here using docker go client.
	client, err := elastic.NewClient(elastic.SetSniff(false))

	expected := engine.Item{
		Url:  "http://asdasd",
		Type: "zhenai",
		Id:   "231",
		Payload: model.Profile{
			Name:   "胡汉三",
			Age:    34,
			Height: 170,
			Weight: 60,
		},
	}
	const index = "dating_test"
	//Save expected item
	err = save(client, index, expected)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index(index).Type(expected.Type).Id(expected.Id).Do(context.Background())

	if err != nil {
		panic(err)
	}
	t.Logf("%s", *resp.Source)

	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if err != nil {
		panic(err)
	}
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}

}
