package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() //处理
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error : Status Code ", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body) //获取数据
}
