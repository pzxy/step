go相关的包

````
archive : 文档操作,解压,压缩
bufio : 读取的一些操作,Reader,Scanner等结构
builtin :基本类型 int bool等
bytes : 关于字节的操作,读取,写入等
cmd : 直行外部指令,比如执行shell的一些指令
compress:也是和压缩的一些相关,还有霍夫曼的一些东西
container:堆栈,列表,环等数据结构
context : 上下文的一些操作
cryptp :加密的一些东西,md5,sha等
database:数据库的操作
debug : 没用过
encoding: 里面有json包 可以操作 结构体和json的转换
errors :错误
expvar:基本类型Int String 等的对象和对象的一些方法
flag :接受命令行参数,windows下的.exe的文件的执行的一些参数
fmt :打印相关
go : 里面有很多包,写编译器和路径的一些东西,类似ast抽象语法树,等等
hash:哈希操作,还有crc这些
html:和html相关,没用过
image:图像的操作,可以画图,但是仅仅支持2d
index:可以让数组快速的索引
internal:关于cpu系统的一些
io: 包括ioutil等,输出输出相关,readAll等
log :日志
math:数学操作,比特操作,有rand包,.s文件很多
mime:实现了MIME规范,扩展了邮件类型,比如声音,图像二进制文件等(gif图像,au声音等)
net:网络相关
os:包括打文件,so.Open,os.Stat这些等等,用的比较多
path:匹配一些拼接一些路径
plugin:Package 插件实现 Go 插件的加载和符号解析。
一个插件是一个带有导出函数和变量的 Go 主包，这些函数和变量已经被构建：在1.8版本中开放插件（Plugin）的支持，这意味着现在能从Go中动态加载部分函数。
reflect:反射相关的包
regexp:正则
runtime : 运行时交互的的一些函数,比如垃圾回收,goroutine等
sort:排序
strconv :格式转换
strings :字符串的操作,对比这些
sync :同步
syscall:软件包系统调用包含一个到低级操作系统基元的接口。具体细节取决于底层系统，默认情况下，
 godoc 将显示当前系统的系统调用文档。如果您希望 godoc 显示另一个系统的系统调用文档，
 请将 $ GOOS 和 $ GOARCH 设置为所需的系统。
 例如，如果你想在 linux / amd64 上查看 freebsd / arm 的文档，
 可以把 $ GOOS 设置为 freebsd 以及把 $ GOARCH 设置为 arm。
 系统调用的主要用途在于为系统提供更加便携的界面的其他软件包，
 例如“os”，“time”和“net”。如果可以的话，使用这些软件包而不是这一个。
testing:测试相关
text:文件
time:时间操作相关,睡眠,时间戳等
unicode:编码相关,例如utf-8,utf-16
unsafe:软件包的不安全包括围绕 Go 程序类型安全的操作步骤。
导入不安全的程序包可能不可移植，并且不受 Go 1 兼容性准则的保护。它表示任意Go表达式的类型。
````
发现了新大陆[腾讯云](https://cloud.tencent.com/developer/doc/1101)
TestPickUpCarShiftPsTrayDstIsChargingRelatedNodeButRobotNotCharging4TestMoveAwayRobot