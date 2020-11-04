package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	A("/Users/pzxy/WorkSpace/Go/src/step/grammar/advance/verify/example.pb.go")
}

func A(fileName string) {
	fs := token.NewFileSet()
	astNode, err := parser.ParseFile(fs, fileName, nil, parser.ParseComments)
	if err != nil {
		return
	}
	ast.Inspect(astNode,
		func(node ast.Node) bool {
			if ret, ok := node.(*ast.GenDecl); ok {
				for _, v := range ret.Specs {
					recStruct(v, "")
				}
			}
			return true
		})
}

func recStruct(v interface{}, head string) {

	t, isSpec := v.(*ast.TypeSpec)
	if !isSpec {
		return
	}
	st, isStruct := t.Type.(*ast.StructType)
	if !isStruct {
		return
	}
	fmt.Println(head + t.Name.Name)
	head += "-"

	for _, field := range st.Fields.List {
		starExpr, isStarExpr := field.Type.(*ast.StarExpr)
		if !isStarExpr {
			continue
		}
		ident, isIdent := starExpr.X.(*ast.Ident)
		if !isIdent {
			continue
		}
		recStruct(ident.Obj.Decl, head)
	}
}
