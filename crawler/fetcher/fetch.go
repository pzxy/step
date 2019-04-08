package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/**
爬取方法
*/
var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	client := &http.Client{}
	//生成要访问的url
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	//增加header选项
	reqest.Header.Add("Cookie", "Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1554643975; sid=d617da78-0a5c-41ec-89c3-d39d5b02e3f3; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1554644048")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36")
	reqest.Header.Add("Upgrade-Insecure-Requests", "1")
	resp, _ := client.Do(reqest)
	//resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() //处理

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error : Status Code ", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body) //获取数据
}
