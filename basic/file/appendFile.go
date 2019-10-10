package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//注意 不能同一个file对象 开两个 br := bufio.NewReader(file)   会有一个为空
	fileNameSrc, _ := os.Open(`D:\workspace\Go\src\step\basic\file\fileSrc`)
	fileNameSrc2, _ := os.Open(`D:\workspace\Go\src\step\basic\file\fileSrc`)
	//testResultFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	fileAppendDst, _ := os.OpenFile(`D:\workspace\Go\src\step\basic\file\fileAppendDst`, os.O_APPEND|os.O_CREATE, os.ModePerm)
	fileCopyDst, _ := os.OpenFile(`D:\workspace\Go\src\step\basic\file\fileCopyDst`, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	defer fileNameSrc.Close()
	defer fileAppendDst.Close()
	defer fileCopyDst.Close()
	copyFile2(fileCopyDst, fileNameSrc)
	result := readCoverFileLine(fileNameSrc2)
	for i := 1; i < len(result); i++ {
		withReturn := fmt.Sprintf("%v\n", result[i])
		fileAppendDst.Write([]byte(withReturn))
	}
}

func readCoverFileLine(file *os.File) (result []string) {
	//之所以这里入参是 io.Reader类型能传入 *os.File类型,
	//是由于*File实现了 io.Reader 这个接口
	br := bufio.NewReader(file) //返回这个文件足够大小的reader对象的缓冲区类型
	fmt.Printf("result %v \n", br)
	for {
		if a, _, c := br.ReadLine(); c != io.EOF { //返回行内容,若是太长的话,返回头的前缀,然后下次继续读
			result = append(result, string(a))
			continue
		}
		fmt.Printf("result %v \n", result)
		break
	}
	return

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
