package main

//#cgo CFLAGS: -DPNG_DEBUG=1 -I./pcre/include
//#cgo LDFLAGS: -L./pcre/lib -lpcre2-8
//#define PCRE2_CODE_UNIT_WIDTH 8
//#include <stdio.h>
//#include "./pcre/include/pcre2.h"
/*
#include <stdio.h>

void printint(int v)  {
    printf("printint: %d\n", v);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	startMatchIdx := C.ulong(2)
	errNumber := C.int(0)
	errOffset := C.ulong(0)
	pattern := `\d{4}([^0-9^\s]{3,11})[^\s]`
	src := `a;jhgoqoghqoj0329 u0tyu10hg0h9Y0Y9827342482y(Y0y(G)_)lajf;lqjfgqhgpqjopjqa=)*(^!@#$%^&*())9999999`

	value := C.CString(pattern)
	re := C.pcre2_compile(
		C.PCRE2_SPTR(unsafe.Pointer(value)), /* the å */
		C.ulong(len(pattern)),               /* indicates pattern is zero-terminated */
		0,                                   /* default options */
		&errNumber,                          /* for error number */
		&errOffset,                          /* for error offset */
		nil)                                 /* use default compile context */
	if re == nil {
		return
	}
	subject := C.PCRE2_SPTR(unsafe.Pointer(C.CString(src)))
	matchData := C.pcre2_match_data_create_from_pattern(re, nil)
	rc := C.pcre2_match(
		re,                /* the compiled pattern */
		subject,           /* the subject string */
		C.ulong(len(src)), /* the length of the subject */
		startMatchIdx,     /* start at offset 0 in the subject */
		0,                 /* default options */
		matchData,         /* block for storing the result */
		nil)               /* use default match context */
	if rc < 0 {
		switch rc {
		case C.PCRE2_ERROR_NOMATCH:
			fmt.Println("No match")
			return
		default:
			fmt.Println("Matching error ", rc)
			return
		}
	}
	if rc == 0 {
		fmt.Println("ovector was not big enough for all the captured substrings")
		return
	}

	ovector := C.pcre2_get_ovector_pointer(matchData)
	matchCount := C.pcre2_get_ovector_count(matchData)
	fmt.Println("matchCount:", matchCount)
	fmt.Println("subject:", subject)
	fmt.Println("ovector:", ovector)
	fmt.Println("ovector:", unsafe.Pointer(ovector))
	fmt.Println("ovector:", uintptr(unsafe.Pointer(ovector)))
	// uintptr(unsafe.Pointer(ovector) 转为可以计算的指针，也是数组的起始地址。
	// unsafe.Sizeof(C.ulong(1))计算数组中单位元素占用内存大小。
	// 0*unsafe.Sizeof(C.ulong(1))) 数组下标为0的元素的占用内存大小。
	v1 := *(*C.ulong)(unsafe.Pointer(uintptr(unsafe.Pointer(ovector)) + 0*unsafe.Sizeof(C.ulong(1))))
	v2 := *(*C.ulong)(unsafe.Pointer(uintptr(unsafe.Pointer(ovector)) + 1*unsafe.Sizeof(C.ulong(1))))
	v3 := *(*C.ulong)(unsafe.Pointer(uintptr(unsafe.Pointer(ovector)) + 2*unsafe.Sizeof(C.ulong(1))))
	v4 := *(*C.ulong)(unsafe.Pointer(uintptr(unsafe.Pointer(ovector)) + 3*unsafe.Sizeof(C.ulong(1))))
	v5 := *(*C.ulong)(unsafe.Pointer(uintptr(unsafe.Pointer(ovector)) + 4*unsafe.Sizeof(C.ulong(1))))

	fmt.Println("v1: ", v1)
	fmt.Println("v2: ", v2)
	fmt.Println("v3: ", v3)
	fmt.Println("v4: ", v4)
	fmt.Println("v5: ", v5)
	fmt.Println(src[v1:v2])

	//fmt.Println(C.PCRE2_SIZE(ovector))
	//fmt.Println(uintptr(unsafe.Pointer(subject)) + uintptr(*ovector))
	//fmt.Println(matchData)
	//	for i := 0; i < int(rc); i++ {
	//		substring_start := subject + ovector[C.ulong(2*i)]
	//		substring_length := ovector[2*i+1] - ovector[2*i]
	//		fmt.Println(substring_length, substring_start)
	//	}
}
