package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

//在复杂的json数据中取出自己想要的数据,并且去重
//想读读源码,看了好久,勉强看懂一部分,写的很牛呀,把我都难住了 : )
type people struct {
	//其实就是把自己想要的数据的那部分的结构写出来就行了.
	ViewModel struct {
		AccessoryData struct {
			Crm_sup_base_SUP_ATTACHMENT []struct {
				AttachPath string
			}
		}
	}
}

type SupChange struct {
	Time   string
	Reason string
	Change string
}
type Params struct {
	Flag int
	SupChange SupChange
}
type People2 struct {
	Params Params
}


func main() {
	path := strings.Replace(`D:\workspace\Go\src\step\basic\json\re2.json`, "\\", "/", -1)
	//data, err := ioutil.ReadFile("/home/pzxy/go/src/step/basic/json/re.json")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if error := recover(); error != nil {
			fmt.Printf("have one error : %v", error)
		}
	}()
	var peo []People2
	err = json.Unmarshal(data, &peo)
	if err != nil {
		panic(err)
	}
	//for _,v := range peo{
	//	type Params struct {
	//		Flag int
	//		SupChange SupChange
	//	}
	//}

	fmt.Printf("%v\n", peo)
	//array := peo.ViewModel.AccessoryData.Crm_sup_base_SUP_ATTACHMENT
	//var smap sync.Map
	//for _, v := range array {
	//	smap.Store(v.AttachPath, "attachPath")
	//}
	//smap.Range(func(key, value interface{}) bool {
	//	fmt.Printf("%v : %v\n", value, key)
	//	return true
	//})

}
