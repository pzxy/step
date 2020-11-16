package main

import (
	"io"
	"net"
	"os"
	"sync"
	"sync/atomic"
)

/**
失败2：
既然执行过 Once.Do 方法也可能因为函数执行失败的原因未初始化资源，并且以后也没机会再次初始化资源，


*/
func main() {
	var once sync.Once
	var googleConn net.Conn // 到Google网站的一个连接

	once.Do(func() {
		// 建立到google.com的连接，有可能因为网络的原因，googleConn并没有建立成功，此时它的值为nil
		googleConn, _ = net.Dial("tcp", "google.com:80")
	})
	// 发送http请求
	googleConn.Write([]byte("GET / HTTP/1.1\r\nHost: google.com\r\n Accept: */*\r\n\r\n"))
	io.Copy(os.Stdout, googleConn)
}

/**
解决方案：

我们所做的改变就是 Do 方法和参数 f 函数都会返回 error，如果 f 执行失败，会把这个错误信息返回。

对 slowDo 方法也做了调整，如果 f 调用失败，我们不会更改 done 字段的值，
这样后续 degoroutine 还会继续调用 f。如果 f 执行成功，才会修改 done 的值为 1。
*/

// 一个功能更加强大的Once
type Once struct {
	m    sync.Mutex
	done uint32
}

// 传入的函数f有返回值error，如果初始化失败，需要返回失败的error
// Do方法会把这个error返回给调用者
func (o *Once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 { //fast path
		return nil
	}
	return o.slowDo(f)
}

// 如果还没有初始化
func (o *Once) slowDo(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 { // 双检查，还没有初始化
		err = f()
		if err == nil { // 初始化成功才将标记置为已初始化
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}
