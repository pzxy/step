package main

//extern int go_qsort_compare(void* a, void* b);
import "C"

import (
	"fmt"
	"step/cgo/book/demo2/qsort"
	"unsafe"
)

// 我们要使用Sort函数，还是要自己定义比较函数。
// 并且这个比较函数是用c来写的，那么如何才能用go写比较函数呢。参考demo3

//export go_qsort_compare
func go_qsort_compare(a, b unsafe.Pointer) C.int {
	pa, pb := (*C.int)(a), (*C.int)(b)
	return C.int(*pa - *pb)
}

func main() {
	values := []int32{42, 9, 101, 95, 27, 25}

	qsort.Sort(unsafe.Pointer(&values[0]),
		len(values), int(unsafe.Sizeof(values[0])),
		qsort.CompareFunc(C.go_qsort_compare),
	)
	fmt.Println(values)
}
