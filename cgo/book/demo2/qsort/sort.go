package qsort

/*
#include <stdlib.h>

typedef int (*qsort_cmp_func_t)(const void* a, const void* b);
*/
import "C"
import "unsafe"

// CompareFunc 这里支所以重新定义一下 C.qsort_cmp_func_t 类型，是因为cgo编译后，变量名称会变化，所以要重新定义一下。
// 比如上面的这个类型，展开以后就会是这样的：_Ctype_qsort_cmp_func_t
// 所以这里重新定义一下
type CompareFunc C.qsort_cmp_func_t

// Sort 这样以来外部使用的时候就不会依赖c包中的函数了。
func Sort(base unsafe.Pointer, num, size int, cmp CompareFunc) {
	C.qsort(base, C.size_t(num), C.size_t(size), C.qsort_cmp_func_t(cmp))
}
