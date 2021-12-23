package main

import "fmt"

//组合模式  目录-文件 ； 省-城市，类似这种业务结构的都能这样用。

func main() {
	f1 := &file{"文件1"}
	f2 := &file{"文件2"}
	f3 := &file{"文件3"}
	f4 := &file{"文件4"}

	root := &folder{}
	dir1 := &folder{}

	root.add(f1)
	root.add(dir1)
	dir1.add(f2)
	dir1.add(f3)
	dir1.add(f4)

	root.execute()
}

type Composite interface {
	execute()
}

type folder struct {
	c []Composite
}

func (f *folder) execute() {
	for _, composite := range f.c {
		composite.execute()
	}
}

func (f *folder) add(composite Composite) {
	f.c = append(f.c, composite)
}

type file struct {
	name string
}

func (f *file) execute() {
	fmt.Println("file-name:", f.name)
}


