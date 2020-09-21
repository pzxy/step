package main

import (
	"fmt"
	"golang.org/x/exp/mmap"
	"os"
	"step/grammar/codeskill/log"

	"syscall"
)

//https://www.cnblogs.com/huxiao-tee/p/4660352.html#4008787
func main() {
	mmapDemo()
}

/**
syscall包中的mmap
*/
func syscallMmapDemo() {
	///
	//f- fd：待映射的文件描述符。
	//- offset：映射到内存区域的起始位置，0 表示由内核指定内存地址。
	//- length：要映射的内存区域的大小。
	//- prot：内存保护标志位，可以通过或运算符`|`组合
	//    - PROT_EXEC  // 页内容可以被执行
	//    - PROT_READ  // 页内容可以被读取
	//    - PROT_WRITE // 页可以被写入
	//    - PROT_NONE  // 页不可访问
	//- flags：映射对象的类型，常用的是以下两类
	//    - MAP_SHARED  // 共享映射，写入数据会复制回文件, 与映射该文件的其他进程共享。
	//    - MAP_PRIVATE // 建立一个写入时拷贝的私有映射，写入数据不影响原文件。

	fd, err := syscall.Open("./mmap.txt", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.ErrLog(err)
		return
	}
	//可以读，但是不知道怎么写
	p, err := syscall.Mmap(fd, 0, 4096, syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		log.ErrLog(err)
		return
	}
	//p, _ = json.Marshal("data")
	fmt.Println(p)
}

//mmap包
//"golang.org/x/exp/mmap"
//go get github.com/golang/exp

//是不是说复制时候很好用呢？
func mmapDemo() {
	r, err := mmap.Open("./mmap.txt")
	if err != nil {
		log.ErrLog(err)
		return
	}
	buff := make([]byte, 10)
	//读入的长度为slice预设的长度，0是offset。预设长度过长将会用0填充。
	r.ReadAt(buff, 0)
	fmt.Println(string(buff))
	r.Close()

}
