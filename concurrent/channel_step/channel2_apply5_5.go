package main

import "fmt"

/**
Map-Reduce模式(hadoop)
map-reduce 是一种处理数据的方式，最早是由 Google 公司研究提出的一种面向大规模数据处理的并行计算模型和方法，
开源的版本是 hadoop，前几年比较火。

不过，我要讲的并不是分布式的 map-reduce，而是单机单进程的 map-reduce 方法。

map-reduce 分为两个步骤，
第一步是映射（map），处理队列中的数据，
第二步是规约（reduce）(明明就是分解归纳的意思)，把列表中的每一个元素按照一定的处理方式处理成结果，放入到结果队列中。
就像做汉堡一样，map 就是单独处理每一种食材，reduce 就是从每一份食材中取一部分，做成一个汉堡。
*/

func mapChan(in <-chan interface{}, fn func(interface{}) interface{}) <-chan interface{} {
	out := make(chan interface{}) //创建一个输出chan
	if in == nil {                // 异常检查
		close(out)
		return out
	}

	go func() { // 启动一个goroutine,实现map的主要逻辑
		defer close(out)
		for v := range in { // 从输入chan读取数据，执行业务操作，也就是map操作
			out <- fn(v)
		}
	}()

	return out
}

func reduce(in <-chan interface{}, fn func(r, v interface{}) interface{}) interface{} {
	if in == nil { // 异常检查
		return nil
	}

	out := <-in         // 先读取第一个元素
	for v := range in { // 实现reduce的主要逻辑
		out = fn(out, v)
	}

	return out
}

// 生成一个数据流
func asStream2(done <-chan struct{}) <-chan interface{} {
	s := make(chan interface{})
	values := []int{1, 2, 3, 4, 5}
	go func() {
		defer close(s)             //todo 这种写法需要学习一下，创建的管道，推出的时候要关闭
		for _, v := range values { // 从数组生成
			fmt.Println(v)
			select {
			case <-done:
				return
			case s <- v: //todo 没有缓冲区的chan，放不进去的话，在chan自己逻辑的代码处理的地方会阻塞住。
			}
		}
	}()
	return s
}

func main() {
	in := asStream2(nil)

	// map操作: 乘以10
	mapFn := func(v interface{}) interface{} {
		return v.(int) * 10
	}

	// reduce操作: 对map的结果进行累加
	reduceFn := func(r, v interface{}) interface{} {
		return r.(int) + v.(int)
	}

	sum := reduce(mapChan(in, mapFn), reduceFn) //返回累加结果
	fmt.Println(sum)
}
