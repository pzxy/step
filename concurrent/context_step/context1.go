package main

/**
context的用途：主要是取消goroutine的运行:
1. 上下文信息传递 （request-scoped），比如处理 http 请求、在请求处理链路上传递信息；
2. 控制子 goroutine 的运行；
3. 超时控制的方法调用；
4. 可以取消的方法调用。
*/
