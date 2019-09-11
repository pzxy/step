package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"pss/middleware/utils/helper"
	"strconv"
)

func main() {
	//"D://m0_38005557"
	//result := comparefile("D://m0_38005557","D://dt_data_file1")
	//fmt.Printf("结果: %v\n",result)
	var dstByte []byte
	var srcByte []byte
	dstByte = []byte("你好呀")
	srcByte = []byte("我太难了33")
	f := test()
	f.Write(dstByte)
	f.Write(srcByte)
	f.Close()

}
func test()*os.File{
	var f *os.File
	fileSuffix := helper.FormatUint8(1)
	path := "D://dt_data_file" + fileSuffix
	f, _  = os.OpenFile(path,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	//const (
	//
	//                O_RDONLY int = syscall.O_RDONLY // 只读打开文件和os.Open()同义
	//
	//                O_WRONLY int = syscall.O_WRONLY // 只写打开文件
	//
	//                O_RDWR   int = syscall.O_RDWR   // 读写方式打开文件
	//
	//                O_APPEND int = syscall.O_APPEND // 当写的时候使用追加模式到文件末尾
	//
	//                O_CREATE int = syscall.O_CREAT  // 如果文件不存在，此案创建
	//
	//                O_EXCL   int = syscall.O_EXCL   // 和O_CREATE一起使用, 只有当文件不存在时才创建
	//
	//                O_SYNC   int = syscall.O_SYNC   // 以同步I/O方式打开文件，直接写入硬盘.
	//
	//                O_TRUNC  int = syscall.O_TRUNC  // 如果可以的话，当打开文件时先清空文件
	//
	//        )
	//————————————————
	//f, _ = os.Create(path)
	//_, err := os.Stat(path)
	//if err == nil {
	//	f, _ = os.Open(path)
	//}
	//if os.IsNotExist(err) {
	//	f, _ = os.Create(path)
	//}
	//测试时使用,认为目录一定存在
	return f
}
func writeDtData2File() {
	data, err := ioutil.ReadFile("D:/workspace/Go/src/step/basic/file/re.json")
	if err != nil {
		panic(err)
	}
	fileSuffix := strconv.Itoa(int(1))
	path := "D://dt_data_file" + fileSuffix
	f, _ := os.Create(path) //测试时使用,认为目录一定存在
	f.Write([]byte(data))
	defer f.Close()
}

func comparefile(spath, dpath string) bool {
	sFile, err := os.Open(spath)

	if err != nil {

		return false
	}
	dFile, err := os.Open(dpath)

	if err != nil {

		return false
	}
	b := comparebyte(sFile, dFile)


	sFile.Close()
	dFile.Close()

	return b
}
func comparebyte(sfile *os.File, dfile *os.File) bool {

	var sbyte []byte = make([]byte, 512)

	var dbyte []byte = make([]byte, 512)

	var serr, derr error

	for {
		_, serr = sfile.Read(sbyte)
		_, derr = dfile.Read(dbyte)

		if serr != nil || derr != nil {

			if serr != derr {

				return false
			}

			if serr == io.EOF {

				break
			}
		}

		if bytes.Equal(sbyte, dbyte) {

			continue
		}

		return false
	}

	return true

}
