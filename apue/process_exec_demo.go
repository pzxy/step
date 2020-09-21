package main

import (
	"fmt"
	"os"
	"os/exec"
	"step/grammar/codeskill/log"
	"syscall"
	"time"
)

func main() {
	execDemo2()
}

/**
一般fork和exec结合使用
exec并不会开一个子进程，只是执行完以后把这个进程关闭了，所以后面的不会打印了
*/
func execDemo() {
	fmt.Println("开始")
	argv := []string{"ls", "-l"}
	err := syscall.Exec("/bin/ls", argv, nil) //最后一个参数可以指定环境变量
	if err != nil {
		log.ErrLog(err)
	}
	fmt.Println("按理说这里不会打印没因为exec会替换栈中的这个方法")

}

/**
syscall.Exec需要三个参数：

第一个参数是可执行文件的路径，注意不会自动从PATH下面去搜索，所以：
1.1 要么是显式的指定全路径：/path/to/executable
1.2 要么是显式的指定相对路径: ./relpath/to/executable
1.3 要么通过exec.LookPath从PATH里面搜索出来，如本例子。
第二个参数是参数列表
2.1 注意args[0]是可执行程序名，这个内容会显示在ps -ef的输出中。用户可以改这个值，例如明明执行的是/usr/bin/sleep的可执行程序，但是这里可以改成任意字符串，例如ls，这样用户在ps -ef查看到的就是ls的命令在运行，而不是sleep命令，混淆用户。
2.2 后面是正常的参数。
第三个参数是环境变量
3.1 如果没有传，那么不会自动继承caller的环境变量的。
*/
func execDemo2() {
	fmt.Println(os.Getpid())
	fmt.Println(os.Getppid())
	fmt.Println(os.Getuid())
	fmt.Println(os.Getegid())
	fmt.Println("开始")
	path, err := exec.LookPath("ls") //path:/bin/ls,err=nil
	if err != nil {
		log.ErrLog(err)
		return
	}
	argv := []string{"ls", "-l"}
	envv := os.Environ()
	//这里随便输入路径的话会显示，权限不足
	err = syscall.Exec(path, argv, envv) //最后一个参数可以指定环境变量,可以不填,但是第一个命令的路径是必须填的
	fmt.Println("按理说这里不会打印没因为exec会替换栈中的这个方法")
	if err != nil {
		log.ErrLog(err)
		return
	}
	fmt.Println("按理说这里不会打印没因为exec会替换栈中的这个方法")

}

/**
执行命令但是不开子进程的方法
*/

func execDemo3() {
	file, err := os.OpenFile("./cmd.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {
		log.ErrLog(err)
		return
	}
	//cmd := exec.Command("echo", "hello")
	//-c 表示输出
	cmd := exec.Command("bash", "-c", fmt.Sprintf("cd %v && ls -l", "/"))
	cmd.Stdout = file
	cmd.Stderr = file
	if err := cmd.Run(); err != nil {
		log.ErrLog(err)
		return
	}
	time.Sleep(time.Second)
	file.Close()

	fmt.Println("按理说这里会打印,,,,因为没开启新进程替换栈中的这个方法")

}
