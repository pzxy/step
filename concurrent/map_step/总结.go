package main

/**
map：并发读写风险：
加锁map：性能不行：
分片锁map：好用
sync.Map：特定场景，读多写少，不怎么增加元素。或者多个goroutine操作不同的key。
*/
