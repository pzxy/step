package goconvey

import "fmt"

func Sum(a int, b int) int {
	return a + b
}

type MathDemo struct {
	a int
	b int
}

func (m *MathDemo) Sub() int {
	fmt.Println("sub")
	a := m.a - m.b
	return a
}
