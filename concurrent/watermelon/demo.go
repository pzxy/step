package main

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"time"
)

//周一至周五每天早上10点至下午6点提醒喝水
/**
cron
*/
func main() {
	fmt.Println("程序开始时间======", time.Now())
	c := cron.New()
	//spec := "0 0 9,10,11,12,14,15,16,17,18 * * ? "
	//spec := "0 0 9,10,11,12,14,15,16,17,18 * * ?"
	spec1 := "*/5 * * * * ?"
	spec2 := "*/5 * * * * ?"
	err := c.AddFunc(spec1, func() {
		fmt.Println("time=====", time.Now().Location())
		fmt.Println("=======================开始执行喝水任务=======================")

		//url := "http://ep.yoozoo.com/api/feishu/msg"
		//payload := strings.NewReader("{\n\t\"msg_type\": \"text\",\n    \"account\": \"oc_f1d263808c91f074b0eef45a782c92fd\",\n    \"send_obj\": \"chat\",\n    \"text\": \"记得喝水哦\"\n}")
		//req, _ := http.NewRequest("POST", url, payload)
		//
		//req.Header.Add("content-type", "application/json")
		//req.Header.Add("cache-control", "no-cache")

		//_, e := http.DefaultClient.Do(req)
		_, e := json.Marshal("sss")

		if e != nil {
			fmt.Println("http do err:======", e)
			return
		}

	})
	err = c.AddFunc(spec2, func() {
		fmt.Println("=======================开始执行吃饭任务=======================")
		fmt.Println("time=====", time.Now())

	})
	if err != nil {
		fmt.Println(" AddFunc err=====", err)
	}
	c.Start()
	select {}

}
