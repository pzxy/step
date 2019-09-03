package main

import (
	"io/ioutil"
	"os"
	"strconv"
)

func writeDtData2File(dtData []byte,robotId uint8){
	fileSuffix := strconv.Itoa(int(robotId))
	path := "D://dt_data_file"+fileSuffix
	f, _ := os.Create(path)//测试时使用,认为目录一定存在
	f.Write([]byte(dtData))
	defer f.Close()
}
func main (){
	data, err := ioutil.ReadFile("D:/workspace/Go/src/step/basic/file/re.json")
	if err != nil {
		panic(err)
	}
	writeDtData2File(data,1)
}