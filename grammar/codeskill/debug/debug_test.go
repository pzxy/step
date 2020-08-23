package debug

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestDebugFunc(t *testing.T) {
	Init(true)
	for i := 0; i < 1000000; i++ {
		time.Sleep(time.Millisecond * 100)
		str := ""
		for idx := 0; idx < i; idx++ {
			str = str + fmt.Sprintf("%x", idx)
		}
		RcvRawData.Write(str)
		HeartBeatCounter.Write(strconv.Itoa(i))
		SendRcvPkgMarshalBody.Write("??????")
	}
}
