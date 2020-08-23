package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//文件的读取
//readLine更低级,是每行的度,太长的话返回前缀,下次再度
//readByte跟高级,单个字节的读取
//readString 可以指定读取的分界符 单引号 '\n'
func main() {
	file, err := os.Open(`D:\workspace\Go\src\step\basic\file\fileSrc`)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	result := readLine(file)
	//result := readByte(file)
	//result := readString(file)
	fileNameDst, err := os.OpenFile(`D:\workspace\Go\src\step\basic\file\fileDst`, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer fileNameDst.Close()

	fmt.Printf("read file data %s \n", result)
}

func readLine(file *os.File) (result []string) {
	//之所以这里入参是 io.Reader类型能传入 *os.File类型,
	//是由于*File实现了 io.Reader 这个接口
	br := bufio.NewReader(file) //返回这个文件足够大小的reader对象的缓冲区类型
	for {
		if a, _, c := br.ReadLine(); c != io.EOF { //返回行内容,若是太长的话,返回头的前缀,然后下次继续读
			result = append(result, string(a))
			continue
		}
		break
	}
	return
}

//ReadBytes('\n')或ReadString('\n')，或者使用扫描器。相比ReadLine更高级
func readByte(file *os.File) (result []byte) {
	//之所以这里入参是 io.Reader类型能传入 *os.File类型,
	//是由于*File实现了 io.Reader 这个接口
	br := bufio.NewReader(file) //返回这个文件足够大小的reader对象的缓冲区类型
	//br.ReadByte()
	for {
		if a, err := br.ReadByte(); err == nil { //返回行内容,若是太长的话,返回头的前缀,然后下次继续读
			result = append(result, a)
			continue
		}
		break
	}
	return
}

func readString(file *os.File) (result []string) {
	//之所以这里入参是 io.Reader类型能传入 *os.File类型,
	//是由于*File实现了 io.Reader 这个接口
	br := bufio.NewReader(file) //返回这个文件足够大小的reader对象的缓冲区类型
	br.ReadByte()
	for {
		if a, err := br.ReadString('\n'); err == nil { //返回行内容,若是太长的话,返回头的前缀,然后下次继续读
			result = append(result, a)
			continue
		}
		break
	}
	return
}
