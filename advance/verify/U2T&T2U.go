package main

import "fmt"

//u2t 无符号转补码（有符号）
//t2u 补码（有符号）转无符号

func main() {
	trunc_u()
}

func T2U() {
	var i int16 = -12345
	fmt.Printf("有符号: %b, \n无符号:%b \n", i, uint16(i))
	//有符号:-11000000111001,
	//无符号:1100111111000111
	//-11000000111001, 的补码和1100111111000111的值一样。有无符号，只是换了一种解释位模式的方式而已。
}

func U2T() {
	var i uint8 = 129
	fmt.Printf("有符号: %b,%v, \n无符号:  %b,%v \n", int8(i), int8(i), i, i)

}

func U2T_2() {
	var i int8 = -126
	fmt.Printf("有符号: %v,%v, \n无符号:  %v,%v \n", byte(i), byte(i), i, i)

}

//总结：如果8位的话，在0-127的范围里面转换是没问题的。
//如果不在这个范围， 无符号转补码（有符号），则需要减去 -2^8
// 补码（有符号）转无符号，则需要加上。2^8

func compare() { //go里面没有隐式的类型转换。
	//i := int8(2-3)
	//ii :=  uint8(1)- i
	//fmt.Println()
}

func T2U_2() {
	var i int32 = -12345

	fmt.Printf("有符号: %b, \n无符号:%b \n", i, uint16(i))
	//有符号:-11000000111001,
	//无符号:1100111111000111
	//-11000000111001, 的补码和1100111111000111的值一样。有无符号，只是换了一种解释位模式的方式而已。
}

//无符号的截断就是无符号，
// 有符号(补码)的截断是将截断后的数，再求补码，就是截断的值。
// 注意，是对原来的值的补码的截取，原值是有没有符号没关系。）
func trunc() {

	var i int16 = 129 //00000000 10000001
	ii := int8(i)     //11111111 ，10000001的补码

	fmt.Println(i, ii)
	fmt.Printf("%b,%b \n", i, ii)

}

func trunc_u() {
	var i int16 = -267
	ii := uint8(i)

	fmt.Println(i, ii)
	fmt.Printf("%b,%b \n", i, ii)

}
