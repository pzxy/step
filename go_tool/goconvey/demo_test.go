package goconvey

import (
	"bou.ke/monkey"
	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestDemo1(t *testing.T) {
	//多个断言
	Convey("TestSum", t, func() {
		Convey("TestSum", func() {
			// 场景：地图中找不到该点
			result := Sum(1, 2)
			resultP := result == 3
			So(resultP, ShouldBeTrue)
		})
		Convey("TestSub", func() {
			// 场景：地图中找不到该点
			m := &MathDemo{10, 2}
			result := m.Sub()
			resultP := result == 8
			So(resultP, ShouldBeTrue)
		})
	})
}

//使用打桩方法测试
func TestDemo2(t *testing.T) {
	md := &MathDemo{}
	//结构体的方法打桩
	monkey.PatchInstanceMethod(reflect.TypeOf(md), "Sub", SubStub)
	//简单方法打桩
	gostub.Stub(&Sum, SumStub)

	Convey("TestSumStub", t, func() {
		Convey("TestSumStub", func() {
			// 场景：地图中找不到该点

			result := Sum(1, 2)
			resultP := result == 6
			So(resultP, ShouldBeTrue)
		})
		Convey("TestSubStub", func() {
			// 场景：地图中找不到该点
			m := &MathDemo{10, 2}
			result := SubStub(m)
			resultP := result == 5
			So(resultP, ShouldBeTrue)
		})
	})
}
