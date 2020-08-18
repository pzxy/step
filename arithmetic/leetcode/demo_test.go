package leetcode

import (
	"fmt"
	"testing"
)

func Test_LengthOf(t *testing.T) {
	fmt.Println(lengthOf("aaaaadcaddd"))
}

func Test_Fibonacci(t *testing.T) {
	n := 7
	i := Fibo1(n)
	j := Fibo2(n)
	k := Fibo3(n)
	fmt.Printf("%v  %v  %v", i, j, k)
}

func Test_FindRepeatFrom100Int(t *testing.T) {
	sli := initSli(uint8(2))
	length := len(sli)
	n := len(sli) - 1

	if int(sli[n]) == length {
		fmt.Println(sli[n])
		return
	}

	for i, j := 0, n; i < j; i, j = i+1, j-1 {
		if sli[i]+sli[j] != 101 {
			if int(sli[i])+j == 101 {
				fmt.Println(sli[i])
			} else {
				fmt.Println(sli[j] + 1)
			}
			return
		}
	}
}

func Test_Prime(t *testing.T) {
	//test2()
	fmt.Print(allPrimeSum(1, 100000, 0))
}

func Test_RevSlice(t *testing.T) {
	//test2()
	revSlice()
}

func Test_RevSingleList(t *testing.T) {
	var node link
	node.value = 1
	node.next = &link{value: 2}
	node.next.next = &link{value: 3}
	node.next.next.next = &link{value: 4}
	node.next.next.next.next = &link{value: 5}
	node.next.next.next.next.next = &link{value: 6}
	node.next.next.next.next.next.next = &link{value: 7}

	//node.reverseLink().printLink()//全部反转
	//每三个反转
	node.revlink2()
	node.printLink()
}

func Test_StepGrid(t *testing.T) {
	m := 6
	n := 4
	fmt.Printf("%d * %d的方格一共 %d 种走法\n", m, n, walk(m, n))
	fmt.Printf("%d * %d的方格一共 %d 种走法\n", m, n, walks(m, n))

}

func Test_Int2Chinese(t *testing.T) {
	fmt.Println(int2Chinese(33402))
}

func Test_largestNumber(t *testing.T) {
	//s1 := []int{3, 30, 34, 5, 9}
	//s2 := []int{824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247}
	s3 := []int{1, 2, 3, 1}
	fmt.Println(largestNumber(s3))
}
