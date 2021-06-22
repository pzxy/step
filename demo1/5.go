package main

//https://bou.ke/blog/monkey-patching-in-go/
//https://www.hopperapp.com/
//http://disasm.pro/
func a() int {
	return 1
}
func main() {
	f := a
	f()
}
