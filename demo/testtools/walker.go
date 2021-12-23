package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func handleTestFile(path string, info os.FileInfo) bool {
	if !strings.HasSuffix(path, "_test.go") {
		return false
	}

	testCaseWaitGroup.Add(1)
	go genTestFuncNamesByAst(path)

	return false
}

func handleTestDirAndFile(filter []string, path string, info os.FileInfo) (err error) {
	if info.IsDir() ||
		filterContains(filter, path) {
		return nil
	}

	if !handleTestFile(path, info) {
		return nil
	}

	return nil
}

func handleTester(filter []string, path string, info os.FileInfo, err error) error {
	if err != nil {
		log(fmt.Sprintf("error in handleTestFile(%v)", err))
		return nil
	}

	return handleTestDirAndFile(filter, path, info)
}

func runTester(testFilePath string, filter []string) error {
	return filepath.Walk(testFilePath,
		func(path string, info os.FileInfo, err error) error {
			return handleTester(filter, path, info, err)
		})
}
