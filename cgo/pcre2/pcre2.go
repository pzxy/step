package pcre2

//#cgo CFLAGS: -DPNG_DEBUG=1 -I./pcre/include
//#cgo LDFLAGS: -L./pcre/lib -lpcre2-8
//#include <stdio.h>
//#include "./pcre/include/pcre2.h"
import "C"
import (
	"fmt"
	"unsafe"
)

// Match matchIdx：起始的匹配位置，函数运行结束后会填充匹配到的最大位置。
func Match(pattern string, subject string, matchIdx *int) ([]string, error) {
	var err error
	errNumber := C.int(0)
	errOffset := C.ulong(0)
	value := C.CString(pattern)
	re := C.pcre2_compile(
		C.PCRE2_SPTR(unsafe.Pointer(value)), /* the å */
		C.ulong(len(pattern)),               /* indicates pattern is zero-terminated */
		0,                                   /* default options */
		&errNumber,                          /* for error number */
		&errOffset,                          /* for error offset */
		nil)                                 /* use default compile context */
	if re == nil {
		var buffer *C.uchar
		C.pcre2_get_error_message(errNumber, buffer, C.ulong(unsafe.Sizeof(buffer)))
		return nil, fmt.Errorf("PCRE2 compilation failed at offset %v: %v", errOffset, buffer)
	}
	subjectC := C.PCRE2_SPTR(unsafe.Pointer(C.CString(subject)))
	matchData := C.pcre2_match_data_create_from_pattern(re, nil)
	defer func() {
		C.pcre2_match_data_free(matchData) /* Release memory used for the match */
		C.pcre2_code_free(re)              /*   data and the compiled pattern. */
	}()

	// options https://pcre2project.github.io/pcre2/doc/html/pcre2_match.html
	rc := C.pcre2_match(
		re,                    /* the compiled pattern */
		subjectC,              /* the subject string */
		C.ulong(len(subject)), /* the length of the subject */
		C.ulong(*matchIdx),    /* start at offset 0 in the subject */
		0,                     /* default options */
		matchData,             /* block for storing the result */
		nil)                   /* use default match context */

	if rc < 0 {
		switch rc {
		case C.PCRE2_ERROR_NOMATCH:
			err = fmt.Errorf("no match")
		default:
			err = fmt.Errorf("matching error %v", rc)
		}
		return nil, err
	}
	if rc == 0 {
		err = fmt.Errorf("ovector was not big enough for all the captured substrings")
		return nil, err
	}

	var ret []string
	outputVector := C.pcre2_get_ovector_pointer(matchData)
	idx1, idx2 := -1, 0
	for i := 0; idx1 != idx2; i += 2 {
		idx1 = int(*(*C.ulong)(unsafe.Pointer(uintptr(unsafe.Pointer(outputVector)) + uintptr(i)*unsafe.Sizeof(C.ulong(1)))))
		idx2 = int(*(*C.ulong)(unsafe.Pointer(uintptr(unsafe.Pointer(outputVector)) + uintptr(i+1)*unsafe.Sizeof(C.ulong(1)))))
		if idx1 > idx2 {
			err = fmt.Errorf("from end to start the match was: %s", subject[idx2:idx1])
			return nil, err
		}
		if idx1 == idx2 {
			break
		}
		*matchIdx = max(idx2, *matchIdx)
		ret = append(ret, subject[idx1:idx2])
	}
	return ret, nil
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
