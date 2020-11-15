package main

/**
流模式
Stream这里我来介绍一种把 Channel 当作流式管道使用的方式，
也就是把 Channel 看作流（Stream），提供跳过几个元素，或者是只取其中的几个元素等方法。
*/

func asStream(done <-chan struct{}, values ...interface{}) <-chan interface{} {
	s := make(chan interface{}) //创建一个unbuffered的channel
	go func() {                 // 启动一个goroutine，往s中塞数据
		defer close(s)             // 退出时关闭chan
		for _, v := range values { // 遍历数组
			select {
			case <-done:
				return
			case s <- v: // 将数组元素塞入到chan中
			}
		}
	}()
	return s
}

/**
流创建好以后，该咋处理呢？下面我再给你介绍下实现流的方法。
takeN：只取流中的前 n 个数据；
takeFn：筛选流中的数据，只保留满足条件的数据；
takeWhile：只取前面满足条件的数据，一旦不满足条件，就不再取；
skipN：跳过流中前几个数据；skipFn：跳过满足条件的数据；
skipWhile：跳过前面满足条件的数据，一旦不满足条件，当前这个元素和以后的元素都会输出给 Channel 的 receiver。
*/

func takeN(done <-chan struct{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{}) // 创建输出流
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ { // 只读取前num个元素
			select {
			case <-done:
				return
			case takeStream <- <-valueStream: //从输入流中读取元素
			}
		}
	}()
	return takeStream
}
