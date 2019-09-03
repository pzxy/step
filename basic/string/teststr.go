package main

import (
	"fmt"
	"strconv"
)

func main(){
	robotId := 1
	test01(robotId)
}
func test01(robotId int){
	s := strconv.Itoa(robotId)
	path := "C://dt_data_file"+s
	fmt.Println(path)
}