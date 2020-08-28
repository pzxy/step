package leetcode

import "math"

/**
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），
每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？
例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

示例 1：

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1
示例2:

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 ×3 ×4 = 36

2 <= n <= 58
*/

/**
贪心算法
这里是数学的方式，已经推论出了最大段是3才是最大。
只要保证3，还有就是如果余1那么就分22，或者不分，直接为4。
贪心算法是建立在数学推论已有的基础上的。
*/

func cuttingRope1(n int) int {
	if n <= 3 {
		return n - 1
	}
	a := n / 3
	b := n % 3
	if b == 0 {
		return int(math.Pow(3, float64(a)))
	}
	if b == 1 {
		return int(math.Pow(3, float64(a-1)) * 4)
	}
	return int(math.Pow(3, float64(a)) * 2)

}

/**
动态规划是特殊情况返回。然后以一般情况拆成小问题。
*/
func cuttingRope(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3

	for i := 4; i <= n; i++ {
		var max float64
		for j := 1; j <= i/2; j++ { //之所以要除以2是因为，1和3与3和1是一样的。
			max = math.Max(max, float64(dp[j]*dp[i-j]))
		}
		dp[i] = int(max)
	}
	return dp[n]
}
