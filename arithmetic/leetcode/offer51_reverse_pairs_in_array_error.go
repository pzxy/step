package leetcode

/**
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。
输入一个数组，求出这个数组中的逆序对的总数。

示例 1:

输入: [7,5,6,4]
输出: 5

限制：

0 <= 数组长度 <= 50000
*/

func reversePairs3(nums []int) (ret int) {
	lenNums := len(nums)
	if lenNums == 0 || lenNums == 1 {
		return 0
	}
	tmp := make([]int, 1)
	tmp[0] = nums[0]
	var ifAdd bool
	var count int
	for _, v := range nums[1:] {
		ifAdd = false
		count = 0
		for i, v2 := range tmp {
			if v < v2 {
				continue
			}
			if v == v2 {
				count--
				continue
			}
			t1 := make([]int, i)
			copy(t1, tmp[:i])
			t1 = append(t1, v)
			tmp = append(t1, tmp[i:]...)
			ret += i + count
			ifAdd = true
			break
		}
		if !ifAdd {
			ret += len(tmp) + count
			tmp = append(tmp, v)
		}

	}
	return
}

func reversePairs2(nums []int) int {
	return mergeSort2(nums, 0, len(nums)-1)
}

//官方
func mergeSort2(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end-start)/2
	cnt := mergeSort2(nums, start, mid) + mergeSort2(nums, mid+1, end)
	tmp := []int{}
	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			cnt += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
		cnt += end - (mid + 1) + 1
	}
	for ; j <= end; j++ {
		tmp = append(tmp, nums[j])
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}
	return cnt
}

func reversePairs(a []int) int {
	n := len(a)
	return mergeSort(a, 0, n-1)
}

func mergeSort(a []int, p int, r int) int {
	if p >= r {
		return 0
	}
	q := p + (r-p)>>1
	r1 := mergeSort(a, p, q)   //前半部分
	r2 := mergeSort(a, q+1, r) //后半部分
	r3 := doMerge(a, p, q, r)  //排序
	return r1 + r2 + r3
}

/**
我们申请一个临时数组 tmp，大小与 A[p…r]相同。我们用两个游标 i 和 j，分别指向 A[p…q]和 A[q+1…r]的第一个元素。
比较这两个元素 A[i]和 A[j]，如果 A[i]<=A[j]，我们就把 A[i]放入到临时数组 tmp，并且 i 后移一位，否则将 A[j]放入到数组 tmp，j 后移一位。

继续上述比较过程，直到其中一个子数组中的所有数据都放入临时数组中，再把另一个数组中的数据依次加入到临时数组的末尾，
这个时候，临时数组中存储的就是两个子数组合并之后的结果了。最后再把临时数组 tmp 中的数据拷贝到原数组 A[p…r]中。
*/
func doMerge(pr []int, p, q, r int) (ret int) {
	if len(pr) == 1 {
		return
	}

	i := p
	j := q + 1
	k := 0
	tmp := make([]int, r-p+1) //创建临时数组

	for i <= q && j <= r { //排序
		if pr[i] < pr[j] {
			tmp[k] = pr[i]
			k++
			i++
		} else {
			tmp[k] = pr[j]
			k++
			j++
			ret += q - i + 1
		}
	}

	var start, end int //剩下的部分添加到临时数组tmp中
	if i != q {
		start = j
		end = r
	} else {
		start = i
		end = q
	}
	for start <= end {
		tmp[k] = pr[start]
		start++
		k++
	}

	//全部复制到原始数组中
	for i, v := range tmp {
		pr[p+i] = v
	}
	return
}
