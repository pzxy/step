package main

import (
	"testing"
	"unicode/utf8"
)

// FuzzReverse 模糊测试
// 使用命令 go test -fuzz=Fuzz。来随机输入值
// go test -fuzz=Fuzz测试，几秒钟后，停止用 模糊测试ctrl-C。
// go test -fuzz=Fuzz -fuzztime 30s如果没有发现失败，它会在退出前 fuzz 30 秒。
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	// 除了返回err停止外，还可以调用t.Skip()以停止执行该模糊输入
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
