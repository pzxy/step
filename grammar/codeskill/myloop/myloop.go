package myloop

import "time"

//protect loop broke pss when system on gc state
func IdleSleep() {
	time.Sleep(time.Millisecond * 1)
}
