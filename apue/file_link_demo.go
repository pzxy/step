package main

import (
	"fmt"
	"os"
	"step/grammar/codeskill/log"
	"syscall"
)

func main() {
	symlinkDemo()
}

/**
link创建一个硬连接,除了文件名字以为一摸一样，包括大小。可以使用stat查看
硬连接数也会加一

硬链接不允许跨文件系统，到那时posix支持

符号链接没有文件系统限制

不允许创建目录的硬连接，因为可能循环调用。

增加硬连接数应该是一个原子操作，为了保证线程安全
*/
func linkDemo() {
	err := os.Link("./link1", "./link1")
	if err != nil {
		log.ErrLog(err)
		return
	}
}

/**
创建一个符号链接，或者说是软链接
使用stat查看后，发现硬连接数没有增加，大小也不一样。
*/
func symlinkDemo() {
	err := os.Symlink("./link1", "./sym_link")

	if err != nil {
		log.ErrLog(err)
		return
	}
}

/**
他只读符号链接文件，也就是软链接文件,然后返回符号链接指向的文件名字
只读文件名字，不读内容，然后返回长度
*/
func readlinkDemo() {
	buf := make([]byte, 1024)

	n, err := syscall.Readlink("./sym_link", buf) //buf中应为link1
	if err != nil {
		log.ErrLog(err)
		return
	}

	if n < 0 {
		return
	}
	//打印出来不好看，os包中的这个函数更好用一点
	fmt.Println(string(buf))

	s, err := os.Readlink("./sym_link")
	if err != nil {
		log.ErrLog(err)
		return
	}
	fmt.Println(s)
}

/**
unlink
1 如果是符号链接就删除符号链接
2 如果是硬连接，硬连接数减1，为0时，释放数据块和inode
3 如果硬连接数为零，但是有进程正打开文件，那么等进程结束后再删除,这样creat一个文件后马上 unlink，就相当于一个临时文件，等进程结束就会删除
*/

/**
发现一特特别神奇的事情，将源文件删除后，通过readlink读取符号链接，得到源名字后。我们再创建一个和源文件一样的文件，随便写的什么东西，这个时候再看符号链接文件就能看到以前的信息了。
这里我理解错了，之所以会这样，是因为goland缓存了数据，我关了又打开，只能看到新数据，看不到原来的数据。
*/
func unlinkDemo() {
	//这里软硬链接要好好理解，软连接mac下我测试的，源文件被删除后，再创建一个同名文件，软连接有重新起作用了也可以用
	err := syscall.Unlink("./link1") //会删除软连接，或者入参为硬连接，删除入参文件，硬连接减1
	if err != nil {
		log.ErrLog(err)
		return
	}

}

/**

rename
文件重命名
*/

func renameDemo() {
	os.Rename("./link1", "./link1")
}
