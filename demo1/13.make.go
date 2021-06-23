package main

import "os"

func main() {
	os.RemoveAll("./tedge")
	os.MkdirAll("./tedge", 0766)
}
