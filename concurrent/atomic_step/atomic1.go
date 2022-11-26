package main

const x int64 = 1 + 1<<33

func main() {
	var i = x
	_ = i
}

/**
GOARCH=386 go tools compile -N -l atomic1.go;GOARCH=386 go tools objdump -gnu atomic1.o
*/
