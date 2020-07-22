package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"step/utils/log"
	"syscall"
)

/**

目录操作
mkdir,rmdir,opendir/fdopendir,readdir,rewinddir,telldir/seekdir,closedir,dup
*/
func main() {
	dupDemo()
}

/**
创建dir
*/
func mkdirDemo() {
	if err := os.Mkdir("./dir1/dir2", 0755); err != nil {
		log.ErrLog(err)
	}
}

/**
删除目录，只能删除空目录
也就是说删除的时候先删除文件，在删除目录，或者说删除目录下的目录中文件，递归
*/
func rmdirDemo() {
	if err := syscall.Rmdir("./dir1"); err != nil {
		log.ErrLog(err)
		return
	}
}

/**
打开目录
*/

func opendirDemo() {
	//不存在这个 opendir，可能是因为go认为用readir就够了吧
}

/**
读取目录
*/
func readirDemo() {
	file, err := ioutil.ReadDir("./dir1")
	if err != nil {
		log.ErrLog(err)
		return
	}
	fmt.Printf("file:%+v \n", file[1])

	fd, err := syscall.Open("./test.txt", os.O_RDWR, 0755)
	if err != nil {
		log.ErrLog(err)
		return
	}
	buf := make([]byte, 1024)
	//todo 这里要传进去一个文件描述符。但是目录没有文件描述符吧。这里整不了
	n, err := syscall.ReadDirent(fd, buf)
	if err != nil {
		log.ErrLog(err)
		return
	}
	fmt.Println(n)
}

/**

rewinddir
将目录指针恢复到目录的起始位置

*/

func rewinddirDemo() {
	//todo 公众也没找到这个名字得到方法，不知道有没有别的名称相近的方法。

}

/**
telldir/seekdir
设置目录指针位置
*/

func telldirDemo() {
	//todo  go中这两个也没有
}

/**
closedir
*/
func closedirDemo() {
	//todo go中也没有这个

}

/**
递归遍历目录
*/

func traverseDir() {
	file, err := ioutil.ReadDir("./dir1")
	if err != nil {
		log.ErrLog(err)
		return
	}
	fmt.Printf("file:%+v \n", file[1])
	//todo 按理说这里应该是目录的
	if (file[1].Mode() & syscall.S_IFMT) == syscall.S_IFDIR {
		fmt.Println("是目录")
	}
}

func traverseDir2(path1 string) {
	file, err := ioutil.ReadDir(path1)
	if err != nil {
		log.ErrLog(err)
		return
	}
	for _, fi := range file {
		if fi.IsDir() {
			fmt.Println("-" + fi.Name())
			traverseDir2(path.Join(path1, fi.Name()))
		} else {
			fmt.Println(fi.Name())
		}
	}
}

/**
dup 复制一个文件描述符，新的文件描述符fd值自动增加。
dup2  复制一个文件描述符，但是要指定新的fd1，如果指定的fd1已经被人使用，那么就关闭这个文件，然后再使用fd1。
也就是说dup2将指定的第二个参数fd重定向到前面的fd指向的内容。
*/
func dupDemo() {

	b1, _ := json.Marshal("文件重定向1")
	b2, _ := json.Marshal("文件重定向2")
	fd, _ := syscall.Open("./test.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
	saveFd, _ := syscall.Dup(syscall.Stdout)

	syscall.Dup2(fd, syscall.Stdout)
	//向文件描述符写，但是屏幕没有输出
	syscall.Write(syscall.Stdout, b1)

	//再次将保存的fd重新设置到syscall.Stdout上
	syscall.Dup2(saveFd, syscall.Stdout)
	//可以看到有输出了,而且 b1 输入到了 fd，test.txt 文件中了
	syscall.Write(syscall.Stdout, b2)
	syscall.Write(saveFd, b2)

	syscall.Close(fd)

}
