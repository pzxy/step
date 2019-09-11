package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	sum := Test11(2,3)
	Convey("TestDt", t, func() {
		Convey("Data transfer integrity verification failed ", func() {
			time.Sleep(time.Second * 3)
			b := sum == 5
			So(b, ShouldBeTrue)
		})
	})
}

