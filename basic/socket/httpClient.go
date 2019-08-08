package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//
func main() {

	test02()
}

func test01() {

	//这里演示了http.Get和httpClient类型Get
	url1 := "http://www.baidu.com"
	fmt.Printf("Send request to %q with method GET ...\n", url1)
	//这里添加Transport,里面包括了请求时间,代理很多设置.
	//若是不设置的话,他会有默认值,DefaultTransport的值
	myTransport := &http.Transport{
		Proxy:                  nil,
		DialContext:            nil,
		Dial:                   nil,
		DialTLS:                nil,
		TLSClientConfig:        nil,
		TLSHandshakeTimeout:    0,
		DisableKeepAlives:      false,
		DisableCompression:     false,
		MaxIdleConns:           0,
		MaxIdleConnsPerHost:    0,
		MaxConnsPerHost:        0,
		IdleConnTimeout:        0,
		ResponseHeaderTimeout:  0,
		ExpectContinueTimeout:  0,
		TLSNextProto:           nil,
		ProxyConnectHeader:     nil,
		MaxResponseHeaderBytes: 0,
	}
	myclient := http.Client{
		Transport: myTransport,
		Timeout:   1 * time.Second,
	}
	resp1, err := myclient.Get(url1)
	//var httpClient http.Client//和http.Get()的等价,http.Get()底层就是Client

	if err != nil {
		fmt.Printf("request sending error: %v\n", err)
	}
	defer resp1.Body.Close()
	line1 := resp1.Proto + " " + resp1.Status
	fmt.Printf("The first line of response:\n%s\n", line1)
}

func test02() {
	//这里演示分步请求
	//先是用http.NewRequest方法指定了请求方法以及请求url构造请求对象req，
	//并进一步对req的Header进行设置，最后再通过Client.Do方法将请求发出去。
	client := &http.Client{}

	//生成要访问的url
	//提交请求
	reqest, err := http.NewRequest("GET", "http://www.baidu.com", nil)

	//设置其他内容,或者在创建client的时候设置 client := &http.Client{transport:...}
	/*Transport :=	&http.Transport{
			Proxy:                  nil,
			DialContext:            nil,
			Dial:                   nil,
			DialTLS:                nil,
			TLSClientConfig:        nil,
			TLSHandshakeTimeout:    0,
			DisableKeepAlives:      false,
			DisableCompression:     false,
			MaxIdleConns:           0,
			MaxIdleConnsPerHost:    0,
			MaxConnsPerHost:        0,
			IdleConnTimeout:        0,
			ResponseHeaderTimeout:  0,
			ExpectContinueTimeout:  0,
			TLSNextProto:           nil,
			ProxyConnectHeader:     nil,
			MaxResponseHeaderBytes: 0,
		}
	client.Transport = Transport*/

	//设置超时时间
	client.Timeout = time.Second * 1
	//设置
	if err != nil {
		return
	}
	//增加header选项
	//reqest.Header.Add("Cookie", "Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1562577869;sid=02b77da0-94ec-4b76-a1dc-6964f872284d;Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1562640252")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36")
	/*reqest.Header.Add("Upgrade-Insecure-Requests", "1")*/
	resp, _ := client.Do(reqest)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {

		//return nil, fmt.Errorf("Error : Status Code ", resp.StatusCode) //这里相当于创建了一个error对象
	}
	//读取返回体
	bs, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bs))

}
