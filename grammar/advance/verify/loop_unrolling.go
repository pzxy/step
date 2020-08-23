package main

import (
	"fmt"
	"time"
)

/**
验证循环展开
 */

/**
总结：
前提：基于本测试代码
当执行循环大于200～300的时候，甚至更高优化就会起作用
 */
func main() {

	fmt.Println("********* add ********")
	sli := demoInit(9999)
	demoNormal_add(sli)
	demoLoopUnrolling_add(sli)
	demoLoopUnrolling2_add(sli)
	fmt.Println("********* mul ********")
	sli2 := demoInit_mul(1000)
	demoNormal_mul(sli2)
	demoLoopUnrolling_mul(sli2)
	demoLoopUnrolling2_mul(sli2)
	demoLoopUnrolling3_mul(sli2)
}

func demoInit(n uint32) []uint32 {
	sli := make([]uint32, 0)
	for i := uint32(1); i <= n; i++ {
		sli = append(sli, i)
	}
	return sli
}

func demoInit_mul(n uint32) []uint32 {
	sli := make([]uint32, 0)
	for i := uint32(1); i < n; i++ {
		sli = append(sli, 1)
	}
	return sli
}
func demoNormal_add(sli []uint32) {
	var sum uint32
	t1 := time.Now()
	n := len(sli) //这里将这个移动到这里，确实能提高代码的效率
	for i := 0; i < n; i++ {
		sum += sli[i]
	}
	t2 := time.Now()
	fmt.Printf("normal add:time %v \n", t2.Sub(t1))
	fmt.Println(sum)

}

//2 *1 在机器上还是顺序的去执行了，还要等前面的加起来再加后面的值
// 但是编译器会做这些循环展开的工作，如果选择更高优化等级的话
func demoLoopUnrolling_add(sli []uint32) {
	var sum uint32
	var i int
	length := len(sli)
	n := len(sli) - 1
	t1 := time.Now()
	for i = 0; i < n; i += 2 {
		sum = sum + sli[i] + sli[i+1]
	}
	for ; i < length; i++ {
		sum = sum + sli[i]
	}
	t2 := time.Now()
	fmt.Printf("loop unrolling  add:  time %v \n", t2.Sub(t1))
	fmt.Println(sum)

}

//2*2 循环展开，并行操作,速度得到明显的提升

func demoLoopUnrolling2_add(sli []uint32) {
	var sum1 uint32
	var sum2 uint32
	var i int
	length := len(sli)
	n := len(sli) - 1
	t1 := time.Now()
	for i = 0; i < n; i += 2 {
		sum1 = sum1 + sli[i]
		sum2 = sum2 + sli[i+1]
	}
	for ; i < length; i++ {
		sum1 += sli[i]
	}
	t2 := time.Now()
	fmt.Printf("loop unrolling add & 并行 :  time %v \n", t2.Sub(t1))
	fmt.Println(sum1 + sum2)

}

func demoNormal_mul(sli []uint32) {
	sum := uint32(1)
	t1 := time.Now()
	n := len(sli) //这里将这个移动到这里，确实能提高代码的效率
	for i := 0; i < n; i++ {
		sum = sum * sli[i]
	}
	t2 := time.Now()
	fmt.Printf("normal mul:time %v \n", t2.Sub(t1))
	fmt.Println(sum)

}


//并行
func demoLoopUnrolling2_mul(sli []uint32) {
	sum1 := uint32(1)
	sum2 := uint32(1)
	var i int
	length := len(sli)
	n := len(sli) - 1
	t1 := time.Now()
	for i = 0; i < n; i += 2 {
		sum1 = sum1 * sli[i]
		sum2 = sum2 * sli[i+1]
	}
	for ; i < length; i++ {
		sum1 = sum1 * sli[i]
	}
	t2 := time.Now()
	fmt.Printf("loop unrolling  mul 并行:  time %v \n", t2.Sub(t1))
	fmt.Println(sum1 * sum2)

}
func demoLoopUnrolling_mul(sli []uint32) {
	sum := uint32(1)
	var i int
	length := len(sli)
	n := len(sli) - 1
	t1 := time.Now()
	for i = 0; i < n; i += 2 {
		sum = (sum * sli[i]) * sli[i+1]
	}
	for ; i < length; i++ {
		sum = sum * sli[i]
	}
	t2 := time.Now()
	fmt.Printf("loop unrolling  mul 向前结合:  time %v \n", t2.Sub(t1))
	fmt.Println(sum)

}
//k*a1 操作，性能与demoLoopUnrolling()相比也变快了
func demoLoopUnrolling3_mul(sli []uint32) {
	sum := uint32(1)
	var i int
	length := len(sli)
	n := len(sli) - 1
	t1 := time.Now()
	for i = 0; i < n; i += 2 {
		sum = sum * (sli[i] * sli[i+1])
	}
	for ; i < length; i++ {
		sum = sum * sli[i]
	}
	t2 := time.Now()
	fmt.Printf("loop unrolling  mul 后面的结合:  time %v \n", t2.Sub(t1))
	fmt.Println(sum)

}
