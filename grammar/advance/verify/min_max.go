package main

import (
	"fmt"
	"time"
)

/**
测试：选择分支
两个数组a[i]和b[i]
要求a[i]设置成a【ℹ️】，b[i]中较小的一个，b[i]设置成a[i],b[i]中较大的一个。
 */
//即使不使用，go run -gcflags='-l -N' min_max.go，优化后的也比前面的快的多，
//要多次测试，因为可能分支全部选择正确的，效率是差不多的

func main() {
	slia := demo2Init(1000)
	slib := demo2Init(1000)
	n := len(slib)
	demo2Normal(slia, slib, n)
	demo2MinMax(slia, slib, n)

}

func demo2Init(n uint32) []uint32 {
	sli := make([]uint32, 0)
	for i := uint32(1); i <= n; i++ {
		sli = append(sli, i)
	}
	return sli
}
func demo2Normal(a []uint32, b []uint32, n int) {
	t1 := time.Now()
	for i := 0; i < n; i++ {
		if a[i] > b[i] {
			a[i], b[i] = b[i], a[i]
		}
	}
	t2 := time.Now()
	fmt.Printf("normal branch predicate:time %v \n", t2.Sub(t1))

}

//使用条件传送的方法去选择分支,但是go里面没三木运算
func demo2MinMax(a []uint32, b []uint32, n int) {

	t1 := time.Now()
	for i := 0; i < n; i++ {
		//min := a[i] > b[i]? b[i]:a[i]
		//
		//max := a[i], b[i] = b[i], a[i]
		midA := a[i]
		midB := b[i]
		flags := midA > midB
//奇怪这里如果是将a的值赋值给a，那么会非常快，将b的值给a，和平时一样
		if flags {
			a[i] = midB
			b[i] = midA
		}

	}
	t2 := time.Now()
	fmt.Printf("optimizing branch prediction:time %v \n", t2.Sub(t1))

}
