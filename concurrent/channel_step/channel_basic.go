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

/**
make channel初始化：
会调用makechan()，然后
1. chan的size或者元素的size是0，不必创建buf
2. 元素不是指针，分配一块连续的内存给hchan数据结构和buf
3. 元素包含指针，那么单独分配buf
最后记录下元素大小，类型，容量。

send:
编译的时候会将调用chansend1函数。
大致处理如下：
1. 如果 chan 是 nil 的话，就把调用者 goroutine park（阻塞休眠）， 调用者就永远被阻塞住了，所以，第 11 行是不可能执行到的代码。
2. 当你往一个已经满了的 chan 实例发送数据时，并且想不阻塞当前调用，那么这里的逻辑是直接返回。
3. 如果 chan 已经被 close 了，再往里面发送数据的话会 panic。
4. 如果等待队列中有等待的 receiver，那么这段代码就把它从队列中弹出，然后直接把数据交给它（通过 memmove(dst, src, t.size)），而不需要放入到 buf 中，速度可以更快一些。
5. 没有 receiver，需要把数据放入到 buf 中，放入之后，就成功返回了。
6. 如果 buf 满了，发送者的 goroutine 就会加入到发送者的等待队列中，直到被唤醒。这个时候，数据或者被取走了，或者 chan 被 close 了。

recv:
Go 会把代码转换成 chanrecv1 函数，如果要返回两个返回值，会转换成 chanrecv2，chanrecv1 函数和 chanrecv2 会调用 chanrecv。
chanrecv1 和 chanrecv2 传入的 block 参数的值是 true，都是阻塞方式，所以我们分析 chanrecv 的实现的时候，不考虑 block=false 的情况。
大致处理如下：
1. chan 为 nil 的情况。和 send 一样，从 nil chan 中接收（读取、获取）数据时，调用者会被永远阻塞。
2. chanrecv1 和 chanrecv2 传入的 block 参数的值是 true，都是阻塞方式，所以我们分析 chanrecv 的实现的时候，不考虑 block=false 的情况。
3. chan 已经被 close 的情况。如果 chan 已经被 close 了，并且队列中没有缓存的元素，那么返回 true、false。
4. 处理 buf 满的情况。这个时候，如果是 unbuffer 的 chan，就直接将 sender 的数据复制给 receiver，否则就从队列头部读取一个值，并把这个 sender 的值加入到队列尾部。
5. 处理没有等待的 sender 的情况。这个是和 chansend 共用一把大锁，所以不会有并发的问题。如果 buf 有元素，就取出一个元素给 receiver。
6. 处理 buf 中没有元素的情况。如果没有元素，那么当前的 receiver 就会被阻塞，直到它从 sender 中接收了数据，或者是 chan 被 close，才返回。


close:
通过 close 函数，可以把 chan 关闭，编译器会替换成 closechan 方法的调用。
逻辑如下：
1. 如果 chan 为 nil，close 会 panic；
2. 如果 chan 已经 closed，再次 close 也会 panic。
3. 如果 chan 不为 nil，chan 也没有 closed，就把等待队列中的 sender（writer）和 receiver（reader）从队列中全部移除并唤醒。
*/

/**
并发原语的选择：
1. 共享资源的并发访问使用传统并发原语；
2. 复杂的任务编排和消息传递使用 Channel；
3. 消息通知机制使用 Channel，除非只想 signal 一个 goroutine，才使用 Cond；
4. 简单等待所有任务的完成用 WaitGroup，也有 Channel 的推崇者用 Channel，都可以；
5. 需要和 Select 语句结合，使用 Channel；
6. 需要和超时配合时，使用 Channel 和 Context。
*/
