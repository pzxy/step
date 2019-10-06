package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func main() {
	var fileName = `C:/Users/pzxy/Desktop/test/test2/dt_starttask_test.go`
	testParseFile(fileName)
}

//解析文件
func testParseFile(fileName string) {
	//parser.ParseComments

	astNode, err := parser.ParseFile(token.NewFileSet(), fileName, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	ast.Inspect(astNode, func(node ast.Node) bool {
		if ret, ok := node.(*ast.FuncDecl); ok {
			if strings.HasPrefix(ret.Name.String(), "Test") {
				fmt.Printf("fileName:%v \ntestFuncNames:%v\n", fileName, ret.Name.String())
			}
		}

		return true
	})

}
