package main

import (
	"fmt"
	"sync"
)

//并发读写map syncMap 特性 并发效率还比较高

func main() {

	var scene sync.Map
	//将键值对保存到sync.Map中
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	//从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	//根据键删除对应的键值对
	scene.Delete("london")
	//遍历所有数据
	scene.Range(func(key, value interface{}) bool {
		fmt.Println("iterate:", key, value)
		return true
	})
}
