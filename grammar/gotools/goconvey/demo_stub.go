package goconvey

//Sum的打桩测试方法
func SumStub(a int, b int) int {
	return a + b + 3
}

//Sub的打桩测试方法
func SubStub(m *MathDemo) int {
	return m.a - m.b - 3
}
