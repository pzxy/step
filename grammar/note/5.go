package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

var c = make(chan bool, 2)

func main() {

	//for i := 0; i < 10; i++ {
	//	go func(a int) {
	//		c <- true
	//		time.Sleep(time.Second)
	//		fmt.Println(a)
	//		<-c
	//	}(i)
	//	time.Sleep(time.Millisecond * 100)
	//}
	//time.Sleep(time.Second * 10)

	data := []byte("this is test, hello world, keep coding")
	fmt.Printf("%x \n", sha1.Sum(data))

	h := sha1.New()
	io.WriteString(h, "this is test, hello world, keep coding")
	fmt.Printf("%x \n", h.Sum(nil))

	//fmt.Printf("%x \n", shaFile("./file.txt"))

}
