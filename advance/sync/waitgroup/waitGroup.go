package waitgroup

import (
	"fmt"
	"sync"
)

//同时读写
func writeData2(intChan chan int, wg *sync.WaitGroup) {
	defer close(intChan)
	defer wg.Done()

	for i := 1; i <= 5000; i++ {
		intChan <- i
	}
}
func readData2(intChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读出来的数据: %v\n", v)
	}
}

func main() {
	intChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	//必须传地址
	go readData2(intChan, &wg)
	go writeData2(intChan, &wg)
	wg.Wait()
}
