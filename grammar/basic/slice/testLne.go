package main

import "fmt"

func main() {
	var sli []int
	sli = append(sli, 1)
	sli = append(sli, 1)
	sli = append(sli, 1)
	sli = append(sli, 1)
	sli = append(sli, 1)
	sli = append(sli, 1)
	sli = append(sli, 1)
	sli = append(sli, 1)
	sli = append(sli, 1)
	sli = append(sli, 1)
	sli = append(sli, 1)

	fmt.Println(len(sli), cap(sli))
	fmt.Println(sli)
}
