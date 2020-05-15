package log

import (
	"fmt"
	"testing"
)

func TestCreateDir(t *testing.T) {
	if !CreateDir("logs") {
		fmt.Println("CreateDir failed!")
	}
	fmt.Println("CreateDir success!")
}
