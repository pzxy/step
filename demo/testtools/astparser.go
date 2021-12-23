package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
)

func genTestFuncNamesByAst(fileName string) {
	result := make([]testItem, 0)
	astNode, err := parser.ParseFile(token.NewFileSet(), fileName, nil, parser.ParseComments)
	if err != nil {
		log(err.Error())
		return
	}

	ast.Inspect(astNode,
		func(node ast.Node) bool {
			if ret, ok := node.(*ast.FuncDecl); ok {
				if strings.HasPrefix(ret.Name.String(), "Test") {
					result = append(result,
						testItem{
							dirPath:       filepath.Dir(fileName),
							fileName:      fileName,
							testFuncNames: ret.Name.String(),
						},
					)
				}
				return true
			}
			return true
		})

	writeIntoTaskCasesQueue(result)

	return
}
