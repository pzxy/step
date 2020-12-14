package bytedance

//核心思想,先将每个数组排序成有序，然后再分别进行两个有序数组的归并排序。
func Merge4ThreeArray(a, b, c []int) []int {
	return Merge4TwoArray(Merge4TwoArray(a, b), c)
}

func Merge4TwoArray(a, b []int) []int {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	Merge4Array(a, 0, len(a)-1)
	Merge4Array(b, 0, len(b)-1)
	i, j := 0, 0
	lenAIdx := len(a) - 1
	lenBIdx := len(b) - 1
	var ret []int
	for i <= lenAIdx && j <= lenBIdx {
		if a[i] <= b[j] {
			ret = append(ret, a[i])
			i++
		} else {
			ret = append(ret, b[j])
			j++
		}
	}
	if i > lenAIdx {
		ret = append(ret, b[j:]...)
	} else {
		ret = append(ret, a[i:]...)
	}
	return ret
}

func Merge4Array(a []int, p, r int) {
	if p >= r {
		return
	}
	q := p + (r-p)>>1
	Merge4Array(a, p, q)
	Merge4Array(a, q+1, r)
	merge(a, p, q, r)
}

func merge(a []int, p, q, r int) {
	var tmp []int
	i, j := p, q+1
	for i <= q && j <= r {
		if a[i] < a[j] {
			tmp = append(tmp, a[i])
			i++
		} else {
			tmp = append(tmp, a[j])
			j++
		}
	}
	if i > q {
		for j <= r {
			tmp = append(tmp, a[j])
			j++
		}
	} else {
		for i <= q {
			tmp = append(tmp, a[i])
			i++
		}
	}

	for _, v := range tmp {
		if p > r {
			break
		}
		a[p] = v
		p++
	}
}
