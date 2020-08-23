package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	testMd5ByUrl()
	//testMd5ByFile()
}

//byte  string互相转换
func testMd5ByUrl() {
	imgUrl := "ftp://192.168.2.118:9965/Template.bin"
	res, err := http.Get(imgUrl)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("%d \n", len(body))
	md5Url := md5.Sum(body)
	tmp := make([]byte, 0)
	tmp = append(tmp, md5Url[:]...)
	mStr := hex.EncodeToString(tmp)

	fmt.Printf("m1 : %x \n", md5Url)
	fmt.Println("..ssss.....", mStr)

	//src := "5a1f3dbbae52785f58bdea8db575050b"
	//
	//str, _ := hex.DecodeString(src)
	//
	//if mStr == src {
	//	fmt.Println("....chengogn ...", str)
	//}

}
func testMd5ByFile() {
	data1, _ := ioutil.ReadFile("D:/workspace/Go/robot2.bin")
	fmt.Printf("%d \n", len(data1))
	m1 := md5.Sum(data1)
	fmt.Printf("%x \n", m1)
	//data2, _ := ioutil.ReadFile("D:/workspace/Go/robot(3).bin")
	//fmt.Printf("%d \n", len(data2))
	//m2 := md5.Sum(data2)
	//fmt.Printf("%x \n", m2)

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
