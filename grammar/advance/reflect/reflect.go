package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := "2021"
	v := reflect.TypeOf(&a)
	fmt.Println(v.Elem())
	vv := reflect.ValueOf(&a) //vv是中typ *rtype是 *string类型元数据
	vvv := vv.Elem()          //将*string类型元数据转为 string类型元数据。因为只有作用原变量身上才有意义。
	vvv.SetString("2022")
	fmt.Println(vvv.String())

}
