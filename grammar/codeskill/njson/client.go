package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	httpCli = &http.Client{
		Timeout: time.Second * 5,
	}

	cache          [][]byte
	cacheCnt       int
	cacheUploadMax = 100 * 1024
	inputCh        = make(chan []byte, 64)
)

const (
	minGZSize = 1024

	httpDiv = 100
	httpOk  = 2
	httpBad = 4
	httpErr = 5

	serverUrl = "http://0.0.0.0:9090"
)

var testData = []byte(`{"Name": "阿斯顿"}`)

func main() {
	// 启动IO
	go startIO()

	// mock test data
	for i := 0; i < 100000; i++ {
		//time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		time.Sleep(time.Millisecond * 10)
		Write(testData)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
}

func Write(msg []byte) (err error) {
	err = doFeed([]byte(msg))

	return err
}

func doFeed(data []byte) error {
	select {
	// send data
	case inputCh <- data:
	}

	return nil
}

func doFlush(bodies [][]byte, url string) error {
	if bodies == nil {
		return nil
	}

	body, gz, err := buildBody(bodies)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
		return err
	}

	if gz {
		req.Header.Set("Content-Encoding", "gzip")
	}

	resp, err := httpCli.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}

	defer resp.Body.Close()
	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	switch resp.StatusCode / httpDiv {
	case httpOk:
		log.Printf("post to %s ok\n", url)
	case httpBad:
		log.Printf("post to %s failed(HTTP: %d): %s, data dropped \n", url, resp.StatusCode, string(respbody))
	case httpErr:
		log.Printf("post to %s failed(HTTP: %d): %s \n", url, resp.StatusCode, string(respbody))
		return fmt.Errorf("server internal error")
	}

	// 初始化cache
	cache = nil
	cacheCnt = 0

	return nil
}

func startIO() {
	f := func() {
		tick := time.NewTicker(4 * time.Second)
		defer tick.Stop()
		log.Printf("io interval: %v", 1*time.Second)

		for {
			select {
			case d := <-inputCh:
				if d == nil {
					log.Println("get empty data, ignored")
				} else {
					cache = append(cache, d)
					cacheCnt += len(d)

					if cacheCnt >= cacheUploadMax {
						doFlush(cache, serverUrl)
					}
				}

			case <-tick.C:
				doFlush(cache, serverUrl)
			}
		}
	}

	log.Println("starting...")
	f()
}

func buildBody(bodies [][]byte) (body []byte, gzon bool, err error) {
	body = bytes.Join(bodies, []byte("\n"))

	// gzip压缩处理(todo)
	return
}
