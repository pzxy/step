package main

import (
	"fmt"
	"regexp"
)

func main() {
	test()
}

func test() {
	imgUrl := "https://bbs.csdn.net/topics/3601078.bin"
	re := regexp.MustCompile(`.*/([^/]+)$`)
	matchs := re.FindAllStringSubmatch(imgUrl, -1)

	var name string
	for _, m := range matchs {
		name = m[1]
	}
	fmt.Println(name)
	//path := "D://"+name
	//f , err := os.Create(path)
	//if err != nil {
	//	panic(err)
	//}
	//
	//res, err := http.Get(imgUrl)
	//if err != nil {
	//	return
	//}
	//
	//defer res.Body.Close()
	//
	//written, err := io.Copy(f, res.Body)
	//
	//body,err:= ioutil.ReadFile(path)
	//
	//fmt.Printf("Total length: %d", written)
	//fmt.Printf("Total length: %s", string(body))
}
