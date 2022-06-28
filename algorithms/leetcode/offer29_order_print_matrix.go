package leetcode

/**

输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：

输入：matrix =[[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]


限制：

0 <= matrix.length <= 100
0 <= matrix[i].length<= 100

*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	rowMin, rowMax := 0, len(matrix)-1
	colMin, colMax := 0, len(matrix[0])-1

	var ret []int
	for {
		//行->右
		for j := colMin; j <= colMax; j++ {
			ret = append(ret, matrix[rowMin][j])
		}
		rowMin++
		if rowMin > rowMax {
			break
		}
		//列->下
		for i := rowMin; i <= rowMax; i++ {
			ret = append(ret, matrix[i][colMax])
		}
		colMax--
		if colMin > colMax {
			break
		}
		//行->左
		for j := colMax; j >= colMin; j-- {
			ret = append(ret, matrix[rowMax][j])
		}
		//列->上
		rowMax--
		if rowMin > rowMax {
			break
		}
		for i := rowMax; i >= rowMin; i-- {
			ret = append(ret, matrix[i][colMin])
		}
		colMin++
		if colMin > colMax {
			break
		}
	}
	return ret
}

func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	var res []int
	l := 0
	r := len(matrix[0]) - 1
	t := 0
	b := len(matrix) - 1
	for {
		for i := l; i <= r; i++ {
			res = append(res, matrix[t][i])
		}
		t++
		if t > b {
			break
		}
		for i := t; i <= b; i++ {
			res = append(res, matrix[i][r])
		}
		r--
		if l > r {
			break
		}
		for i := r; i >= l; i-- {
			res = append(res, matrix[b][i])
		}
		b--
		if t > b {
			break
		}
		for i := b; i >= t; i-- {
			res = append(res, matrix[i][l])
		}
		l++
		if l > r {
			break
		}
	}
	return res
}
