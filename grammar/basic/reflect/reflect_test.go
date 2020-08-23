package main

import (
	"reflect"
	"testing"
)

//测试原生结构体成员赋值
type data struct {
	Hp int
}

func BenchmarkNativeAssign(b *testing.B) {
	v := data{Hp: 2}
	//重置基准计时器
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v.Hp = 3
	}
}

//测试反射结构体成员赋值f.SetInt()
func BenchmarkReflectAssign(b *testing.B) {
	v := data{Hp: 2}
	vv := reflect.ValueOf(&v).Elem()
	f := vv.FieldByName("Hp")
	//重置基准计时器
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		f.SetInt(3)
	}
}

//结构体成员搜索并赋值对比 vv.FieldByName("Hp").
func BenchmarkReflectFindFieldAndAssign(b *testing.B) {
	v := data{Hp: 2}
	vv := reflect.ValueOf(&v).Elem()
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vv.FieldByName("Hp").SetInt(3)
	}
}

//原生调用函数
func foo(v int) {

}

func BenchmarkNaviteCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo(0)
	}
}

//反射调用函数
func BenchmarkReflectCall(b *testing.B) {
	v := reflect.ValueOf(foo)
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		//reflect.ValueOf(2)将2 构造为反射值对象,因为反射函数调用的参数必须全是反射值对象
		//再使用[]reflect.Value构造多个参数列表传给反射值对象的Call()方法进行调用
		v.Call([]reflect.Value{reflect.ValueOf(2)}) //Call内部好复杂 反正没看懂 :(
	}
}

/** 测试结果
[pzxy@pzxy-pc reflect]$ go test -v -bench=.
goos: linux
goarch: amd64
pkg: godemo/basic/reflect
BenchmarkNativeAssign-4                 2000000000               0.37 ns/op   原生赋值
BenchmarkReflectAssign-4                300000000                4.81 ns/op   反射结构体赋值
BenchmarkReflectFindFieldAndAssign-4    20000000               100 ns/op     反射结构体FieldByName("Hp")搜索后赋值
BenchmarkNaviteCall-4                   2000000000               0.37 ns/op  原生调用函数
BenchmarkReflectCall-4                  10000000               165 ns/op	反射调用函数 简直太慢了
PASS
ok      godemo/basic/reflect    7.431s

*/
//解决方案是 提前缓冲对象或者函数列表
