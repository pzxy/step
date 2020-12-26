package bytedance

//https://zhuanlan.zhihu.com/p/35797448
type MyMap struct {
	bucket []Entry
	cap    int
}

type Entry struct {
	key, value string
	next       *Entry
}

func NewMyMap(cap int) *MyMap {
	return &MyMap{
		bucket: make([]Entry, 10),
		cap:    cap,
	}
}

func hashCode(key string, length int) int {
	b := []byte(key)
	var sum int
	for _, v := range b {
		sum += int(v)
	}
	return sum % length
}

func (m *MyMap) Put(key, value string) {

	entry := Entry{key, value, nil}
	index := hashCode(key, m.cap)

	e := &m.bucket[index]
	//没找到

	if e.key == "" {
		*e = entry
		return
	}

	if e.key == entry.key {
		*e = entry
		return
	}

	e = e.next
	for e != nil {
		if e.key == entry.key {
			*e = entry
			return
		}
		e = e.next
	}
}
