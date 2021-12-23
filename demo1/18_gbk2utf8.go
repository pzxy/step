package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/fs"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

// 将文件从gbk转换为utf-8

func main() {
	//s := "/Users/pzxy/WorkSpace/Go/src/Java"
	s := "/Users/pzxy/WorkSpace/java/Java"
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

func showFileList(root string) {
	err := filepath.Walk(root, func(p string, info fs.FileInfo, err error) error {
		if info == nil {
			// 文件名称超过限定长度等其他问题也会导致info == nil
			// 如果此时return err 就会显示找不到路径，并停止查找。
			println("can't find:(" + p + ")")
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if !strings.EqualFold(path.Ext(info.Name()), ".txt") {
			fmt.Println("这不是.txt文件")
			return nil
		}

		//if !strings.EqualFold(path.Ext(info.Name()), ".java") {
		//	fmt.Println("这不是.java文件")
		//	return nil
		//}

		if err := convertFile(p); err != nil {
			fmt.Println(err)
		}

		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() error: %v\n", err)
	}
	return
}
