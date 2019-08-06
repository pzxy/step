package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// strings.Reader类型的值,
// 可以让我们很方便地读取一个字符串中的内容。
// 在读取的过程中,Reader值会保存已读取的字节的计数

func main() {

	// 省略若干代码。
	file, _ := ioutil.ReadFile("basic/json/re.json")
	reader1 := strings.NewReader(string(file))

	// 不读取后五个
	for reader1.Len() > 5 {
		b, err := reader1.ReadByte() // 读取1 byte
		fmt.Println(string(b), err, reader1.Len(), reader1.Size())
		// h <nil> 10 11
		// e <nil> 9 11
		// l <nil> 8 11
		// l <nil> 7 11
		// o <nil> 6 11
		//   <nil> 5 11
	}
	fmt.Println("....................")
	readingIndex := reader1.Size() - int64(reader1.Len()) // 计算出的已读计数。
	fmt.Println(readingIndex)

	offset2 := int64(17)
	expectedIndex := reader1.Size() - int64(reader1.Len()) + offset2
	fmt.Printf("Seek with offset %d and whence %d ...\n", offset2, io.SeekCurrent)
	readingIndex, _ = reader1.Seek(offset2, io.SeekCurrent)
	fmt.Printf("The reading index in reader: %d (returned by Seek)\n", readingIndex)
	fmt.Printf("The reading index in reader: %d (computed by me)\n", expectedIndex)

	s := "hello world"
	// 创建 Reader
	r := strings.NewReader(s)

	fmt.Println(r)        // &{hello world 0 -1}
	fmt.Println(r.Size()) // 11 获取字符串长度
	fmt.Println(r.Len())  // 11 获取未读取长度

	// 不读取最后五个字符
	for r.Len() > 5 {
		b, err := r.ReadByte() // 读取1 byte
		fmt.Println(string(b), err, r.Len(), r.Size())
		// h <nil> 10 11
		// e <nil> 9 11
		// l <nil> 8 11
		// l <nil> 7 11
		// o <nil> 6 11
		//   <nil> 5 11
	}

	// 读取还未被读取字符串中5字符的数据
	b_s := make([]byte, 1)
	n, err := r.Read(b_s)
	fmt.Println(string(b_s), n, err) // world 5 <nil>
	fmt.Println(r.Size())            // 11
	fmt.Println(r.Len())             // 0

}
