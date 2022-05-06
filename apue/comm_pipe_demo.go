package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
)

//https://www.jianshu.com/p/d82f8187c02c
func main() {
	UnixPipe()
}

/**

该包包含了一些操作系统底层的接口，实现细节依赖于底层的系统，默认下，
godoc只会展示当前系统的syscall文档。如果你想展示其他系统的syscall文档，需设置$GOOS和$GOARCH到目标系统。
syscall主要的使用是提供当前系统的接口给其他代码包如os，time，net。Go推荐使用那些包而不是这个syscall包。

*/
/**

os.Pipe() vs io.Pipe()
os.Pipe()基于syscall.Pipe()实现，比较底层，实现了基于操作系统级别的管道，而io.Pipe()基于channel和互斥锁实现数据通信
os.Pipe()不保证原子操作，io.Pipe是并发安全的
io.Pipe()因为使用了无缓冲channel， 读写是阻塞的

*/

//将命令ps aux 的输出管道连接到grep pipe的输入管道
//并把输出管道里的数据全部写到输入管道里

//说实话这里和linux中的pipe还是有区别的，pipe是一个管道有读端和写端。这里是两个管道。当然这里的管道也有输入和输出。
//这里以后还需要精炼一下，自己还需要学习的东西还有很多。

//简单介绍下Go实现执行外部命令的细节，exec.Command函数会调用exec的LookPath函数，
// 会去OS的PATH路径下查找，是否有对应的可执行程序文件，比如这里是echo，会依次访问path的路径，
// 看有没有echo这个文件，在我的Mac上，最终找到的路径为/bin/echo。
// 在exec.Start函数内部会调用exec.StartProcess来创建一个新的进程，在运行时执行fork系统调用

//上面的go程序使用了cmd.StdoutPipe()方法，
// 在StdoutPipe内部使用了下文要说的os.Pipe()函数，
// 而在这个函数内部调用了syscall.Pipe()，这个函数最终调用了pipe系统调用。

func UnixPipe() { //匿名管道

	//go里面操作命令行的正确姿势
	//	r, w := io.Pipe()
	//	cmd := exec.Command(nm, path)
	//	cmd.Stdout = w
	//	cmd.Stderr = os.Stderr

	fmt.Println("Run command `ps aux | grep apipe`: ")
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "apipe")
	stdout1, err := cmd1.StdoutPipe() //cmd1上建立一个输出管道，为*io.Reader类型
	if err != nil {
		fmt.Printf("Error: Can not obtain the stdout pipe for command: %s", err)
		return
	}
	if err := cmd1.Start(); err != nil {
		fmt.Printf("Error: The command can not running: %s\n", err)
		return
	}
	outputBuf1 := bufio.NewReader(stdout1) //避免数据过多带来的困扰，使用带缓冲的读取器来获取输出管道中的数据
	stdin2, err := cmd2.StdinPipe()        //cmd2上建立一个输入管道
	if err != nil {
		fmt.Printf("Error: Can not obtain the stdin pipe for command: %s\n", err)
		return
	}

	var outputBuf2 bytes.Buffer //获取cmd2的输出数据的字节缓冲器
	cmd2.Stdout = &outputBuf2   //将缓冲器赋值给cmd2的输出字段，这样cmd2的所有输出内容就会被写入到缓冲器中
	if err := cmd2.Start(); err != nil {
		fmt.Printf("Error: The command can not be startup: %s\n", err)
		return
	}

	outputBuf1.WriteTo(stdin2) //将缓冲读取器里的输出管道数据写入输入管道里
	err = stdin2.Close()       //关闭cmd2的输入管道
	if err != nil {
		fmt.Printf("Error: Can not close the stdio pipe: %s\n", err)
		return
	}
	if err := cmd2.Wait(); err != nil { //为了获取cmd2的所有输出内容，调用Wait()方法一直阻塞到其所属所有命令执行完
		fmt.Printf("Error: Can not wait for the command: %s\n", err)
		return
	}
	fmt.Printf("%s\n", outputBuf2.Bytes()) //输出执行结果
}

/**
系统调用 pipe
*/
