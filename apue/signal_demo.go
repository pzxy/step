package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/**
1) SIGHUP:当用户退出shell时，由该shell启动的所有进程将收到这个信号，默认动作为终止进程
2）SIGINT：当用户按下了<Ctrl+C>组合键时，用户终端向正在运行中的由该终端启动的程序发出此信号。默认动
作为终止里程。
52 第5章 信号
3）SIGQUIT：当用户按下<ctrl+\>组合键时产生该信号，用户终端向正在运行中的由该终端启动的程序发出些信
号。默认动作为终止进程。
4）SIGILL：CPU检测到某进程执行了非法指令。默认动作为终止进程并产生core文件
5）SIGTRAP：该信号由断点指令或其他 trap指令产生。默认动作为终止里程 并产生core文件。
6 ) SIGABRT:调用abort函数时产生该信号。默认动作为终止进程并产生core文件。
7）SIGBUS：非法访问内存地址，包括内存对齐出错，默认动作为终止进程并产生core文件。
8）SIGFPE：在发生致命的运算错误时发出。不仅包括浮点运算错误，还包括溢出及除数为0等所有的算法错误。默
认动作为终止进程并产生core文件。
9）SIGKILL：无条件终止进程。本信号不能被忽略，处理和阻塞。默认动作为终止进程。它向系统管理员提供了可
以杀死任何进程的方法。
10）SIGUSE1：用户定义 的信号。即程序员可以在程序中定义并使用该信号。默认动作为终止进程。
11）SIGSEGV：指示进程进行了无效内存访问。默认动作为终止进程并产生core文件。
12）SIGUSR2：这是另外一个用户自定义信号 ，程序员可以在程序中定义 并使用该信号。默认动作为终止进程。1
13）SIGPIPE：Broken pipe向一个没有读端的管道写数据。默认动作为终止进程。
14) SIGALRM:定时器超时，超时的时间 由系统调用alarm设置。默认动作为终止进程。
15）SIGTERM：程序结束信号，与SIGKILL不同的是，该信号可以被阻塞和终止。通常用来要示程序正常退出。执行
shell命令Kill时，缺省产生这个信号。默认动作为终止进程。
16）SIGCHLD：子进程结束时，父进程会收到这个信号。默认动作为忽略这个信号。
17）SIGCONT：停止进程的执行。信号不能被忽略，处理和阻塞。默认动作为终止进程。
18）SIGTTIN：后台进程读终端控制台。默认动作为暂停进程。
19）SIGTSTP：停止进程的运行。按下<ctrl+z>组合键时发出这个信号。默认动作为暂停进程。
21）SIGTTOU:该信号类似于SIGTTIN，在后台进程要向终端输出数据时发生。默认动作为暂停进程。
22）SIGURG：套接字上有紧急数据时，向当前正在运行的进程发出些信号，报告有紧急数据到达。如网络带外数据
到达，默认动作为忽略该信号。
23）SIGXFSZ：进程执行时间超过了分配给该进程的CPU时间 ，系统产生该信号并发送给该进程。默认动作为终止
进程。
24）SIGXFSZ：超过文件的最大长度设置。默认动作为终止进程。
25）SIGVTALRM：虚拟时钟超时时产生该信号。类似于SIGALRM，但是该信号只计算该进程占用CPU的使用时间。默
认动作为终止进程。
26）SGIPROF：类似于SIGVTALRM，它不公包括该进程占用CPU时间还包括执行系统调用时间。默认动作为终止进
程。
27）SIGWINCH：窗口变化大小时发出。默认动作为忽略该信号。
28）SIGIO：此信号向进程指示发出了一个异步IO事件。默认动作为忽略。
29）SIGPWR：关机。默认动作为终止进程。
30）SIGSYS：无效的系统调用。默认动作为终止进程并产生core文件。
31）SIGRTMIN～（64）SIGRTMAX：LINUX的实时信号，它们没有固定的含义（可以由用户自定义）。所有的实时信
号的默认动作都为终止进程。
*/
func main() {
	listenSignalDemo()
}

/**
捕获Ctrl + c信号:syscall.SIGINT
*/
func catchSignalDemo() {
	fmt.Printf("请输入 Ctrl + C \n")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt) //监听指定信号
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		for s := range sigChan {
			switch s {
			case syscall.SIGINT:
				fmt.Printf("捕获了Ctrl + C信号:%v \n", s)
				wg.Done()
			default:
				fmt.Printf("不处理此信号:%v \n", s)
			}
		}
	}()
	time.Sleep(time.Second)
	wg.Wait()
	fmt.Println("程序结束")
}

/**
忽视Ctrl + c信号,syscall.SIGINT
*/
func ignoreSignalDemo() {
	fmt.Printf("请输入 Ctrl + C,会忽视掉此信号 \n")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		for s := range sigChan {
			switch s {
			/*case syscall.SIGINT:
			fmt.Printf("捕获了Ctrl + C信号:%v \n", s)
			wg.Done()*/
			default:
				fmt.Printf("不处理此信号:%v，如要关闭，请直接杀死进程 \n", s.String())
			}
		}
	}()
	time.Sleep(time.Second)
	wg.Wait()
	fmt.Println("程序结束")
}

/**
all signal
*/

func listenSignalDemo() {
	ntf := make(chan int, 1)
	clean := func() {
		//fmt.Println("开始执行清理")
		//fmt.Println("结束清理")
		ntf <- 0
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c) //监听所有信号

	go func() {
		for s := range c {
			switch s {
			case syscall.SIGINT:
				fmt.Printf("退出请调节窗口大小 \n")
				//clean()
			case syscall.SIGFPE:
				//fmt.Printf("捕获了信号:%v \n", s.String())
				//clean()
			case syscall.SIGWINCH:
				//fmt.Printf("捕获了信号:%v \n", s.String())
				clean()
			default:
				fmt.Printf("不处理此信号:%v \n", s.String())
				//clean()
			}
		}
	}()

	<-ntf
	fmt.Println("程序结束")
}

/**
1.当进程处理SIGINT信号时，临时阻塞SIGQUIT信号。
2.进程间利用信号传参，来控制数据同步，如两个进程交叉报数。
*/

func practiseDemo() {

}
