package main

import (
	"fmt"
	"sync"
)

//sync.Map的基本用法
//loadOrStore不能更改value
//使用方法:
//创建对象:var mapInt = new(sync.Map) 或者 mapInt := sync.Map{}
//CRUD:
//增加:mapInt.Store(1,"波")
//删除:	mapInt.Delete(2) 其中key不存在不会报错
//查询:v,ok := mapInt.Load(2) 返回value和ok,未查到ok为false
//修改:mapInt.Store(1,"波") 若是key已经存在,则会修改
//查询增加:	v,ok := mapInt.LoadOrStore(2,"囧")
//若是不存在返回false,并且将值加进去,存在则返回true和value,不能更改值

//粗浅原理:
//用了两个map两存储数据read(只读 原子操作) dirty(读写,需要加锁) 存储数据的键值对是:map[interface{}]*entry,值是*entry指针
//entry三种值
//nil: entry已被删除了，并且m.dirty为nil
//expunged: entry已被删除了，并且m.dirty不为nil，而且这个entry不存在于m.dirty中
//其它： entry是一个正常的值

//无论是增加修改,删除,查询,都会从read开始的,然后再到dirty
//比如增加的话是从read开始的,read只能读,只是从read中读取这个值,看看有没有这个值,然后再加锁对dirty操作,给dirty增加了值
//以后,查的时候从read开始查,查不到就给Map里面的misses字段增加一,
// 当misses>=len(m.dirty)的长度时候,就把dirty提升为read,又开始从read中查询,read中是原子操作不需要加锁.
//删除的话read中标记为nil就可以了,dirty中是直接删除的
//注意的几点:
//LoadOrStore方法如果提供的key存在，则返回已存在的值(Load)，否则保存提供的键值(Store)
//Range方法调用前会做一个m.dirty的提升
//
//sync.Map优化点:
//空间换时间。 通过冗余的两个数据结构(read、dirty),实现加锁对性能的影响。
//使用只读数据(read)，避免读写冲突。
//动态调整，miss次数多了之后，将dirty数据提升为read。
//double-checking。
//延迟删除。 删除一个键值只是打标记，只有在提升dirty的时候才清理删除的数据。
//优先从read读取、更新、删除，因为对read的读取不需要锁。

//完整的来回顾这个流程:
//增加或者修改:我们执行增加(Store)操作,首先会检验read中是否存在这个值
// 存在的话而且没有被标记为nil则直接存储这这个值中;此时因为m.dirty也指向这个entry,dirty也保存的这个值,
// 若果不存在则给dirty增加这个值
// 或者被标记为nil的话,改变标记,更新值

func main() {
	//var mapInt = new(sync.Map)
	mapInt := sync.Map{}
	//add element
	mapInt.Store(1, "波")
	mapInt.Store(2, "多")
	mapInt.Store(3, "野")
	fmt.Println("before delete key:")

	//iterator
	mapInt.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	//delete
	mapInt.Delete(4)
	fmt.Println("after delete key:")
	//iterator
	mapInt.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)

		return true
	})
	//query

	v, ok := mapInt.Load(2)
	if ok {
		fmt.Println(v)
	}
	//query or add
	v, ok = mapInt.LoadOrStore(2, "囧")
	fmt.Println("loadOrStore :", v, ok)
	//iterator
	fmt.Println("after query or add key:")
	mapInt.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	//add
	mapInt.Store(1, "囧")
	//iterator
	fmt.Println("after add key:")
	mapInt.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
