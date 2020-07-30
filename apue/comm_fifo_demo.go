package main

import (
	"encoding/json"
	"fmt"
	"io"
	"step/utils/log"
	"time"
)

func main() {
	fifoDemo()
}

/**
go中的这个pipe()管道是命名管道，是fifo，不是linux中的pipe,os/exec包才对匿名管道支持
*/
//https://blog.csdn.net/Meng_zj/article/details/80391485?utm_source=blogxgwz3
func fifoDemo() {
	r, w := io.Pipe()
	go func() {
		b, _ := json.Marshal("hello world")
		_, err := w.Write(b)
		if err != nil {
			log.ErrLog(err)
			return
		}
		b2, _ := json.Marshal("hello world2")
		_, err = w.Write(b2)
		if err != nil {
			log.ErrLog(err)
			return
		}
	}()

	time.Sleep(time.Second * 2)
	buf := make([]byte, 512)
	_, err := r.Read(buf)
	if err != nil {
		log.ErrLog(err)
		return
	}

	r.Close()
	w.Close()
	length := len(buf)
	fmt.Printf("%s \n", buf[:length])

}
