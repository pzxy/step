package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"step/grammar/codeskill/log"
	"syscall"
)

func main() {
	pathconfDemo()
}

/**
更改工作路径
*/
func chdirGetcwdDemo() {
	getwdDemo()
	f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR, 0755)
	f.Close()

	err = os.Chdir("/")
	if err != nil {
		log.ErrLog(err)
		return
	}
	getwdDemo()
	//如果没切换工作目录的话，是不会报错的。
	err = syscall.Access("./test.txt", unix.W_OK) //
	if err != nil {
		log.ErrLog(err)
		return
	}
}

/**
获取工作目录
*/
func getwdDemo() {
	//打印工作路径，c中是getpwd，go里面是getwd
	path, _ := os.Getwd()
	fmt.Printf("path :%v \n", path)
}

/**
获取系统的资源限制，比如文件名字最大多长,目录路径最长是多少，io缓冲区是多大等
fpathconf 和 pathconf 两个
*/
///* configurable pathname variables */
//#define _PC_LINK_MAX             1
//#define _PC_MAX_CANON            2
//#define _PC_MAX_INPUT            3
//#define _PC_NAME_MAX             4
//#define _PC_PATH_MAX             5
//#define _PC_PIPE_BUF             6
//#define _PC_CHOWN_RESTRICTED     7
//#define _PC_NO_TRUNC             8
//#define _PC_VDISABLE             9
//
//#if !defined(_POSIX_C_SOURCE) || defined(_DARWIN_C_SOURCE)
//#define _PC_NAME_CHARS_MAX       10
//#define _PC_CASE_SENSITIVE               11
//#define _PC_CASE_PRESERVING              12
//#define _PC_EXTENDED_SECURITY_NP        13
//#define _PC_AUTH_OPAQUE_NP      14
//#endif
//
//#define _PC_2_SYMLINKS          15      /* Symlink supported in directory */
//#define _PC_ALLOC_SIZE_MIN      16      /* Minimum storage actually allocated */
//#define _PC_ASYNC_IO            17      /* Async I/O [AIO] supported? */
//#define _PC_FILESIZEBITS        18      /* # of bits to represent file size */
//#define _PC_PRIO_IO             19      /* Priority I/O [PIO] supported? */
//#define _PC_REC_INCR_XFER_SIZE  20      /* Recommended increment for next two */
//#define _PC_REC_MAX_XFER_SIZE   21      /* Recommended max file transfer size */
//#define _PC_REC_MIN_XFER_SIZE   22      /* Recommended min file transfer size */
//#define _PC_REC_XFER_ALIGN      23      /* Recommended buffer alignment */
//#define _PC_SYMLINK_MAX         24      /* Max # of bytes in symlink name */
//#define _PC_SYNC_IO             25      /* Sync I/O [SIO] supported? */
//#define _PC_XATTR_SIZE_BITS     26      /* # of bits to represent maximum xattr size */
//#define _PC_MIN_HOLE_SIZE       27      /* Recommended minimum hole size for sparse files */
func pathconfDemo() {
	//按理说应该是 _PC_NAME_MAX,但是这里面没有定义这个东西，这个函数第二个参数，不知道应该传进去什么
	//我打印了c中的常量后发现_PC_NAME_MAX值为5,但是go里面没定义这个。
	n, err := syscall.Pathconf("./test.txt", 1)
	if err != nil {
		log.ErrLog(err)
		return
	}

	fmt.Printf("file name max length:%v \n", n)
}
