package leetcode

import "strconv"

/**
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），
也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，
因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？


示例 1：

输入：m = 2, n = 3, k = 1
输出：3
示例 2：

输入：m = 3, n = 1, k = 0
输出：1
*/
/**
回溯法是深度优先,有的题深度优先和广度优先都可以，有的题只能深度优先，
深度优先，用递归来做
广度优先用循环来做Q队列
*/
var directions = []point2{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func movingCount2(m, n, k int) int {
	if m == 0 || n == 0 {
		return 0
	}
	count := 1
	if k == 0 {
		return count
	}
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	grid[0][0] = 1
	Q := []point2{{0, 0}}
	for len(Q) > 0 {
		p := Q[0]
		Q = Q[1:]
		for _, r := range directions {
			curr := p.add(r)
			if !curr.check(grid, m, n, k) {
				continue
			}
			count++
			grid[curr.x][curr.y] = count
			Q = append(Q, curr)
		}
	}
	return count
}

func (p point2) add(r point2) point2 {
	return point2{p.x + r.x, p.y + r.y}
}
func (p point2) check(grid [][]int, m, n, k int) bool {
	if p.x < 0 || p.x >= m {
		return false
	}
	if p.y < 0 || p.y >= n {
		return false
	}
	if grid[p.x][p.y] != 0 {
		return false
	}
	var sum int
	x := strconv.Itoa(p.x)
	for _, i2 := range x {
		sum += int(i2 - 48)
	}
	y := strconv.Itoa(p.y)
	for _, i2 := range y {
		sum += int(i2 - 48)
	}
	if sum > k {
		return false
	}
	return true

}

type point2 struct {
	x, y int
}

//func movingCount(m, n, k int) int {
//	if m == 0 || n == 0 {
//		return 0
//	}
//	gird := make([][]int, m)
//	for i := range gird {
//		gird[i] = make([]int, n)
//	}
//
//	count := find2BFS(gird, point2{0, 0}, m, n, k)
//
//	return count
//}
//func movingCount2(m, n, k int) int {
//	if m == 0 || n == 0 {
//		return 0
//	}
//	gird := make([][]int, m)
//	for i := range gird {
//		gird[i] = make([]int, n)
//	}
//
//	count := find2DFS(gird, point2{0, 0}, m, n, k)
//
//	return count
//}
//
//type point2 struct {
//	i int
//	j int
//}
//
//var dirs2 = [4]point2{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
//
//func (p point2) add(r point2) point2 {
//	return point2{p.i + r.i, p.j + r.j}
//}
//func (p point2) at(gird [][]int, m, n, k int) bool {
//	if p.i < 0 || p.i >= m {
//		return false
//	}
//	if p.j < 0 || p.j >= n {
//		return false
//	}
//	if gird[p.i][p.j] != 0 {
//		return false
//	}
//	if !p.kOk1(k) {
//		return false
//	}
//	return true
//
//}
//
//func find2BFS(gird [][]int, start point2, m, n, k int) int {
//	Q := []point2{start}
//	count := 1
//	gird[start.i][start.j] = count
//
//	for len(Q) > 0 {
//		curr := Q[0]
//		Q = Q[1:]
//		for _, r := range dirs2 {
//			next := curr.add(r)
//			if !next.at(gird, m, n, k) {
//				continue
//			}
//			count++
//			gird[next.i][next.j] = count
//			Q = append(Q, next)
//		}
//	}
//
//	return count
//}
//func find2DFS(gird [][]int, start point2, m, n, k int) int {
//	if !start.at(gird, m, n, k) {
//		return 0
//	}
//	count := 1
//	gird[start.i][start.j] = count
//	for _, r := range dirs2 {
//		next := start.add(r)
//		count += find2DFS(gird, next, m, n, k)
//	}
//
//	return count
//}
//func (p point2) kOk2(k int) bool {
//	res := 0
//	m := p.i
//	n := p.j
//	for m > 0 {
//		res += m % 10
//		m /= 10
//	}
//	for n > 0 {
//		res += n % 10
//		n /= 10
//	}
//	if res > k {
//		return false
//	}
//	return true
//}
//func (p point2) kOk1(k int) bool {
//	sI := strconv.Itoa(p.i)
//	sum := 0
//	for _, i3 := range sI {
//		v, _ := strconv.Atoi(string(i3))
//		sum += v
//	}
//	if sum > k {
//		return false
//	}
//	sJ := strconv.Itoa(p.j)
//	for _, i3 := range sJ {
//		v, _ := strconv.Atoi(string(i3))
//		sum += v
//	}
//	if sum > k {
//		return false
//	}
//	return true
//}
//
//func checkCoordinate(gird [][]int, i, j, k int) bool {
//	if gird[i][j] != 0 {
//		return false
//	}
//	sI := strconv.Itoa(i)
//	sum := 0
//	for _, i3 := range sI {
//		v, _ := strconv.Atoi(string(i3))
//		sum += v
//	}
//	if sum > k {
//		return false
//	}
//	sJ := strconv.Itoa(j)
//	for _, i3 := range sJ {
//		v, _ := strconv.Atoi(string(i3))
//		sum += v
//	}
//	if sum > k {
//		return false
//	}
//	return true
//}
