package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	testMd5ByUrl()
}

//byte  string互相转换
func testMd5ByUrl() {
	imgUrl := "https://dl.pconline.com.cn/html_2/1/117/id=10699&pn=0&linkPage=1.html"
	res, err := http.Get(imgUrl)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	m1 := md5.Sum(body)
	fmt.Printf("%x \n", m1)
	//s := string(m1[1])
	md5Byte := make([]byte, 0)
	md5Byte = append(md5Byte, m1[:]...)

	s := string(md5Byte)

	var md5Src [16]byte
	md5SrcStr := s
	copy(md5Src[:], md5SrcStr)

	if m1 == md5Src {
		fmt.Printf("%x", md5Src)
	}

}

//第二种加密写法
func testMd5ByUrl2() {
	imgUrl := "https://music.163.com/#/download"
	res, err := http.Get(imgUrl)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}
	h := md5.New()
	body, err := ioutil.ReadAll(res.Body)
	io.WriteString(h, string(body))
	fmt.Printf("%x", h.Sum(nil))
}

//读取文件加密

func testMd5ByReadFile() {
	imgUrl := "D://dt_data_file1"
	body, _ := ioutil.ReadFile(imgUrl)
	m1 := md5.Sum(body)
	fmt.Printf("%x", m1)
}
func testMd5() {
	s := string("asdasd")
	body := []byte(s)
	m1 := md5.Sum(body)
	fmt.Printf("%x", m1)
}
