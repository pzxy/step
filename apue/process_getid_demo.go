package main

import (
	"fmt"
	"os"
	"step/grammar/codeskill/log"
)

func main() {
	getGidDemo()
}

func getPidDemo() {
	fmt.Printf("pid:%v \n", os.Getpid())
	fmt.Printf("ppid:%v \n", os.Getppid())
}

func getUidDemo() {
	f, err := os.OpenFile("./getuid", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.ErrLog(err)
		return
	}
	fmt.Printf("uid:%v \n", os.Getuid())  //执行这个文件的用户id
	fmt.Printf("eid:%v \n", os.Geteuid()) //实际的用户id
	f.Close()

	//如果将这段代码编译以后放到跟目录下的话，那么文件的所有者是root，
	// 我们再用pzxy用户来执行文件，可以执行，但是因为是755，所以创建不了文件。
	//这时我们将编译文件的权限改为04755就可以执行了。效果为-rwsr-xr-x,这里的s代表设置用户ID。
	//可以看到显示出来的uid为501，eid为0，也就是root用户。执行的时候设置用户id为文件的所有者。
	// /usr/bin/passwd   -rwsr-xr-x  1 root  wheel  45232 Dec  5  2019 /usr/bin/passwd
	//我们通过passwd改密码的时候，虽然是普通用户但是能改root文件所有者的文件，那是因为passwd设置了 s位

}

func getGidDemo() {
	fmt.Printf("gid:%v \n", os.Getgid())   //执行这个文件的用户组id
	fmt.Printf("egid:%v \n", os.Getegid()) //实际的用户组id

	// 如果要设置组用户id，那么是06777，6也就是4，2
	// 如果是07777的话，那么显示出来是这样的，-rwsrwsrwt,最后的t代表粘住位。
	// 粘住位很少用了
	// 如果在一个执行文件设置了该位，则执行该文件且进程结束后，系统会把该进程的正文部分放置磁盘的交换区中，
	// 在交换区中文件是连续存放的，不像非交换区一样，一个文件的内容分散在磁盘的几个块中。
	// 所以在加载该执行文件时就可以加快速度启动，直接从交换区中把进程的正文部分取至内存中运行。
}
