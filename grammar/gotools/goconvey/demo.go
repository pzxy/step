package goconvey

var Sum = func(a int, b int) int {
	return a + b
}

type MathDemo struct {
	a int
	b int
}

func (m *MathDemo) Sub() int {
	return m.a - m.b
}
