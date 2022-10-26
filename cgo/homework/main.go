package main

//#cgo CFLAGS: -DPNG_DEBUG=1 -I./pcre/include
//#cgo LDFLAGS: -L./pcre/lib -lpcre2-8
//#define PCRE2_CODE_UNIT_WIDTH 8
//#include <stdio.h>
//#include "./pcre/include/pcre2.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	errNumber := C.int(0)
	errOffset := C.ulong(0)
	pattern := `^[0-9]*`
	value := C.CString(pattern)
	re, err := C.pcre2_compile(
		C.PCRE2_SPTR(unsafe.Pointer(value)), /* the å */
		C.ulong(len(pattern)),               /* indicates pattern is zero-terminated */
		0,                                   /* default options */
		&errNumber,                          /* for error number */
		&errOffset,                          /* for error offset */
		nil)                                 /* use default compile context */
	if err != nil {
		return
	}
	if re == nil {
		return
	}
	src := `1234567890abc`
	subject := C.PCRE2_SPTR(unsafe.Pointer(C.CString(src)))
	matchData := C.pcre2_match_data_create_from_pattern(re, nil)
	rc, err := C.pcre2_match(
		re,                /* the compiled pattern */
		subject,           /* the subject string */
		C.ulong(len(src)), /* the length of the subject */
		0,                 /* start at offset 0 in the subject */
		0,                 /* default options */
		matchData,         /* block for storing the result */
		nil)               /* use default match context */
	if err != nil {
		return
	}
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
	fmt.Println(ovector)
	fmt.Println(C.PCRE2_SPTR(unsafe.Pointer(ovector)))
	fmt.Println(matchData)
	//	for i := 0; i < int(rc); i++ {
	//		substring_start := subject + ovector[C.ulong(2*i)]
	//		substring_length := ovector[2*i+1] - ovector[2*i]
	//		fmt.Println(substring_length, substring_start)
	//	}
}
