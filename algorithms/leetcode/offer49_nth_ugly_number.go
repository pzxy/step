package leetcode

/**
我们把只包含质因子1、2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。

示例:

输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
说明:

1是丑数。
n不超过1690。

/**
丑数由 2*i,3*i,5*i 组成。
第一个丑数是 1
第二个丑数是 min(2*1,3*1,5*1)=2
第三个丑数是 min(2*2,3*1,5*1)=3
第四个丑数是 min(2*2,3*2,5*1)=4
第五个丑数是 min(2*3,3*2,5*1)=5
第六个丑数是 min(2*3,3*2,5*2)=6
第k个丑数是 dp[k]=min(2*dp[t2],3*dp[t3],5*dp[t5])
*/
func nthUglyNumber(n int) int {
	dp := make([]int, n+2)
	dp[0] = 0
	dp[1] = 1
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	for a, b, c, i := 1, 1, 1, 2; i <= n; i++ {
		dp[i] = min(dp[a]*2, min(dp[b]*3, dp[c]*5))
		if dp[i] == dp[a]*2 {
			a++
		}
		if dp[i] == dp[b]*3 {
			b++
		}
		if dp[i] == dp[c]*5 {
			c++
		}
	}
	return dp[n]
}
