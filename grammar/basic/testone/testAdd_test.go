package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Test11(2, 3)
	if sum == 5 {
		t.Log("测试通过")
	}
}
