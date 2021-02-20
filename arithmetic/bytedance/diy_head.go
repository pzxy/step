package bytedance

type MyHead struct {
	a   []int //从index 1 开始
	len int
	cap int
}

func NewMyHead(cap int) *MyHead {
	return &MyHead{make([]int, cap+1), 0, cap}
}

func (h *MyHead) insert(v int) {
	if h.len >= h.cap {
		return
	}
	h.len++
	h.a[h.len] = v
	i := h.len
	rootIdx := h.len / 2
	for rootIdx > 0 && h.a[i] > h.a[rootIdx] {
		h.a[rootIdx], h.a[i] = h.a[i], h.a[rootIdx]
		i = rootIdx
		rootIdx = i / 2
	}
}

func (h *MyHead) headify(limit int, i int) {
	maxIdx := i
	for {
		if 2*i <= limit && h.a[i] < h.a[2*i] {
			maxIdx = 2 * i
		}
		if 2*i+1 <= limit && h.a[maxIdx] < h.a[2*i+1] {
			maxIdx = 2*i + 1
		}
		if i == maxIdx {
			break
		}
		h.a[maxIdx], h.a[i] = h.a[i], h.a[maxIdx]
		i = maxIdx
	}
}

func (h *MyHead) build() {
	if h.len == 0 {
		return
	}
	for i := h.len / 2; i >= 1; i-- {
		h.headify(h.len, i)
	}
}

func (h *MyHead) deleteTop() {
	if h.len == 0 {
		return
	}
	h.a[1] = h.a[h.len]
	h.a = h.a[:h.len]
	h.len--
	h.headify(h.len, 1)
}

func (h *MyHead) sort() {
	h.build()
	limit := h.len
	for limit > 1 {
		h.a[1], h.a[limit] = h.a[limit], h.a[1]
		limit--
		h.headify(limit, 1)
	}
}
