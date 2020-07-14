package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"step/utils/log"
	"syscall"
	"time"
)

func main() {
	lstatDemo()
}

func statDemo() {
	/**
	os 包中的stat
	*/
	file, err := os.Stat("./test3.txt")
	if err != nil {
		log.ErrLog(err)
		return
	}
	fmt.Printf("name:%v,\n isDir:%v,\n mode:%v,\n modTime:%v,\n Size:%v,\n Sys:%+v \n", file.Name(),
		file.IsDir(), file.Mode(), file.ModTime(), file.Size(), file.Sys()) //os调用 Sys就相当于Unix下的Stat

	/**
	syscall包中的stat
	*/
	var stat syscall.Stat_t
	err = syscall.Stat("./test.txt", &stat)
	fmt.Println()
	fmt.Println("************** syscall 包中的 stat **************")
	if err != nil {
		log.ErrLog(err)
		return
	}
	//复习一下时间错转时间
	fmt.Println(time.Unix(stat.Atimespec.Sec, 0).Format("2006/01/02 15:04:05"))
	fmt.Printf("%+v \n", stat)

	/**
	unix包中的stat,这两个包内容其实差不多

	*/
	fmt.Println()
	fmt.Println("************** unix 包中的 stat **************")
	var stat2 unix.Stat_t
	err = unix.Stat("./test.txt", &stat2)
	if err != nil {
		log.ErrLog(err)
		return
	}
	//复习一下时间错转时间
	fmt.Println(time.Unix(stat.Atimespec.Sec, 0).Format("2006/01/02 15:04:05"))
	fmt.Printf("%+v", stat)
}

/**
fstat和stat一样，只不过fstat是操作打开的文件，通过文件描述符来操作，而stat是操作为打开的文件
*/
func fstatDemo() {
	fd, err := syscall.Open("./test.txt", syscall.O_CREAT|syscall.O_RDWR, 0655)
	if err != nil {
		log.ErrLog(err)
		return
	}
	var stat syscall.Stat_t
	syscall.Fstat(fd, &stat)
	fmt.Printf("fstat : %+v", stat)
}

/**
ln -s test1.txt ltest1
当执行上面的命令的时候，会创建一个文件，这个文件是指向test1.txt的,相当与一个快捷方式

stat操作的话会操作指向的符号链接的文件test1.txt， 而lstat不跟踪符号链接，而是操作文件的本身
*/
func lstatDemo() {
	//可以看到打开相同的文件，文件的大小是不一样的。stat是打开文件符号链接指向的文件，而lstat是获取文件的本身
	f1, err := os.Stat("./ftest")
	if err != nil {
		log.ErrLog(err)
		return
	}
	fmt.Printf("stat size:%v\n", f1.Size())
	f2, err2 := os.Lstat("./ftest")
	if err2 != nil {
		log.ErrLog(err)
		return
	}
	fmt.Printf("fstat size:%v\n", f2.Size())
	//此外还有 syscall unix包也可以使用
}
