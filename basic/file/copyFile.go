package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	srcFile, err := os.Open(`D:\workspace\Go\src\step\basic\file\fileSrc`)
	dstFile, err := os.OpenFile(`D:\workspace\Go\src\step\basic\file\fileCopyDst`, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil{
		panic(err)
	}
	copyFile2(dstFile,srcFile)
	srcFile.Close()
	dstFile.Close()
	//copyFile()

}

func copyFile() (written int64, err error) {
	srcFile, _ := os.Open(`D:\workspace\Go\src\step\basic\file\fileSrc`)
	dstFile, _ := os.OpenFile(`D:\workspace\Go\src\step\basic\file\fileDst`, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(dstFile)
	srcFile.Close()
	dstFile.Close()
	return io.Copy(writer, reader)
}

func copyFile2(dstFile *os.File, srcFile *os.File) {
	reader := bufio.NewReader(srcFile)
	for {
		if a, _, c := reader.ReadLine(); c != io.EOF {
			writeStr := fmt.Sprintf("%v\n", string(a))
			dstFile.Write([]byte(writeStr))
			continue
		}
		break
	}
}
