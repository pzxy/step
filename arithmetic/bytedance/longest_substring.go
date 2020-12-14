package bytedance

func LongestSubstring(str string) (ret string) {
	if len(str) <= 1 {
		return str
	}
	head, end := 0, 1

	m := make(map[uint8]int, 0)
	m[str[head]] = head

	ret = string(str[head])

	for end < len(str) {
		if _, ok := m[str[end]]; ok {
			if end-head > len(ret) {
				ret = str[head:end]
			}
			m = make(map[uint8]int, 0)
			m[str[head]] = head
			head = end
		} else {
			m[str[end]] = end
		}

		end++
	}
	return
}
