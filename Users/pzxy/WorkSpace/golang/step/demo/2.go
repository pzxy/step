package main

import (
	"fmt"
	"math"
)

func main() {
	n := 0
	s := make([]int, 1)
	s[0] = -1
	fmt.Scan(&n)

	for j := 0; j < n; j++ {
		x := 0
		fmt.Scan(&x)
		s = append(s, x)

	}
	fmt.Printf("%d\n", treeNodePath(s))
}

func treeNodePath(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	dfs(s, []int{}, 1)
	return leastPath
}

var leastValue = math.MaxInt64
var leastPath []int

func dfs(s []int, path []int, i int) {
	if i > len(s) || s[i] == -1 {
		if path[len(path)-1] < leastValue {
			leastValue = path[len(path)-1]
			leastPath = path
		}
		return
	}
	path = append(path, s[i])
	dfs(s, path, i*2)
	dfs(s, path, i*2+1)
}
