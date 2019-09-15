package main

import (
	"step/basic/testtwo"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := testtwo.Test11(2,3)
	if sum == 5{
		t.Log("测试通过")
	}
}

