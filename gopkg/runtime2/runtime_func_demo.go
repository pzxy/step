package main

import (
	"fmt"
	"runtime"
	"syscall"
	"time"
)

func main() {
	//runtimeKeepAlive()
	//runtimeLockOsThread()
	//runtimeGC()
	runtimeDemo1()
}

func runtimeDemo1() {
	runtime.GC()
	fmt.Printf("%d\n", runtime.NumCPU())       //逻辑cpu数
	fmt.Printf("%s\n", runtime.GOROOT())       //返回goroot 的路径
	fmt.Printf("%d\n", runtime.GOMAXPROCS(12)) //设置最大逻辑cup数
	fmt.Printf("%d\n", runtime.NumCgoCall())   //当前 进程中调用cgo的数量 就是调用c的库
	fmt.Printf("%s\n", runtime.Version())      //
	fmt.Printf("%s\n", runtime.GOROOT())       //
	//runtime2.GoroutineProfile() 还没测试
}

type File struct{ d int }

func runtimeKeepAlive() {
	d, err := syscall.Open("/Users/pzxy/Workspace/go/src/step/gopkg/runtime2/re.json", syscall.O_RDONLY, 0)
	if err != nil {
		fmt.Printf("syscall.Open err:%s", err)
	}
	p := &File{d}

	runtime.SetFinalizer(p, func(p *File) {
		syscall.Close(p.d)
		fmt.Printf("%s\n", "垃圾会输了")
	}) //给p对象绑定方法，垃圾回收的时候进行监听 https://www.jianshu.com/p/84bac7932394
	var buf [10]byte
	_, err = syscall.Read(p.d, buf[:])
	fmt.Printf("%v\n", p)
	runtime.GC()         //下面还有用到p的话 垃圾不会回收
	runtime.KeepAlive(p) //保持p活跃状态，在syscall。Read返回结果之前，不被回收。
	runtime.GC()         //下面还有用到p的话 垃圾不会回收
	//fmt.Printf("%v\n", p)
	time.Sleep(time.Second * 3)
	//发现 垃圾回收了这句话纸打印了一次，说明 在runtime.KeepAlive(p) 这之前的的情况下，没有被打印，垃圾回收只会回收下面用不到的对象，
	//其实这里不必用runtime.KeepAlive(p)，即使打印一个 p对象 他也不会被回收
}

type Student struct {
	name string
}

func runtimeGC() {
	var i *Student = new(Student)
	runtime.SetFinalizer(i, func(i interface{}) {
		println("垃圾回收了")
	})
	runtime.GC()
	time.Sleep(time.Second * 3)
}

//runtime2.LockOSThread和runtime.UnlockOSThread函数
//前者调用会使调用他的Goroutine与当前运行它的M锁定到一起，后者调用会解除这样的锁定。
//多次调用前者不会出现任何问题，但最后一次调用的记录会被保留，
//即时之前没有调用前者，对后者的调用也不会产生任何副作用
func runtimeLockOsThread() {
	go calcSum2()
	go calcSum1()
	fmt.Printf("%d\n", runtime.NumGoroutine()) //当前 携程树木
	time.Sleep(3 * time.Second)
	fmt.Printf("%d\n", runtime.NumGoroutine()) //当前 携程树木

	time.Sleep(time.Second * 100)
	//总结 2先执行，先打印的1 ，这是由于LockOsThread 将这个 goroutine 放到了一个线程中，这个线程中只有这一个goroutine，
	//所以我们可以吧一些优先级高的方法 使用LockOsThread 的方式让其先执行
}

func calcSum1() {
	runtime.LockOSThread()
	start := time.Now()
	fmt.Println("1 start :", start)
	count := 0
	for i := 0; i < 10000000000; i++ {
		count += i
	}
	end := time.Now()
	fmt.Println("1 end :", end)
	defer runtime.UnlockOSThread()
}

func calcSum2() {
	start := time.Now()
	fmt.Println("2 start :", start)
	count := 0
	for i := 0; i < 10000000000; i++ {
		count += i
	}
	end := time.Now()
	fmt.Println("2 end :", end)
}
