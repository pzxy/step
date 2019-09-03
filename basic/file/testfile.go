package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	//"D://m0_38005557"
	result := comparefile("D://m0_38005557","D://dt_data_file3")
	fmt.Printf("结果: %v\n",result)
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
