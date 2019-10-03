package main

import (
	"fmt"
	"os"
)

//os包,操作入参若不是绝对路径的话,就是当前路径

func main() {
	//osStatDemo()
	//osRemoveDemo()
	//osRemoveAllDemo()
	//osMkdirAllDemo()
	osOpenDemo()
}

//文件查看
func osStatDemo() {
	fileInfo, err := os.Stat(`C:/Users/pzxy/Desktop/test`)
	if err != nil {
		panic(err)
	}
	fmt.Println(fileInfo.Name())    //文件名字
	fmt.Println(fileInfo.IsDir())   //是否是目录
	fmt.Println(fileInfo.Mode())    //文件的权限信息   drwxrwxrwx
	fmt.Println(fileInfo.ModTime()) //修改时间
	fmt.Println(fileInfo.Size())    //大小
	fmt.Println(fileInfo.Sys())     //文件的数据源，就是 syscall.Win32FileAttributeData 包括文件属性位，三个时间戳，大小高低位
}

//文件删除
func osRemoveDemo() {
	fileInfo, err := os.Stat(`C:/Users/pzxy/Desktop/test`)
	if err != nil {
		panic(err)
	}

	if err := os.Remove(fileInfo.Name()); err != nil { //只能删除文件,不能删除目录
		log(fmt.Sprintf("delete touch failed for (%v)", err))
	}

}

//删除目录

func osRemoveAllDemo() {
	fileInfo, err := os.Stat(`C:/Users/pzxy/Desktop/test`)
	if err != nil {
		panic(err)
	}

	if err := os.RemoveAll(fileInfo.Name()); err != nil { //删除目录
		log(fmt.Sprintf("delete file failed for (%v)", err))
	}

}

//创建目录
func osMkdirAllDemo() { //目录路径和文件的权限
	if err := os.MkdirAll(`C:/Users/pzxy/Desktop/test`, os.ModePerm); err != nil {
		log(fmt.Sprintf("mkdir file failed for (%v)", err))
	}
}

//打开文件 里面调用openFile()
func osOpenDemo() {
	var f *os.File
	f, err := os.Open(`C:/Users/pzxy/Desktop/test`) //绝对或者相对路径
	if err != nil {
		log(fmt.Sprintf("open file failed for (%v)", err))
	}
	fmt.Printf("file name %s \n",f.Name())
	//fmt.Printf("file name %s \n",f.Stat())//类似os.Stat 返回文件的具体属性
	fmt.Printf("file name %s \n",f.Chdir())
	//fmt.Printf("file name %s \n",f.Chmod())//重新改变文件的权限信息和os.chmod一样,由于file的chmod已经知道文件路径了,所致只有一个参数
}
func log(msg string) {
	fmt.Println(msg)
}
