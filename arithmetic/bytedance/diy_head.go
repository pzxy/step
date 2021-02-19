package bytedance

type MyHead struct {
	a   []int //从index 1 开始
	len int
	cap int
}

func NewMyHead(cap int) *MyHead {
	return &MyHead{make([]int, cap+1), 0, cap}
}

func (h *MyHead) Insert(v int) {
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

func (h *MyHead) headify(i int) {
	for {

	}
}
