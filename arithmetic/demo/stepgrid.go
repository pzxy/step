package main

import "fmt"

//m*n的方格从左下角走到右上角,只能向上 向右走

func createGrip(m int, n int) [][]int {
	row := make([][]int, m)
	for i := range row {
		row[i] = make([]int, n)
	}
	return row
}

//递归法
func walk(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return walk(m-1, n) + walk(m, n-1)
}

//动态规划
func walks(m int, n int) int {
	row := make([][]int, m)
	for i := range row {
		row[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		row[i][0] = 1
	}
	for i := 0; i < n; i++ {
		row[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			row[i][j] = row[i][j-1] + row[i-1][j]
		}
	}
	return row[m-1][n-1]

}
func main() {
	m := 6
	n := 4
	fmt.Printf("%d * %d的方格一共 %d 种走法\n", m, n, walk(m, n))
	fmt.Printf("%d * %d的方格一共 %d 种走法\n", m, n, walks(m, n))

}
