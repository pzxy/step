package main

import (
	"bytes"
	"sync"
)

var buffers = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func GetBuffer() *bytes.Buffer {
	return buffers.Get().(*bytes.Buffer)
}

func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	buffers.Put(buf)
}

/**
pin 方法会将此 goroutine 固定在当前的 P 上，避免查找元素期间被其它的 P 执行。
固定的好处就是查找元素期间直接得到跟这个 P 相关的 local。有一点需要注意的是，pin 方法在执行的时候，
如果跟这个 P 相关的 local 还没有创建，或者运行时 P 的数量被修改了的话，就会新创建 local。
*/

/**
func (p *Pool) Get() interface{} {
    // 把当前goroutine固定在当前的P上
    l, pid := p.pin()

    x := l.private // 优先从local的private字段取，快速
    l.private = nil
    if x == nil {
        // 从当前的local.shared弹出一个，注意是从head读取并移除
        x, _ = l.shared.popHead()
        if x == nil { // 如果没有，则去偷一个
            x = p.getSlow(pid)
        }
    }
    runtime_procUnpin()
    // 如果没有获取到，尝试使用New函数生成一个新的
    if x == nil && p.New != nil {
        x = p.New()
    }
    return x
}
*/
