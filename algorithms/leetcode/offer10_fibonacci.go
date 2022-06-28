package leetcode

//递归实现
func Fibo1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n > 1 {
		return Fibo1(n-1) + Fibo1(n-2)
	} else {
		return -1
	}
}

//迭代实现
func Fibo2(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	} else {
		a, b := 1, 1
		result := 0
		for i := 3; i <= n; i++ {
			result = a + b
			a, b = b, result
		}
		return result
	}
}

//利用闭包
func Fibo3(n int) int {
	if n < 0 {
		return -1
	}

	f := Fibonacci()
	result := 0
	for i := 0; i < n; i++ {
		result = f()
	}
	return result

}

func Fiv() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func Fibonacci4(n int) int {
	if n <= 0 {
		return -1
	}
	f := func() func() int {
		a, b := 0, 1
		return func() int {
			a, b = b, a+b
			return a
		}
	}()
	result := 0

	for i := 0; i < n; i++ {
		result = f()
	}
	return result

}
