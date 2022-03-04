package goconvey

import (
	"bou.ke/monkey"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestDemo1(t *testing.T) {
	Convey("TestSum", t, func() {
		Convey("TestSum", func() {
			result := Sum(1, 2)
			resultP := result == 3
			So(resultP, ShouldBeTrue)
		})
		Convey("TestSub", func() {
			m := &MathDemo{10, 2}
			result := m.Sub()
			resultP := result == 8
			So(resultP, ShouldBeTrue)
		})
	})
}

//打桩 go get github.com/prashantv/gostub

//使用打桩方法测试
func TestDemo2(t *testing.T) {
	md := &MathDemo{}
	//结构体的方法打桩
	monkey.PatchInstanceMethod(reflect.TypeOf(md), "Sub", SubStub)
	//简单方法打桩
	//monkey.Patch(Sum, SumStub)

	//Convey("TestSumStub", t, func() {
	//	Convey("TestSumStub", func() {
	// 场景：地图中找不到该点

	m := &MathDemo{10, 2}

	fmt.Println(m.Sub())
	//p := result == 6
	//So(p, ShouldBeTrue)
	//})
	//Convey("TestSubStub", func() {
	//	// 场景：地图中找不到该点
	//	m := &MathDemo{10, 2}
	//	result := SubStub(m)
	//	p := result == 5
	//	So(p, ShouldBeTrue)
	//})

}
