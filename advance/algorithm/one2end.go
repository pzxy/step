package main

import "fmt"

/**
一个无序数组里有99个不重复正整数，范围从1到100，唯独缺少一个整数。如何找出这个缺失的整数？
*/

func main() {
	sli := initSli(uint8(2))
	length := len(sli)
	n := len(sli) - 1

	if int(sli[n]) == length {
		fmt.Println(sli[n])
		return
	}

	for i, j := 0, n; i < j; i, j = i+1, j-1 {
		if sli[i]+sli[j] != 101 {
			if int(sli[i]) + j == 101{
				fmt.Println( sli[i])
			}else {
				fmt.Println( sli[j]+1)
			}
			return
		}
	}
}

//通过位运算。但是题目要求不能通过位运算
func func1() {
	sli := initSli(uint8(18))
	i := 0
	length := len(sli)
	n := len(sli) - 1
	for i = 0; i < n; i += 2 {
		if sli[i]&1 != 1 {
			fmt.Println(sli[i] - 1)
			return
		}
		if sli[i+1]&1 == 1 {
			fmt.Println(sli[i+1] - 1)
			return
		}
	}
	for ; i < length; i++ {
		if sli[i]&1 != 1 {
			fmt.Println(sli[i] - 1)
			return
		}
		fmt.Println(sli[i] + 1)
	}
}

//相减
func func2() {
	sli := initSli(uint8(12))
	i := 0
	length := len(sli)
	n := len(sli) - 1

	for i = 0; i < n; i++ {
		if sli[i+1] != 1+sli[i] {
			fmt.Println(sli[i] + 1)
			return
		}
	}
	for ; i < length; i++ {
		fmt.Println(sli[i] + 1)
	}
}

func initSli(ign uint8) (sli []uint8) {
	if ign > 100 {
		return
	}
	for i := uint8(1); i <= 100; i++ {
		if i == ign {
			continue
		}
		sli = append(sli, i)
	}
	return sli
}
