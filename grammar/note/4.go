package main

import (
	"fmt"
	"regexp"
)

func main() {
	sss := "--cp="
	s, _ := regexp.Compile("^--?(cp|configProvider)=?")
	if loc := s.FindStringIndex(sss); loc != nil { //最后一个值的下标，从1开始，比如--cp=，最后一个为5
		if sss[loc[1]-1] != '=' {
			fmt.Println("123")
			fmt.Println(string(loc[1] - 1))
		}
		fmt.Println(string(loc[1] - 1))
	}
}
