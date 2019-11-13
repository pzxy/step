package main

import (
	"archive/zip"
	"compress/flate"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	// 用较高的压缩级别覆盖默认的Deflate压缩器。
	//创建一个缓冲区来写入我们的文档
	//buf := new(bytes.Buffer)
	s := `D:\workspace\Go\src\step\gopkg\archive\zip`
	path1 := strings.Replace(s, "\\", "/", -1)
	path2 := path.Join(path1, "test_file.zip")
	// 创建一个缓冲区来写入我们的存档。
	f, err := os.OpenFile(path2, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	//创建一个新的写入文档
	w := zip.NewWriter(f)
	// 注册一个自定义的Deflate压缩机。
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (closer io.WriteCloser, e error) {
		return flate.NewWriter(out, flate.BestCompression)
	})
	// 继续向w添加文件。
	files := []struct {
		Name, Body string
	}{
		{"readme2.txt", "This archive contains some text files."},
		{"gopher2.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo2.txt", "Get animal handling license."},
	}

	for _, file := range files {
		fw, err := w.Create(file.Name)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = fw.Write([]byte(file.Body))
		if err != nil {
			log.Println(err)
		}
	}
	err = w.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
