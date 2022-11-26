package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	var (
		filter       []string
		err          error
		success      bool
		testFilePath string
	)

	clearBufferWorkspace()

	if filter, err = readIgnoreDirectoryList(); err != nil {
		panic(errors.Errorf("load filter.txt failed(%v)", err))
	}
	printIgnoreDirectoryList(filter)

	if testFilePath, success = getTestCasesSearchFromDirPath(); !success {
		panic(errors.Errorf("can`t find test directory"))
	}

	if testCasesFailedFile, err = genTestResultFile(FAILED_TEST_CASES_FILE_NAME); err != nil {
		panic(errors.Errorf("can`t create failed task cases result file"))
	}
	defer testCasesFailedFile.Close()

	if testCoverageFile, err = genTestCoverageFile(COVERAGE_NAME); err != nil {
		panic(errors.Errorf("can`t create coverage file"))
	}
	initCoverageFile()

	defer testCoverageFile.Close()

	log(fmt.Sprintf("starting search test cases in %v\n", testFilePath))

	if err = runTester(testFilePath, filter); err != nil {
		panic(errors.Errorf("runTester failed (%v)", err))
	}

	go RunTestCases()
	testCaseWaitGroup.Wait()
	log(fmt.Sprintf("all test cases over .... %v\n", testFilePath))
}
