package tmp

import (
	"fmt"
	"testing"
)

func TestRangeNilChannel(t *testing.T) {
	names := make(chan int)
	close(names)
	fmt.Println("started")

	for name := range names {
		fmt.Println("name:", name)
	}

	fmt.Println("finished")
}
