package main

/**
使用场景：
那这些特殊的场景是啥呢？官方的文档中指出，在以下两个场景中使用 sync.Map，会比使用 map+RWMutex 的方式，性能要好得多：
1. 只会增长的缓存系统中，一个 key 只写入一次而被读很多次；
2. 多个 goroutine 为不相交的键集读、写和重写键值对。
总结：就是说读多写少，多个goroutine的键不一样。

基本实现：
原则read不需要加锁，dirty需要加锁。
read->dirty：增加一个数时，dirty为nil，就会将read->dirty
dirty->read：如果 miss 数等于 dirty 长度，直接将对象给read。
1. 空间换时间。通过冗余的两个数据结构（只读的 read 字段、可写的 dirty），来减少加锁对性能的影响。
2. 对只读字段（read）的操作不需要加锁。
3. 读，删，更新：优先从 read 字段读取、更新、删除，因为对 read 字段的读取不需要锁。
4. dirty->read 动态调整。miss 次数多了之后，将 dirty 数据提升为 read，避免总是从 dirty 中加锁读取。double-checking。
5. 删：加锁之后先还要再检查 read 字段，确定真的不存在才操作 dirty 字段。延迟删除。
6. 删除一个键值只是打标记，只有在提升 dirty 字段为 read 字段的时候才清理删除的数据。
7，增加的元素(需要锁)放到dirty中，找的时候read中找不到时，dirty->read,miss 次数加1。

*/
