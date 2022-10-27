package pcre2

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

func Match(pattern string, subject string) ([]string, error) {
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

	// 循环匹配多次
	var ret []string
	var startMatchIdx, idx1, idx2 int
	var rc C.int
	for err == nil && startMatchIdx < len(subject) {
		// options https://pcre2project.github.io/pcre2/doc/html/pcre2_match.html
		//  PCRE2_ANCHORED          Match only at the first position
		//  PCRE2_COPY_MATCHED_SUBJECT
		//                          On success, make a private subject copy
		//  PCRE2_ENDANCHORED       Pattern can match only at end of subject
		//  PCRE2_NOTBOL            Subject string is not the beginning of a line
		//  PCRE2_NOTEOL            Subject string is not the end of a line
		//  PCRE2_NOTEMPTY          An empty string is not a valid match
		//  PCRE2_NOTEMPTY_ATSTART  An empty string at the start of the subject is not a valid match
		//  PCRE2_NO_JIT            Do not use JIT matching
		//  PCRE2_NO_UTF_CHECK      Do not check the subject for UTF validity (only relevant if PCRE2_UTF
		//                           was set at compile time)
		//  PCRE2_PARTIAL_HARD      Return PCRE2_ERROR_PARTIAL for a partial match even if there is a full match
		//  PCRE2_PARTIAL_SOFT      Return PCRE2_ERROR_PARTIAL for a partial match if no full matches are found
		rc = C.pcre2_match(
			re,                     /* the compiled pattern */
			subjectC,               /* the subject string */
			C.ulong(len(subject)),  /* the length of the subject */
			C.ulong(startMatchIdx), /* start at offset 0 in the subject */
			0,                      /* default options */
			matchData,              /* block for storing the result */
			nil)                    /* use default match context */

		if rc < 0 {
			switch rc {
			case C.PCRE2_ERROR_NOMATCH:
				err = fmt.Errorf("no match")
			default:
				err = fmt.Errorf("matching error %v", rc)
			}
			break
		}

		if rc == 0 {
			err = fmt.Errorf("ovector was not big enough for all the captured substrings")
			break
		}
		outputVector := C.pcre2_get_ovector_pointer(matchData)
		idx1 = int(*(*C.ulong)(unsafe.Pointer(uintptr(unsafe.Pointer(outputVector)) + 0*unsafe.Sizeof(C.ulong(1)))))
		idx2 = int(*(*C.ulong)(unsafe.Pointer(uintptr(unsafe.Pointer(outputVector)) + 1*unsafe.Sizeof(C.ulong(1)))))
		if idx1 > idx2 {
			err = fmt.Errorf("from end to start the match was: %s", subject[idx2:idx1])
			break
		}
		ret = append(ret, subject[idx1:idx2])
		if idx1 == idx2 {
			break
		}
		startMatchIdx = idx2
	}
	if len(ret) != 0 {
		return ret, nil
	}
	return ret, err
}
