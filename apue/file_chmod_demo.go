package main

import "syscall"

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
	syscall.Chmod("./test2.txt", 0655) //这里必须加行0，因为这个值是八进制才行

}
