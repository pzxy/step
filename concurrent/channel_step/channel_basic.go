package main

/**
qcount：代表 chan 中已经接收但还没被取走的元素的个数。内建函数 len 可以返回这个字段的值。

dataqsiz：队列的大小。chan 使用一个循环队列来存放元素，
循环队列很适合这种生产者 - 消费者的场景（我很好奇为什么这个字段省略 size 中的 e）。

buf：存放元素的循环队列的 buffer。

elemtype 和 elemsize：chan 中元素的类型和 size。
因为 chan 一旦声明，它的元素类型是固定的，即普通类型或者指针类型，所以元素大小也是固定的。

sendx：处理发送数据的指针在 buf 中的位置。一旦接收了新的数据，指针就会加上 elemsize，移向下一个位置。
buf 的总大小是 elemsize 的整数倍，而且 buf 是一个循环列表。

recvx：处理接收请求时的指针在 buf 中的位置。一旦取出数据，此指针会移动到下一个位置。

recvq：chan 是多生产者多消费者的模式，如果消费者因为没有数据可读而被阻塞了，就会被加入到 recvq 队列中。

sendq：如果生产者因为 buf 满了而阻塞，会被加入到 sendq 队列中。
*/
