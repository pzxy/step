package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json:"name`
	Age  int
}

func main() {
	test()
}
func test() {
	//1 反序列化 解析已知类型
	var jsonP = []byte(`[
		{"Name":"wonkung"}
	]`)
	p := make([]People, 0)
	pp := &p
	ppp := &pp
	pppp := &ppp
	err := json.Unmarshal(jsonP, &pppp)
	if err != nil {
		//	log.Errorf("json.Unmarshal :", err)
	}
	fmt.Printf("%v", p)

	fmt.Println()
	fmt.Println("******************")
}

type Student struct {
	Name string
	Age  int
}

func test2() {
	//var as []Student
	//as = append(as, Student{Name: "动物名字", Age: 12})
	//as = append(as, Student{Name: "动物名字", Age: 12})
	//as = append(as, Student{Name: "动物名字", Age: 12})
	//
	//jsonStr, _ := json.Marshal(as)
	//fmt.Printf("src : %s \n", jsonStr)

	//2 反序列化 解析未知类型
	//var f interface{}

	//b := []byte(`{"Name":"wonkung","Age":6,"Parents":["Yly","Whs"]}`)
	//path := strings.Replace(`D:\workspace\Go\src\step\basic\json\nodefile.json`, "\\", "/", -1)
	//b, _ := ioutil.ReadFile(path)
	////json.Unmarshal([]byte(jsonStr), &as)
	//fmt.Printf("f : %s \n", b)
	//
	//regex := `({+)`
	////math := `([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`
	//re := regexp.MustCompile(regex) //must含义 写的肯定是对的 不会出错 表达式加()表示提取
	////match := re.FindAllString(text,-1)//findstring 说明参数是string
	//matchs := re.FindAllStringSubmatch(url3, -1) //提取regex中()的数据
	//for _, m := range matchs {
	//	fmt.Println(m[1])
	//}
	//fmt.Printf("f : %v \n", as)
	//[NodeBookInfo:map[BookReason:0 Booked:false TaskId:] UniqueId:4309199455781888 NodeOccupyInfo:map[Db:map[] RobotSpaceOccupyInfo:map[TaskId: Occupied:false] TraySpaceOccupyInfo:map[Occupied:false TaskId:] CarSpaceOccupyInfo:map[Occupied:false TaskId:]]
	//类型断言
	//for k, v := range f.(map[string]interface{}) {
	//	switch vv := v.(type) {
	//	case string:
	//		fmt.Println(k, "is string", v)
	//	case float64:
	//		fmt.Println(k, "is float64", v)
	//	case []interface{}:
	//		fmt.Println(k, "is arrary :")
	//		for i, j := range vv {
	//			fmt.Println(i, j)
	//		}
	//	}
	//	//	fmt.Println(k,reflect.TypeOf(v))
	//}

}

//	fmt.Println(k,reflect.TypeOf(v))
