package leetcode

/**
请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。
路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。
如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。
例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

["a","b","c","e"],
["s","f","c","s"],
["a","d","e","e"],

但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。



示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false

*/

func exist2(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs2(board, point{i, j}, word, 0) {
				return true
			}
		}
	}
	return false
}

type point struct {
	i int
	j int
}

var dirs = [4]point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} // 上 左 下 右

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) check(gird [][]byte) bool {
	if p.i < 0 || p.i >= len(gird) {
		return false
	}
	if p.j < 0 || p.j >= len(gird[p.i]) {
		return false
	}

	return true
}

func dfs2(board [][]byte, curr point, word string, k int) bool {
	if board[curr.i][curr.j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	temp := board[curr.i][curr.j]
	board[curr.i][curr.j] = byte(' ')
	for _, r := range dirs {
		nextPoint := curr.add(r)
		if !nextPoint.check(board) {
			continue
		}
		if dfs2(board, nextPoint, word, k+1) {
			return true
		}
	}
	board[curr.i][curr.j] = temp
	return false
}

/**

 */
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}
func dfs(board [][]byte, i int, j int, word string, k int) bool {
	if board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	temp := board[i][j]
	board[i][j] = byte(' ')
	if 0 <= i-1 && dfs(board, i-1, j, word, k+1) {
		return true
	}
	if i+1 < len(board) && dfs(board, i+1, j, word, k+1) {
		return true
	}
	if 0 <= j-1 && dfs(board, i, j-1, word, k+1) {
		return true
	}
	if j+1 < len(board[0]) && dfs(board, i, j+1, word, k+1) {
		return true
	}
	board[i][j] = temp
	return false
}
