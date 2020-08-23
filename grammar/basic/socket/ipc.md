
## IPC
ipc是一种进程间通信的概念,或者说规范,方法包括:
- socket(套接字)
- signal(系统信号)
- pipe(管道)
- file lock(文件锁)
- message queue(消息对列)
- semaphore(信号灯.信号量)

1) 其中socket套接字可以实现不同计算机的进程间的通信

2) go中与socket相应的实体都在net包中

3) 其他实现在os包中:
- os/signal(针对系统信号)
- os.Pipe(创建命名管道)
- os/exec(支持匿名管道)


### socket(套接字)
要点:
1) go中syscall包中有一个与socket系统调用相对应的函数
,二者方法签名是一样的.
2) syscall中的Socket函数和平台没关系,go底层做了不同操作系统间的适配.

3) net包中的很多实体会直接或间接的用到syscall.Socket函数,例如net.Dial

函数syscall.Socket:
```go
//syscall.Socket()
func Socket(domain,typ,proto int)(fd int,err error){
    //domain(通信域)
    //typ(数据类型)
    //proto 协议
}
```
domain
- syscall.AF_INET(IPv4域)
- syscall.AF_INET6(IPv6域)
- syscall.AF_UNIX(Unix域)

typ
- syscall.SOCK_DGRAM(数据报文,如UDP)
- syscall.SOCK_STREAM(数据流,如TCP)

proto
 - 设置为0则函数会根据前两个参数值自行选择
   - syscall.AF_INET和syscall.SOCK_DGRAM 则选UDP
   - syscall.AF_INET6和syscall.SOCK_STREAM则选TCP



### net.Dial

net.Dial是对套接字Socket的又一层封装,里面有两个参数network和address,都是string类型

network可选类型
- "tcp"：代表 TCP 协议，其基于的 IP 协议的版本本根据参数address的值自适应。                             
- "tcp4"：代表基于 IP 协议第四版的 TCP 协议。
- "tcp6"：代表基于 IP 协议第六版的 TCP 协议.
- "udp"：代表 UDP 协议，其基于的 IP 协议的版本根据参数address的值自适应。                         
- "udp4"：代表基于 IP 协议第四版的 UDP 协议。
- "udp6"：代表基于 IP 协议第六版的 UDP 协议。
- "unix"：代表 Unix 通信域下的一种内部 socke协议，以 SOCK_STREAM 为 socket 类型。
- "unixgram"：代表 Unix 通信域下的一种内部 socket 协议，以 SOCK_DGRAM 为 socket类型。
- "unixpacket"：代表 Unix 通信域下的一种内部socket 协议，以 SOCK_SEQPACKET 为 socket 类型。
             
          
          
### net.DialTimeout

代表函数位网络链接建立完成而等待的最长时间 ,
时间主要花费在
- 解析参数network和address的值
- 创建socket实例并建立网络连接

底层是通过net.Dialer函数实现,应该了解一下                                  