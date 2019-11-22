package main

import (
	"fmt"
	"sort"
)

func main(){
	sortSlice1()
}

func sortSlice1()  {
	sli := []int{2,3,5,7}
	sort.Slice(sli, func(i, j int) bool {
		if sli[i] > sli[j]{
			return true
		}
		return  true
	})
	fmt.Printf("%v \n",sli)
}
