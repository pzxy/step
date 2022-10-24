package qsort

/*
#include <stdlib.h>

typedef int (*qsort_cmp_func_t)(const void* a, const void* b);
extern int _cgo_qsort_compare(void* a, void* b);
*/
import "C"
import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

// 包装这个比较函数。

//export _cgo_qsort_compare
func _cgo_qsort_compare(a, b unsafe.Pointer) C.int {
	var (
		// array memory is locked
		base     = uintptr(go_qsort_compare_info.base)
		elemsize = uintptr(go_qsort_compare_info.elemsize)
	)

	i := int((uintptr(a) - base) / elemsize)
	j := int((uintptr(b) - base) / elemsize)

	switch {
	case go_qsort_compare_info.less(i, j): // v[i] < v[j]
		return -1
	case go_qsort_compare_info.less(j, i): // v[i] > v[j]
		return +1
	default:
		return 0
	}
}

// 我们来实现一个和go中排序相同的排序方案
//func Slice(slice interface{}, less func(a, b int) bool)

// go_qsort_compare_info 切片首个元素地址，元素个数，元素大小，我们通过通过反射获取。
// 并且将这三个变量放到全局变量中，比较函数为less函数。
var go_qsort_compare_info struct {
	base     unsafe.Pointer
	elemnum  int
	elemsize int
	less     func(a, b int) bool
	sync.Mutex
}

func Slice(slice interface{}, less func(a, b int) bool) {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("qsort called with non-slice value of type %T", slice))
	}
	if sv.Len() == 0 {
		return
	}

	go_qsort_compare_info.Lock()
	defer go_qsort_compare_info.Unlock()

	defer func() {
		go_qsort_compare_info.base = nil
		go_qsort_compare_info.elemnum = 0
		go_qsort_compare_info.elemsize = 0
		go_qsort_compare_info.less = nil
	}()

	// baseMem = unsafe.Pointer(sv.Index(0).Addr().Pointer())
	// baseMem maybe moved, so must saved after call C.fn
	go_qsort_compare_info.base = unsafe.Pointer(sv.Index(0).Addr().Pointer())
	go_qsort_compare_info.elemnum = sv.Len()
	go_qsort_compare_info.elemsize = int(sv.Type().Elem().Size())
	go_qsort_compare_info.less = less

	C.qsort(
		go_qsort_compare_info.base,
		C.size_t(go_qsort_compare_info.elemnum),
		C.size_t(go_qsort_compare_info.elemsize),
		C.qsort_cmp_func_t(C._cgo_qsort_compare),
	)
}
