package datastruct

func r1(n int) int {
	if n == 0 {
		return 1
	}
	return r1(n - 1)
}

func r2(n int, a int) int {
	if n == 0 {
		return a
	}

	return r2(n-1, a*n)
}
