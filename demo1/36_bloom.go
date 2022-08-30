package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	s := strings.Split("10:00:05.308137", ":")
	h, _ := strconv.Atoi(s[0])
	m, _ := strconv.Atoi(s[1])
	sn := strings.Split(s[2], ".")
	sec, _ := strconv.Atoi(sn[0])
	nec, _ := strconv.Atoi(sn[1])
	t := time.Now()
	tt := time.Date(t.Year(), t.Month(), t.Day(), h, m, sec, nec, time.UTC)
	fmt.Println(tt.UnixMilli())

}
