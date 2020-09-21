package main

import (
	"os"
	"step/grammar/codeskill/log"
	"syscall"
	"time"
)

func main() {
	chmodDemo()
}

func chmodDemo() {
	// 执行时设置用户id，当执行这个文件的时候会获得root权限。比如passwd这个文件就设置了这个值，所以执行的时候就有了临时root权限。能够改变shadow中的密码值。
	//	S_ISUID                           = 0x800

	//	S_ISGID                           = 0x400  执行时设置组id

	//设置了这个以后，当内存不够用时交换文件的时候，会尽量将这个文件留在内存中，现在基本上很少使用
	//	S_ISVTX                           = 0x200  黏住位，

	//其他人读写执行
	//	S_IROTH                           = 0x4
	//	S_IWOTH                           = 0x2
	//	S_IXOTH                           = 0x1

	//组读写执行
	//	S_IRGRP                           = 0x20
	//	S_IWGRP                           = 0x10
	//	S_IXGRP                           = 0x8

	//用户读写执行
	//	S_IRUSR                           = 0x100
	//	S_IWUSR                           = 0x80
	//	S_IXUSR                           = 0x40
	//syscall.Chmod("./test2.txt", syscall.S_ISVTX)

	//执行下面代码后，通过stat指令，可以看到文件的权限为被改变了。
	syscall.Chmod("./test2.txt", 0755) //这里必须加行0，因为这个值是八进制才行
	//syscall.FChmod() //通过文件描述符
	//syscall.LChmod() //不跟踪符号链接

}

/**
更改用户id和用户组id，为了安全，只有root用户才有使用这个函数
*/
func chownDemo() {
	err := syscall.Chown("./test2.txt", 1, 1)
	if err != nil {
		log.ErrLog(err)
		return
	}
	//syscall.Chown() //更改用户id和用户组id
	//syscall.Lchown()//通过文件描述符
	//syscall.Fchown()//不跟踪符号链接
}

/**
type Stat_t struct {
	Dev           int32
	Mode          uint16
	Nlink         uint16
	Ino           uint64
	Uid           uint32
	Gid           uint32
	Rdev          int32   //??
	Pad_cgo_0     [4]byte //?? 这个值不知道是多少，这里设置我打印出来是0000，所以下面使用的时候也用0000
	Atimespec     Timespec
	Mtimespec     Timespec
	Ctimespec     Timespec
	Birthtimespec Timespec
	Size          int64
	Blocks        int64
	Blksize       int32
	Flags         uint32
	Gen           uint32 // ??
	Lspare        int32  // ??
	Qspare        [2]int64  // ??
}
*/
func UTimesDemo() {

	Pad_cgo_0 := [4]byte{0, 0, 0, 0}
	timeVals := []syscall.Timeval{
		syscall.Timeval{int64(time.Now().Second()), 0, Pad_cgo_0},
	}
	//todo 这里是无效的参数，说实话，我也不知道怎么去用这个。只能先记住，utime是更改文件的改变时间呢
	err := syscall.Utimes("./test2.txt", timeVals)
	if err != nil {
		log.ErrLog(err)
		return
	}

	//f, _ := os.Stat("./test2.txt")
	//fmt.Println(f.ModTime())
}

/**
截取文件，设置为0后就是清空文件
syscall.truncate unix.truncate都是一样的
*/
func truncateDemo() {
	err := os.Truncate("./test2.txt", 0)
	if err != nil {
		log.ErrLog(err)
	}
	//syscall.Ftruncate() 文件描述符操作
}
