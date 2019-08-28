package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	Url64 := "https://studygolang.com/articles/22602?fr=sidebar"

	testfile2(Url64)
}
func testfile() {
	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//创建一个新的io.Reader，它实现了Read方法
	reader := bufio.NewReader(file)
	//设置读取的长度
	buf := make([]byte, 1024)
	//读取文件
	_, err = reader.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf))
}
func testfile2(url string) {
	//res,_ := http.Get(Url64)
	//创建下载存放exe
	//re := regexp.MustCompile(`([^/]+.bin)`) //must含义 写的肯定是对的 不会出错 表达式加()表示提取
	////match := re.FindAllString(text,-1)//findstring 说明参数是string
	//match := re.FindAllStringSubmatch(url, -1)
	//var name string
	//for _, m := range match {
	//	name = string(m[0])
	//}

	dataAddr := "D:\\" + "360sd-upd-x64.bin"
	f, _ := os.Create(dataAddr)
	defer f.Close()

	res, _ := http.Get(url)
	io.Copy(f, res.Body)

	//	bufio.NewReader()

}
