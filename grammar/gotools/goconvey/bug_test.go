package goconvey

import (
	"fmt"
	"reflect"
	"testing"

	"bou.ke/monkey"
)

type math1 struct {
	a, b int
}

func (m *math1) Add() int {
	fmt.Println("add")
	return m.a + m.b
}

func AddStub(m *math1) int {
	return m.a - m.b
}

func TestAdd(t *testing.T) {
	monkey.PatchInstanceMethod(reflect.TypeOf(&math1{}), "Add", AddStub)

	m := &math1{a: 50, b: 50}
	fmt.Println(m.Add())
}
