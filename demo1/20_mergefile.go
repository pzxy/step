package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	root := "/Users/pzxy/WorkSpace/GO/src/letgo_book"
	targetFile := "letge_book.md"
	f, err := os.OpenFile(targetFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println("open err:", err.Error())
		return
	}

	filepath.Walk(root, func(p string, info fs.FileInfo, err error) error {
		if info == nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		if !strings.EqualFold(path.Ext(info.Name()), ".md") {
			return nil
		}
		data, err := ioutil.ReadFile(p)
		_, err = f.Write(data)
		if err != nil {
			fmt.Println("err ==> :", err.Error())
			panic("123")
		}
		return nil
	})
	f.Close()
}
