package geek

import "fmt"

/**
1. BF算法,Brute Force ，中文叫作暴力匹配算法，也叫朴素匹配算法
2. RK算法,Rabin-Karp ，是由它的两位发明者 Rabin 和 Karp 的名字来命名的。 BF 算法的升级版。
*/

func bfDemo() {
	b := bf("abcdef", "")
	fmt.Println(b)
	fmt.Println()
}

func rkDemo() {
	rk := indexRabinKarp("abcdef", "")
	fmt.Println(rk)
	fmt.Println()
}

/**
1. BF，适用于短字符匹配
*/
func bf(s1 string, s2 string) bool {
	b1 := []byte(s1)
	b2 := []byte(s2)
	n := len(b1)
	m := len(b2)
	var match bool
	for i := 0; i < n-m+1; i++ {
		count := i
		for j, v := range b2 {
			if b1[count] != v {
				break
			}
			if j == m-1 {
				match = true
				break
			}
			count++
		}

		if match {
			return true
		}

	}
	return false
}

/**
2. RK算法，适用于长字符匹配。比较哈希值，哈希值很特殊。
*/

const primeRK = 16777619

func hashStr(sep string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, primeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}
func indexRabinKarp(s, substr string) int {
	// Rabin-Karp search
	hashss, pow := hashStr(substr)
	n := len(substr)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(s[i])
	}
	if h == hashss && s[:n] == substr {
		return 0
	}
	for i := n; i < len(s); {
		h *= primeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashss && s[i-n:i] == substr {
			return i - n
		}
	}
	return -1

}
