package main

import (
	"encoding/json"
	"fmt"
	"io"
	"step/utils/log"
)

func main() {
	fifoDemo()
}

/**
go中的这个pipe()管道是命名管道，是fifo，不是linux中的pipe
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
	}()
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
