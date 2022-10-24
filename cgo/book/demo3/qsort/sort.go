package qsort

/*
#include <stdlib.h>

typedef int (*qsort_cmp_func_t)(const void* a, const void* b);
extern int _cgo_qsort_compare(void* a, void* b);
*/
import "C"
import (
	"sync"
	"unsafe"
)

// 这是Go自己的排序接口。
//
//func main() {
//    values := []int32{42, 9, 101, 95, 27, 25}
//
//    sort.Slice(values, func(i, j int) bool {
//        return values[i] < values[j]
//    })
//
//    fmt.Println(values)
//}
//

// 我们尝试将c中的qsort函数包装成 以下的go函数
//func Sort(base unsafe.Pointer, num, size int, cmp func(a, b unsafe.Pointer) int) {
//
//}

// go函数无法导出为c语言函数，因此无法将闭包函数传入c语言的qsort函数，
// 为此我们可以用go构造出一个可以导出为c语言的代理函数。
// 然后通过一个全局变量临时保存当前的闭包比较函数。
var go_qsort_compare_info struct {
	fn func(a, b unsafe.Pointer) int
	sync.Mutex
}

// _cgo_qsort_compare 是公用的sort函数，内部通过调用go_qsort_compare_info的fn函数。

//export _cgo_qsort_compare
func _cgo_qsort_compare(a, b unsafe.Pointer) C.int {
	return C.int(go_qsort_compare_info.fn(a, b))
}

func Sort(base unsafe.Pointer, num, size int, cmp func(a, b unsafe.Pointer) int) {
	// 每次都需要加锁，因为是全局变量，需要保证回调函数是当前操作的。
	go_qsort_compare_info.Lock()
	defer go_qsort_compare_info.Unlock()

	go_qsort_compare_info.fn = cmp
	// 这里其实是调用的 42行代码。
	C.qsort(base, C.size_t(num), C.size_t(size),
		C.qsort_cmp_func_t(C._cgo_qsort_compare),
	)
}
