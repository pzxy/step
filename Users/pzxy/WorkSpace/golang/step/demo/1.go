package main

import (
	"fmt"
)

func main() {
	var a string
	b := 0
	for {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 {
			break
		} else {

			fmt.Printf("%d\n", printKIdx(a, b))
		}
	}
}

func printKIdx(s string, k int) int {
	m := make(map[rune]int, 0)
	for i, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = i
		}
	}
	sortS := quickSort([]rune(s), 0, len(s)-1)
	if len(s) < k {
		return m[rune(sortS[len(sortS)-1])]
	}
	return m[rune(sortS[k-1])]
}

func quickSort(s []rune, p, r int) string {
	if p >= r {
		return string(s)
	}
	q := partition(s, p, r)
	quickSort(s, p, q-1)
	quickSort(s, q+1, r)
	return string(s)
}
func partition(s []rune, p, r int) int {
	min := s[r]
	i := p
	for j, v := range s[p:r] {
		if v < min {
			s[i], s[p+j] = s[p+j], s[i]
			i++
		}
	}
	s[i], s[r] = s[r], s[i]
	return i
}
