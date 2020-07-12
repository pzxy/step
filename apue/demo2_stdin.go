package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	stdInRead()

}
func stdInRead() {
	s := "测试os.Stdout.Write"
	b, _ := json.Marshal(s)

	_, err := os.Stdin.Read(b) //Stdin 读read就是从键盘等输入设备读取数据。
	if err != nil {

		fmt.Println("read error:", err)
		return

	}

	//_, err2 := os.Stdout.Write(buffer)
	//if err2 != nil {
	//	fmt.Println("read error:", err)
	//	return
	//
	//}

	//fmt.Println("count:", n, ", msg:", string(buffer[:]))
}

func stdOutWrite() { //Stdout对应屏幕，将数据写到Stdout标准输出中，相当于输出到屏幕
	s := "测试os.Stdout.Write"
	b, _ := json.Marshal(s)
	_, err := os.Stdout.Write(b) //就是fmt.println
	if err != nil {
		fmt.Println("read error:", err)
		return

	}
}

func bufRead() {
	reader := bufio.NewReader(os.Stdout)

	result, err := reader.ReadString('\n')
	if err != nil {

		fmt.Println("read error:", err)
	}

	fmt.Println("result:", result)
}
