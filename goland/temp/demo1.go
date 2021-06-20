package temp

import (
	"fmt"
	"strings"
)

//1. 生成接口^i，或者 田 N

type CreateInterface struct { //omitempty 存在，如果name 为空，则反序列化出来的没有 Name
	Name string `json:"name,omitempty"`
	id   string `json:"id,omitempty"`
}

//2.  表达式类型
//表达式类型操作可以按Ctrl-Shift+P调用，如果您需要了解脱字符号处任何表达式的类型，它可以随时为您提供帮助。

func Goland1() {
	strings.Fields("")
}

//3. 使用 docker,创建后在 services 下面

//4. json 转换alt+enter，或者粘贴json 格式的数据。

//5. 快速修复,alt+enter,
func goland2() {
	hi := []byte{'1', '2'}
	m := map[string]string{
		"123": string(hi), //输入 hi 然后alt+enter，可以转换类型
	}
	fmt.Println(m)
}

//6. 根据已有方法，创建结构体。

type Page struct {
	Name string
}

func goland3() {
	_ = Page{ //使用 alt+enter 会生成 Page 结构体。
		Name: "1",
	}

}

//8. 控制可变参数
func goland4() {
	variadic("1", 1, 2, 3)
}
func variadic(a string, xs ...int) {

}
