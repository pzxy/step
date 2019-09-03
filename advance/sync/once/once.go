package once

import (
	"fmt"
	"sync"
	"time"
)

//func main() {
//	for i:=0;i<10;i++{
//		 Test2()
//	}
//	time.Sleep(time.Second*3)
//}
func Test3(){
	for i:=0;i<10;i++{
		Test2()
	}
	time.Sleep(time.Second*3)
}
func test1(){
	var one sync.Once
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func(i int) {
			one.Do(onetest)
			wg.Done()
		}(i)
		go func(i int) {
			one.Do(onetest)
			wg.Done()
		}(i)
		fmt.Println(i)
	}
	wg.Wait()
}
func onetest() {
	fmt.Println("执行一次")
}
var one sync.Once
func Test2(){
	for i := 0; i < 10; i++ {
	one.Do(onetest)
	fmt.Println(i)
	}
}
