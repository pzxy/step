package leetcode

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	a, b := 2, 3
	target := 5
	if Add(a, b) == target {
		log.Fatal("Add Fail")
	}
}
