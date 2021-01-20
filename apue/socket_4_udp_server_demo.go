package main

import (
	"bytes"
	"fmt"
	"step/grammar/codeskill/log"
	"syscall"
)

func main() {
	socketServer4Udp()
}

func socketServer4Udp() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		log.ErrLog(err)
		return
	}
	serverSockAddr := &syscall.SockaddrInet4{
		Port: 9627,
	}
	err = syscall.Bind(fd, serverSockAddr)
	if err != nil {
		log.ErrLog(err)
		return
	}
	fmt.Println("Accepting Connections...")
	buf := make([]byte, 80)
	for {
		n, clientSockAddr, err := syscall.Recvfrom(fd, buf, 0) //flags 可以用 ｜ 拼接
		if err != nil {
			log.ErrLog(err)
			return
		}
		fmt.Printf("server receive data : %s \n", buf[:n])
		fmt.Printf("client info  : %+v \n", clientSockAddr.(*syscall.SockaddrInet4))
		err = syscall.Sendto(fd, bytes.ToTitle(buf[:n]), 0, clientSockAddr)
		if err != nil {
			log.ErrLog(err)
			return
		}
	}

	syscall.Close(fd)
}

/**
flags：调用操作方式。是以下一个或者多个标志的组合体，可通过“ | ”操作符连在一起:
MSG_DONTWAIT：操作不会被阻塞。
MSG_ERRQUEUE： 指示应该从套接字的错误队列上接收错误值，依据不同的协议，错误值以某种辅佐性消息的方式传递进来，使用者应该提供足够大的缓冲区。
导致错误的原封包通过msg_iovec作为一般的数据来传递。导致错误的数据报原目标地址作为msg_name被提供。错误以sock_extended_err结构形态被使用，
定义如下
#define SO_EE_ORIGIN_NONE 0
#define SO_EE_ORIGIN_LOCAL 1
#define SO_EE_ORIGIN_ICMP 2
#define SO_EE_ORIGIN_ICMP6 3
struct sock_extended_err
{
u_int32_t ee_errno;
u_int8_t ee_origin;
u_int8_t ee_type;
u_int8_t ee_code;
u_int8_t ee_pad;
u_int32_t ee_info;
u_int32_t ee_data;
};
MSG_PEEK：指示数据接收后，在接收队列中保留原数据，不将其删除，随后的读操作还可以接收相同的数据。
MSG_TRUNC：返回封包的实际长度，即使它比所提供的缓冲区更长， 只对packet套接字有效。
MSG_WAITALL：要求阻塞操作，直到请求得到完整的满足。然而，如果捕捉到信号，错误或者连接断开发生，或者下次被接收的数据类型不同，仍会返回少于请求量的数据。
MSG_EOR：指示记录的结束，返回的数据完成一个记录。
MSG_TRUNC：指明数据报尾部数据已被丢弃，因为它比所提供的缓冲区需要更多的空间。
MSG_CTRUNC：指明由于缓冲区空间不足，一些控制数据已被丢弃。
MSG_OOB：指示接收到out-of-band数据(即需要优先处理的数据)。
MSG_ERRQUEUE：指示除了来自套接字错误队列的错误外，没有接收到其它数据。
*/
