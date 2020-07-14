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
	statDemo()
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
