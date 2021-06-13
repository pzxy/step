package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	f, _ := os.OpenFile("./prodemo.txt", os.O_CREATE|os.O_RDWR, 0755)
	b, _ := json.Marshal("测试 startProcess")
	f.Write(b)
	time.Sleep(time.Second)
	fmt.Sprint()
}
