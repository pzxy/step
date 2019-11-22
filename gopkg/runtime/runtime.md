## runtime

[链接](https://www.jianshu.com/p/84bac7932394)
````
Package runtime 包含与 Go 的运行时系统交互的操作，
例如控制 goroutines 的函数。
它还包括反映包使用的低级类型信息; 

````


#### 环境变量


````
1 
GOMAXPROCS 变量限制可同时执行用户级 Go 代码的操作系统线程的数量。
代表 Go 代码的系统调用中可以阻止的线程数量没有限制; 
那些不计入 GOMAXPROCS 限制。
该软件包的 GOMAXPROCS 函数查询和更改限制。

2.1
 GOTRACEBACK 变量控制 Go 程序由于未发现的恐慌或意外的运行时情况而失败时产生的输出量。
默认情况下，失败打印当前 goroutine 的堆栈跟踪，在运行系统内部隐藏函数，
然后退出，退出代码
2.2 
如果没有当前的 goroutine 或者失败是失败，
则打印所有 goroutine 的堆栈跟踪内部的运行时间。
GOTRACEBACK = none 完全忽略 goroutine 堆栈轨迹。
GOTRACEBACK = single（默认）的行为如上所述。
GOTRACEBACK = 全部为所有用户创建的例程添加堆栈跟踪。
GOTRACEBACK =系统类似于“全部”，
但为运行时功能添加堆栈帧并显示运行时在内部创建的 goroutines 。
GOTRACEBACK =崩溃与“系统”类似，但是以操作系统特定的方式崩溃而不是退出。
例如，在 Unix 系统上，崩溃引发 SIGABRT 以触发核心转储。
由于历史原因，GOTRACEBACK 设置0,1和2分别是 none，
all 和 system 的同义词。
运行时/调试包的 SetTraceback 函数允许在运行时增加输出量，
但不能减少低于环境变量指定的量。

````

#### 常量

````
编译器是构建运行二进制文件的编译器工具链的名称。已知的工具链是：

1 纠错
gc      Also known as cmd/compile.
gccgo   The gccgo front end, part of the GCC compiler suite.

2
const Compiler = "gc"
GOARCH 是正在运行的程序的架构目标：386，amd64，arm，s390x 等等之一。

3
const GOARCH string = sys.GOARCH
GOOS 是运行程序的操作系统目标：达尔文，freebsd，linux 等等之一。

4 
MemProfileRate 控制内存配置文件中记录和报告的内存分配部分。
分析器旨在按分配的每个 MemProfileRate 字节对平均一个分配进行采样。
要在配置文件中包含每个分配的块，请将 MemProfileRate 设置为1。
要完全关闭分析，请将 MemProfileRate 设置为0。
处理内存配置文件的工具假定配置文件速率在程序的整个生命周期内保持不变,并等于当前值。
程序改变内存分析速率应该尽早执行一次，
尽可能早地执行程序（例如，在 main 的开始处）

5
Caller 返回调用栈的信息，0 从栈顶开始

6
gc 垃圾回收

7
gomaxprocs设置使用的最大逻辑cpu数

8
goroot 赶回goroot路径

9
goexit 终止调用他的goroutine

10
gosched 让出goroutine

11





````