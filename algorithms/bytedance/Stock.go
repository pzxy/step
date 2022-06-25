package bytedance

import "fmt"

/**
给定一个数组代表股票每天的价格，请问只能买卖一次的情况下，最大化利润是多少？日期不重叠的情况下，可以买卖多次呢？
输入: {100, 80, 120, 130, 70, 60, 100, 125}
1）只能买一次：65（60 买进，125 卖出）
2）可以买卖多次: 115（80买进，130卖出；60 买进，125卖出）
输出买卖的序列和最大利润
*/
func stockMore(s []int) {
	if len(s) == 0 {
		return
	}
	lowest, pre := s[0], s[0]
	var ret []int
	for i, v := range s[1:] {
		if v > pre { //升序
			pre = v
			if i == len(s)-2 {
				ret = append(ret, lowest)
				ret = append(ret, pre)
			}
			continue
		}
		//降序
		if pre-lowest > 0 {
			ret = append(ret, lowest)
			ret = append(ret, pre)
		}
		pre = v
		lowest = v
	}

	for i := 0; i < len(ret); i++ {
		fmt.Println(ret[i], ret[i+1])
		i++
	}
}

func stockOnce(s []int) {
	if len(s) == 0 {
		return
	}
	lowest, pre := s[0], s[0]
	var retLowest []int
	var retPre []int
	for i, v := range s[1:] {
		if v > pre { //升序
			pre = v
			if i == len(s)-2 {
				retLowest = append(retLowest, lowest)
				retPre = append(retPre, pre)
			}
			continue
		}
		//降序
		if pre-lowest > 0 {
			retLowest = append(retLowest, lowest)
			retPre = append(retPre, pre)
		}
		pre = v
		lowest = v
	}

	var max int
	var in, out int
	for i := 0; i < len(retLowest); i++ {
		for j := i; j < len(retPre); j++ {
			if retPre[j]-retLowest[i] > max {
				in = retLowest[i]
				out = retPre[j]
				max = retPre[j] - retLowest[i]
			}
		}
	}
	fmt.Println(in, out)
}
