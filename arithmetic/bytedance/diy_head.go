package bytedance

type MyHead struct {
	a   []int //从index 1 开始
	cap int
	len int
}

func NewMyHead(cap int) *MyHead {
	return &MyHead{make([]int, cap+1), cap, 0}
}

func (h *MyHead) Insert(v int) {
	if h.len >= h.cap {
		return
	}
	h.len++
	h.a[h.len] = v
	i := h.len
	ii := i / 2
	for ii > 0 && h.a[i] > h.a[ii] {
		h.a[ii], h.a[i] = h.a[i], h.a[ii]
		i = ii
		ii = i / 2
	}
}

func (h *MyHead) headify(i int) {
	for {

	}
}
