package datastruct

import "fmt"

/**
1. 归并排序
*/
func mergeDemo(a []int) {
	n := len(a)
	mergeSort(a, 0, n-1)
	fmt.Println(a)
}

func mergeSort(a []int, p int, r int) {
	if p >= r {
		return
	}
	q := p + (r-p)>>1
	mergeSort(a, p, q)   //前半部分
	mergeSort(a, q+1, r) //后半部分
	doMerge4(a, p, q, r) //排序
}

/**
我们申请一个临时数组 tmp，大小与 A[p…r]相同。我们用两个游标 i 和 j，分别指向 A[p…q]和 A[q+1…r]的第一个元素。
比较这两个元素 A[i]和 A[j]，如果 A[i]<=A[j]，我们就把 A[i]放入到临时数组 tmp，并且 i 后移一位，否则将 A[j]放入到数组 tmp，j 后移一位。

继续上述比较过程，直到其中一个子数组中的所有数据都放入临时数组中，再把另一个数组中的数据依次加入到临时数组的末尾，
这个时候，临时数组中存储的就是两个子数组合并之后的结果了。最后再把临时数组 tmp 中的数据拷贝到原数组 A[p…r]中。
*/
func doMerge(pr []int, p, q, r int) {
	if len(pr) == 1 {
		return
	}

	i := p
	j := q + 1
	k := 0
	tmp := make([]int, r-p+1) //创建临时数组

	for i <= q || j <= r { //排序
		if i <= q && j <= r {
			if pr[i] < pr[j] {
				tmp[k] = pr[i]
				k++
				i++
			} else {
				tmp[k] = pr[j]
				k++
				j++
			}
			continue
		}

		if i <= q {
			tmp[k] = pr[i]
			i++
			k++
		}

		if j <= r {
			tmp[k] = pr[j]
			j++
			k++
		}
	}

	//var start, end int //剩下的部分添加到临时数组tmp中

	//全部复制到原始数组中
	for i, v := range tmp {
		pr[p+i] = v
	}
}

func doMerge4(pr []int, p, q, r int) {
	if len(pr) == 1 {
		return
	}
	tmp := make([]int, r-p+1)
	t := 0
	for i, j := p, q+1; i <= q || j <= r; {
		if i <= q && j <= r {
			if pr[i] < pr[j] {
				tmp[t] = pr[i]
				t++
				i++
			} else {
				tmp[t] = pr[j]
				t++
				j++
			}
			continue
		}

		if i <= q {
			tmp[t] = pr[i]
			i++
			t++
		}

		if j <= r {
			tmp[t] = pr[j]
			j++
			t++
		}

	}
	for i, v := range tmp {
		pr[p+i] = v
	}

}

/**
2. 快速排序
原地，不稳定
*/

func quickDemo(a []int) {
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)

}
func quickSort(a []int, p, r int) {
	if p >= r {
		return
	}
	q := partition(a, p, r)
	quickSort(a, p, q-1)
	quickSort(a, q+1, r)
}

//2, 1, 3, 5, 2
func partition2(a []int, p, r int) (q int) {
	pivot := a[r]
	i := p
	for j, v := range a[p:r] { //a[p:r]不包括r，也就是最后一个
		if v < pivot {
			a[i], a[p+j] = a[p+j], a[i] //和选择排序差不多这里。用游标将数组分成两部分，找到小的追加到后面
			i++
		}
	}
	a[i], a[r] = a[r], a[i]
	q = i
	return
}

func partition(a []int, p, r int) (q int) {
	pivot := a[r]
	i := p
	for j, v := range a[p:r] {
		if v < pivot {
			a[i], a[p+j] = a[p+j], a[i]
			i++
		}
	}
	a[i], a[r] = a[r], a[i]
	q = i
	return
}

/**
总结：
归并排序
非原地，稳定
时间复杂度：O(nlogn)，最坏情况O(n*n)
空间复杂度：O(n)，因此不被广泛应用。

快速排序：
原地，不稳定
时间复杂度：O(nlogn)，最坏情况O(n*n)
空间复杂度：1

*/
