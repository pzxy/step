package leetcode

/**
二维数组中查找
	a := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
首先这个二维数组是有序的
其次找到查找的办法，我们每次拿最右边的数和要查的数n比较
比如n=7。9>7那么所有数组的第四列都大于7，我们去掉第四列。
然后拿到8和7比较，一样。然后拿到2，2小于7，2是本行的此时的最大值，
然后2下面可能有>=7的数存在，我们去掉第一行，继续拿到4，一直这样比较下去。
*/
func ifFindValFormTwoDimensionalArray(a [][]int, n int) (ret bool) {
	//
	row := len(a)
	if row == 0 {
		return
	}
	col := len(a[0])
	if col == 0 {
		return
	}
	rowIdx := 0
	colIdx := col - 1
	for rowIdx <= row-1 && colIdx >= 0 {
		if a[rowIdx][colIdx] > n {
			colIdx--
			continue
		}
		if a[rowIdx][colIdx] < n {
			rowIdx++
			continue
		}
		if a[rowIdx][colIdx] == n {
			ret = true
			return
		}
	}
	return
}
