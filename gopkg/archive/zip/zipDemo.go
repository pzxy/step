package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	//testWrite()
	//testOpenReader()
	testOpenFile()
}
func testWrite() {
	s := `D:\workspace\Go\src\step\gopkg\archive\zip`
	path1 := strings.Replace(s, "\\", "/", -1)
	path2 := path.Join(path1, "test_file.zip")
	// 创建一个缓冲区来写入我们的存档。
	f, err := os.OpenFile(path2, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	// 创建一个新的zip存档。
	w := zip.NewWriter(f)

	// 将一些文件添加到存档中。将一些文件添加到存档中。
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// 确保在Close时检查错误。
	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}
}

//不使用 zip.OpenReader(path) 创建读取对象
func testOpenFile() {
	s := `D:\workspace\Go\src\step\gopkg\archive\zip\test_file.zip`
	path := strings.Replace(s, "\\", "/", -1)
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	r, err := zip.NewReader(f, 1024)
	if err != nil {
		log.Fatalln(err)
	}
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		r, err := f.Open()
		if err != nil {
			log.Fatalln(err)
		}
		buf := make([]byte, 1024)
		io.CopyBuffer(os.Stdout, r, buf)

		fmt.Println()
		fmt.Println("xxxxxxxxxxxxxx")
		fmt.Printf("buf : %s",buf)

		r.Close()
	}
}
func testOpenReader() {
	s := `D:\workspace\Go\src\step\gopkg\archive\zip\test_file.zip`
	path := strings.Replace(s, "\\", "/", -1)
	// 打开一个zip文件供阅读。
	r, err := zip.OpenReader(path)

	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// 遍历存档中的文件，
	// 打印他们的一些内容。
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		//os.Stdout 输出内容到控制台
		_, _ = io.CopyN(os.Stdout, rc, 68)
		//if err == io.EOF {//这里的文件的结束是指当前打开的readme.txt文件，不是当前.zip目录
		//	log.Fatal(err)
		//}
		rc.Close()
		fmt.Println()
	}

}
