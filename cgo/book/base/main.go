package main

/*
#include <stdio.h>

void printint(int v) {
    printf("printint: %d\n", v);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	v := 42
	C.printint(C.int(v))

	nums := []uint8{1, 2, 3, 4, 5, 6, 7, 8}
	val := &nums[0] // val is the equivalent of the *uint8 the Data function returns
	ptr := unsafe.Pointer(val)
	sixthVal := (*uint8)(unsafe.Pointer(uintptr(ptr) + 5*unsafe.Sizeof(*val)))
	fmt.Println("Sixth element:", *sixthVal)
}
