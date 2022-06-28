package leetcode

/**
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

示例 1:

输入:
[
 [1,3,1],
 [1,5,1],
 [4,2,1]
]
输出: 12
解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物

提示：

0 < grid.length <= 200
0 < grid[0].length <= 200
*/

type point3 struct {
	x, y int
	sum  int
}

func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}
	return walk3(grid)
}

var dirs3 = [2]point3{{x: 1, y: 0}, {x: 0, y: 1}}

func (p point3) add(r point3) point3 {
	return point3{p.x + r.x, p.y + r.y, p.sum}
}

func walk3(grid [][]int) (ret int) {
	start := point3{x: 0, y: 0, sum: grid[0][0]}
	lenX := len(grid)
	lenY := len(grid[0])
	end := point3{x: lenX - 1, y: lenY - 1}
	Q := []point3{start}
	ret = grid[start.x][start.y]
	for len(Q) > 0 {
		q := Q[0]
		Q = Q[1:]
		for _, r := range dirs3 {
			tmp := q.add(r)
			if !tmp.check(lenX, lenY) {
				continue
			}
			tmp.sum = q.sum + grid[tmp.x][tmp.y]
			if tmp.ifEnd(end) {
				if tmp.sum > ret {
					ret = tmp.sum
				}
				continue
			}
			Q = append(Q, tmp)
		}
	}
	return
}

func (p point3) check(lenX, lenY int) bool {
	if p.x < lenX && p.y < lenY {
		return true
	}
	return false
}

func (p point3) ifEnd(end point3) bool {
	if p.x == end.x && p.y == end.y {
		return true
	}
	return false
}

/**
常规动态规划思路
找状态：每一步礼物的最大值
转移方程：dp[i][j]=grid[i-1][j-1]+max(dp[i][j-1],dp[i-1][j])
*/
func maxValue4dp(grid [][]int) int {
	lenGrid := len(grid)
	if lenGrid == 0 {
		return -1
	}
	dp := make([][]int, lenGrid+1)
	for i := 0; i <= lenGrid; i++ {
		dp[i] = make([]int, len(grid[0])+1)
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
		}
	}
	return dp[len(grid)][len(grid[0])]
}

func maxValue4dp2(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}

	dp := make([][]int, len(grid)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[0])+1)
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
