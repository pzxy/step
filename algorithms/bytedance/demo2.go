package bytedance

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

func strConvDemo() {
	var builder strings.Builder
	for i := 0; i < 10; i++ {
		builder.WriteString(strconv.Itoa(i))
	}
	fmt.Println(builder.String())
	//fmt.Println(builder.WriteString())

	fmt.Println("*********")
	b := []byte("0123456789")
	fmt.Println(*(*int32)(unsafe.Pointer(&b)))
}

/**

func pointDemo() {
	a := uint32(1)
	b := int32(2)
	fmt.Println(&a, &b)
	c := &a
	c = &b            //类型不同，无法赋值
	c = (*uint32)(&b) //无效的，无法赋值
	fmt.Println(b)
}
*/

func pointDemo2() {
	a := uint32(1)
	b := int32(1)
	fmt.Println(&a, &b)
	c := &a
	c = (*uint32)(unsafe.Pointer(&b))
	fmt.Println(c)
}
