package main

//#cgo CFLAGS: -DPNG_DEBUG=1 -I./pcre/include
//#cgo LDFLAGS: -L./pcre/lib -lpcre2-8
//#define PCRE2_CODE_UNIT_WIDTH 8
//#include <stdint.h>
//#include <stdio.h>
//#include "./pcre/include/pcre2.h"
import "C"
import "unsafe"

func main() {
	errornumber := C.int(0)
	erroroffset := C.ulong(0)
	value := C.CString(`^[0-9]*$`)
	_ = C.pcre2_compile(
		C.PCRE2_SPTR(unsafe.Pointer(value)), /* the å */
		C.ulong(0),                          /* indicates pattern is zero-terminated */
		0,                                   /* default options */
		&errornumber,                        /* for error number */
		&erroroffset,                        /* for error offset */
		nil)                                 /* use default compile context */
}
