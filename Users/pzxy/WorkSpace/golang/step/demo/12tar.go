package main

import (
	"archive/tar"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var filename string
var directory string

func init() {
	flag.StringVar(&filename, "f", "", "the tar filename")
	flag.StringVar(&directory, "dir", "", "the directory you want to tar")
}

func main() {
	filename = "123.tar"
	directory = "/Users/pzxy/WorkSpace/golang/step/demo"

	//flag.Parse()
	if directory == "" {
		fmt.Println("please specify your directory first!")
		return
	}
	if filename == "" {
		filename = filepath.Base(directory)
	}
	// 创建文件
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// 创建一个Writer
	writer := tar.NewWriter(f)
	defer writer.Close()

	// 遍历需要归档的目录
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		// 如果是目录跳过
		if info.IsDir() {
			return nil
		}

		// 打开文件
		f, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			return err
		}

		// 写入文件头
		hr := &tar.Header{
			Name: path,
			Size: info.Size(),
			Mode: 0666,
		}

		// 将文件头写入文件中
		writer.WriteHeader(hr)
		var buff [1024]byte

		// 不断读取文件中的内容并且写入tar文件中
		for {
			n, err := f.Read(buff[:])
			writer.Write(buff[:n])
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
				return nil
			}
		}
		return nil

	})

}
