package main

import (
	"fmt"
	"os/exec"
	"step/grammar/codeskill/log"
	"syscall"
)

func main() {
	waitDemo()
}

/**
c中wait是等待子进程结束，wait是阻塞的，等结束以后才返回。
后来又有了一个waitpid的函数，这个是函数是不阻塞的，如果没有子进程结束，就会立刻返回。
还可以穿进去一个pid，等在这个进程结束在返回。

可见，调用wait和waitpid不仅可以获得子进程的终止信息，还可以使父进程阻塞等待子进
程终止，起到进程间同步的作用。如果参数status不是空指针，则子进程的终止信息通过
这个参数传出，如果只是为了同步而不关心子进程的终止信息，可以将status参数指定为
NULL。


这几种条件判断正好适应以下几种情况：

1.等待子进程被跟踪，属于TASK_STOPPED,

2.等待子进程僵死，属于TASK_ZOMBIE，

3.参数为WNOHANG，正好符合我们wait4调用参数WNOHANG的意义，就是不等待子进程直接执行完毕。

4.进程收到其它信号，属于其它情况。

5.等待进程的PID不存在或者不是当前进程的子进程，代码写错了。
*/
//http://www.bubuko.com/infodetail-2752650.html

/**
1）参数 pid 为欲等待的子进程的识别码：

　　pid < -1 ；等待进程组 ID 为 pid 绝对值的进程组中的任何子进程；

　　pid = -1 ；等待任何子进程，此时 waitpid() 相当于 wait()。实际上，wait()就是 pid = -1、options = 0 的waitpid()， 且有：

　　pid = 0 ；等待进程组 ID 与当前进程相同的任何子进程（也就是等待同一个进程组中的任何子进程）；

　　pid > 0 ；等待任何子进程 ID 为 pid 的子进程，只要指定的子进程还没有结束，waitpid() 就会一直等下去。

2）参数 options 提供一些额外的选项来控制 waitpid()：

　　WNOHANG；如果没有任何已经结束了的子进程，则马上返回，不等待；

　　WUNTRACED；如果子进程进入暂停执行的情况，则马上返回，但结束状态不予理会；

　　也可以将这两个选项组合起来使用，使用 OR 操作。如果不想使用这两个选项，也可以直接把 options 设为0

3）waitpid() 的返回值，有三种：

	a）正常返回时，waitpid() 返回收集到的子进程的PID；

	b）如果设置了 WNOHANG，而调用 waitpid() 时，没有发现已退出的子进程可收集，则返回0；

	c）如果调用出错，则返回 -1，这时erron 会被设置为相应的值以指示错误所在。（当 pid 所指示的子进程不错在，或此进程存在，
     但不是调用进程的子进程， waitpid() 就会返回出错，这时 erron 被设置为 ECHILD

4）wait3() 和 wait4() 函数除了可以获得子进程状态信息外，还可以获得子进程的资源使用信息，这些信息是通过参数 rusage 得到的。
   而 wait3() 与 wait4() 之间的区别是，wait3() 等待所有进程，而 wait4() 可以根据 pid 的值选择要等待的子进程，
   参数 pid 的意义与 waitpid() 函数的一样。
*/gi

func waitDemo() {
	//开启一个子进程
	cmd := exec.Cmd{
		Path: "nc",
		Args: []string{"ls", "-l", "8888"},
		Dir:  "/usr/bin",
	}

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	fmt.Println("Start child process with pid:", cmd.Process.Pid)

	var waitStatus syscall.WaitStatus
	var rusage syscall.Rusage
	//0是阻塞，其他是不阻塞，不管有没有结束的或者暂停的子进程都立即返回
	wpid, err := syscall.Wait4(-1, &waitStatus, 0, &rusage)
	if err != nil {
		log.ErrLog(err)
		return
	}
	//kill pid
	fmt.Printf("wpid:%v,wait status:%v,rusage:%+v \t", wpid, waitStatus, rusage)
}
