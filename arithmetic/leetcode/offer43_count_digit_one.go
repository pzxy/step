package leetcode

/**
输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。

例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。

示例 1：

输入：n = 12
输出：5
示例 2：

输入：n = 13
输出：6

限制：

1 <= n <2^31
*/

func countDigitOne(n int) int { //此方法当n很大时，会花费很长时间，超出时间限制。

	ret := 0
	for i := 1; i <= n; i++ {
		ii := i
		for ii > 0 {
			if ii%10 == 1 {
				ret++
			}
			ii = ii / 10
		}
	}
	return ret
}

/**
func countDigitOne(n int) int {
    cur,hight,low,digit,res:=n%10,n/10,0,1,0
    for cur!=0||hight!=0{
        if cur==0{
            res+=hight*digit
        }else if cur==1{
            res+=hight*digit+1+low
        }else{
            res+=(hight+1)*digit
        }
        low+=cur*digit
        cur=hight%10
        hight/=10
        digit*=10
    }
    return res
}
*/
func countDigitOne2(n int) int {
	digit := 1
	res := 0
	high := n / 10
	cur := n % 10
	low := 0
	for high != 0 || cur != 0 {
		if cur == 0 {
			res += high * digit
		} else if cur == 1 {
			res += high*digit + low + 1
		} else {
			res += (high + 1) * digit
		}
		low += cur * digit
		cur = high % 10
		high /= 10
		digit *= 10
	}
	return res
}
