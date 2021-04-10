package main

//https://bou.ke/blog/monkey-patching-in-go/
//https://www.hopperapp.com/
//http://disasm.pro/
//func a() int {
//	return 1
//}
//func main() {
//	f := a
//	f()
//}

// go tool compile -N -l main.go
// go tool objdump main.o

func a() int { return 1 }
func b() int { return 2 }
func main() {

}

/**
48 BA
00 00 00 14
MOVD R0, 16(RSP)
1.go:11               0x5e1                   f9400bfa                MOVD 16(RSP), R26
1.go:11               0x5e5                   f9400340                MOVD (R26), R0
1.go:11               0x5e9                   d63f0000                CALL (R0)
*/
