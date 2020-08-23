package main

import "fmt"

//结构体实现接口

type DataWriter interface {
	WriteData(data interface{}) error
}

type file struct {
}

func (d *file) WriteData(data interface{}) error {
	//模拟写入数据
	fmt.Println("writedata", data)
	return nil
}

func main() {
	f := &file{}      //new(file)  实例化的两种方式
	f.WriteData("sd") //自己调用

	var dataWriter DataWriter
	ff, isFile := dataWriter.(*file) //dataWriter.(DataWriter) 左侧必须是接口类型 才能断言 右侧为type(结构体指针  或者  接口类型)
	fmt.Println(ff, isFile)

	dataWriter = f
	dataWriter.WriteData("asd") //接口调用
}
