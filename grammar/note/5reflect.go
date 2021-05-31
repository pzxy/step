package main

import (
	"fmt"
	"reflect"
)

type A43 interface {
	A433()
}
type a43 struct {
	Name string
}

type a42 struct {
	Name string
}

func (a *a43) A433() {

}

func main() {
	fmt.Println(ConfigurationInterfaceName)
}

var ConfigurationInterfaceName = typeInstanceToName((*A43)(nil))
var C1 = a42{}
var C2 = func(a string) string {
	return a
}

func typeInstanceToName(a interface{}) string {
	t := reflect.TypeOf(a)

	if name := t.Name(); name != "" {
		// non-interface types
		fmt.Println(t.PkgPath() + "." + name)
	}
	// interface types
	e := t.Elem()
	fmt.Println(e.PkgPath() + "." + e.Name())
	return "1"
}
