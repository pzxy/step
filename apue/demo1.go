package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//GetAllFile("/Users/pzxy/WorkSpace/Go/src/step")
	b, _ := json.Marshal("asdasdadad")
	os.Stdin.Read(b)
}

func GetAllFile(pathname string) {
	rd, _ := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("%v\n", fi.Name())
			GetAllFile(fmt.Sprintf(pathname+"/%v", fi.Name()))
		} else {
			//fmt.Println(fi.Name())
		}
	}
	return

}
