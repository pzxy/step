package leetcode

import (
	"fmt"
	"runtime"
	"strings"
)

func allPrimeSum(n int, a int, sum int) int {

	if n >= a {
		return sum
	} else if isPrime(n) {
		return allPrimeSum(n+1, a, sum+n)
	} else {
		return allPrimeSum(n+1, a, sum)
	}

}

func isPrime(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 || n == 2 {
		return true
	}

	for j := 2; j < n; j++ {

		if n%j == 0 {
			break
		}

		if n == j+1 {
			return true
		}

	}
	return false
}

func test2() {
	for i := 1; i <= 100; i++ {

		if isPrime(i) {
			fmt.Println(i)
		}
	}
}
func test1() {
	CallerFrameDemo()
	for i := 2; i < 100; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				break
			}
			if i == j+1 {
				fmt.Println(i)
			}
		}
	}
}
func CallerFrameDemo() {
	// Ask runtime2.Callers for up to 10 pcs, including runtime2.Callers itself.
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	if n == 0 {
		// No pcs available. Stop now.
		// This can happen if the first argument to runtime2.Callers is large.
		return
	}

	pc = pc[:n] // pass only valid pcs to runtime2.CallersFrames
	frames := runtime.CallersFrames(pc)

	// Loop to get frames.
	// A fixed number of pcs can expand to an indefinite number of Frames.
	for {
		frame, more := frames.Next()
		// To keep this example's output stable
		// even if there are changes in the testing package,
		// stop unwinding when we leave package runtime2.
		if !strings.Contains(frame.File, "runtime2/") {
			break
		}
		fmt.Printf("- more:%v | %s\n", more, frame.Function)
		if !more {
			break
		}
	}
}
