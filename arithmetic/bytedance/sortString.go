package bytedance

import "fmt"

func SortString(str string) string {
	m := make(map[rune]int, 0)
	for _, s := range str {
		if v, ok := m[s]; ok {
			v++
			m[s] = v
			continue
		}
		m[s] = 1
	}

	var a []rune
	var count []int
	for k, v := range m {
		count = append(count, v)
		a = append(a, k)
	}
	quickSort(a, count, 0, len(a)-1)
	var ret string
	for i, v := range a {
		for j := 0; j < count[i]; j++ {
			ret += string(v)
		}
	}
	fmt.Println(ret)
	return ret
}

func quickSort(a []rune, i []int, p, r int) {
	if p >= r {
		return
	}

	q := partition(a, i, p, r)
	quickSort(a, i, p, q-1)
	quickSort(a, i, q+1, r)
}

func partition(a []rune, count []int, p, r int) int {
	pivot := count[r]
	i := p
	for j, v := range count[p:r] {
		if pivot < v {
			count[i], count[p+j] = count[p+j], count[i]
			a[i], a[p+j] = a[p+j], a[i]
			i++
		}
	}
	a[i], a[r] = a[r], a[i]
	count[i], count[r] = count[r], count[i]
	return i
}
