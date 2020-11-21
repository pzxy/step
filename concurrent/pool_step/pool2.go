package main

/**
线程安全

pool实现原理：
每次垃圾回收的时候，Pool 会把 victim 中的对象移除，然后把 local 的数据给 victim，这样的话，local 就会被清空，
而 victim 就像一个垃圾分拣站，里面的东西可能会被当做垃圾丢弃了，但是里面有用的东西也可能被捡回来重新使用。
victim 中的元素如果被 Get 取走，那么这个元素就很幸运，因为它又“活”过来了。
但是，如果这个时候 Get 的并发不是很大，元素没有被 Get 取走，那么就会被移除掉，因为没有别人引用它的话，就会被垃圾回收掉。


所有当前主要的空闲可用的元素都存放在 local 字段中，请求元素时也是优先从 local 字段中查找可用的元素。
local 字段包含一个 poolLocalInternal 字段，并提供 CPU 缓存对齐，从而避免 false sharing。

poolLocalInternal 也包含两个字段：private 和 shared：
1. private，代表一个缓存的元素，而且只能由相应的一个 P 存取。因为一个 P 同时只能执行一个 goroutine，所以不会有并发的问题。
2. shared，可以由任意的 P 访问，但是只有本地的 P 才能 pushHead/popHead，其它 P 可以 popTail，相当于只有一个本地的 P 作为生产者（Producer），
多个 P 作为消费者（Consumer），它是使用一个 local-free 的 queue 列表实现的。


Get：
首先，从本地的 private 字段中获取可用元素，因为没有锁，获取元素的过程会非常快，如果没有获取到，就尝试从本地的 shared 获取一个，如果还没有，
会使用 getSlow 方法去其它的 shared 中“偷”一个。最后，如果没有获取到，就尝试使用 New 函数创建一个新的。
这里的重点是 getSlow 方法，我们来分析下。看名字也就知道了，它的耗时可能比较长。它首先要遍历所有的 local，尝试从它们的 shared 弹出一个元素。
如果还没找到一个，那么，就开始对 victim 下手了。在 vintim 中查询可用元素的逻辑还是一样的，先从对应的 victim 的 private 查找，如果查不到，
就再从其它 victim 的 shared 中查找。
总结Get查找步骤：
Get()
自己的local：private，share
getSlow()
所有的local：share
自己的victim：private
其他的victim：share
victim都没有，编辑为空，下次跳过
都没找到，如果Pool初始化的时候设置了New字段，则New一个，
否则会返回一个nil


Push：
优先设置本地 private，如果 private 字段已经有值了，那么就把此元素 push 到本地队列中（share）。

坑：
1. 内存泄漏
2. 内存浪费
*/
