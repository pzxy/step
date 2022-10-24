package main

import (
	"fmt"
	"step/cgo/book/demo3/qsort"
	"unsafe"
)

func main() {
	values := []int32{42, 9, 101, 95, 27, 25}
	// 现在已经不用导入C包了，但是还是需要导入不安全unsafe包。 解决参考demo4
	qsort.Sort(unsafe.Pointer(&values[0]), len(values), int(unsafe.Sizeof(values[0])),
		func(a, b unsafe.Pointer) int {
			pa, pb := (*int32)(a), (*int32)(b)
			return int(*pa - *pb)
		},
	)

	fmt.Println(values)
}
