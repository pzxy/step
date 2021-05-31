package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	sss := "--cp="
	s, _ := regexp.Compile("^--?(cp|configProvider)=?")
	if loc := s.FindStringIndex(sss); loc != nil { //最后一个值的下标，从1开始，比如--cp=，最后一个为5
		if sss[loc[1]-1] != '=' {
			fmt.Println("123")

			fmt.Println(string(loc[1] - 1))
			if sss[loc[1]-1] != '=' {
				sss = "-cp=" + "consul.http://localhost:8500"
			}
			fmt.Println(sss)
		}
		fmt.Println(string(loc[1] - 1))
	}

	b := []byte{52, 48, 52, 32, 112, 97, 103, 101, 32, 110, 111, 116, 32, 102, 111, 117, 110, 100, 10}
	fmt.Println(string(b))
	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("2")
		}
	}

}
