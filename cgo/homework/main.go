package main

import "C"
import "unsafe"

//#cgo CFLAGS: -DPNG_DEBUG=1 -I./pcre/include
//#cgo LDFLAGS: -L/usr/local/lib -lpng

//#define PCRE2_CODE_UNIT_WIDTH 8
//#include <stdint.h>
//#include <stdio.h>
//#include "./pcre/include/pcre2.h"
import "C"

func main() {
	errornumber := 0
	erroroffset := 0
	values := []rune(`^[0-9]*$`)
	_ = C.pcre2_compile(
		C._Ctype_uchar(unsafe.Pointer(&values[0])), /* the å */
		C.PCRE2_ZERO_TERMINATED,                    /* indicates pattern is zero-terminated */
		0,                                          /* default options */
		&errornumber,                               /* for error number */
		&erroroffset,                               /* for error offset */
		C.NULL)                                     /* use default compile context */
}
