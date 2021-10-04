package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
	"path/filepath"
	"unicode/utf8"
)

// 将文件从gbk转换为utf-8

func main() {
	//s := "/Users/pzxy/WorkSpace/Go/src/Java"
	s := "/Users/pzxy/WorkSpace/Go/src/Java"
	showFileList(s)
}

func convertFile(path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if utf8.Valid(b) {
		return fmt.Errorf("已经是utf-8了")
	}

	u, err := GbkToUtf8(b)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, u, 0644)
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if info == nil {
		// 文件名称超过限定长度等其他问题也会导致info == nil
		// 如果此时return err 就会显示找不到路径，并停止查找。
		println("can't find:(" + path + ")")
		return nil
	}
	if info.IsDir() {
		//println("This is folder:(" + path + ")")
		return nil
	} else {
		if err := convertFile(path); err != nil {
			fmt.Println(err)
		}
		return nil
	}
}

func showFileList(root string) {
	err := filepath.Walk(root, walkFunc)
	if err != nil {
		fmt.Printf("filepath.Walk() error: %v\n", err)
	}
	return
}
