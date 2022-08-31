package main

import (
	"fmt"
	"reflect"
	"testing"

	"bou.ke/monkey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStub1(t *testing.T) {
	monkey.Patch(demo1, demo1Stub)
	fmt.Println(demo1().Success)
}

type demo2 struct {
	a, b int
}

func (m *demo2) Add() int {
	fmt.Println("----> add")
	return m.a + m.b
}

func Sub(m *demo2) int {
	fmt.Println("----> sub")
	return m.a - m.b
}

func TestStub2(t *testing.T) {
	monkey.PatchInstanceMethod(reflect.TypeOf(&demo2{}), "Add", Sub)
	m := &demo2{a: 50, b: 30}
	fmt.Println(m.Add())
}

func TestStub(t *testing.T) {
	monkey.Patch(demo1, demo1Stub)
	monkey.PatchInstanceMethod(reflect.TypeOf(&demo2{}), "Add", Sub)
	Convey("Stub", t, func() {
		Convey("stub function", func() {
			resp := demo1()
			So(resp.Status, ShouldEqual, "demo1Stub")
		})
		Convey("stub method", func() {
			m := &demo2{a: 50, b: 30}
			v := m.Add()
			So(v, ShouldEqual, 20)
		})
	})
}
