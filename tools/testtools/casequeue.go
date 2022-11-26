package main

import "sync"

type testItem struct {
	dirPath       string
	fileName      string
	testFuncNames string
}

var (
	testCaseQueue     = make(chan []testItem, SIZE_OF_TEST_CASES)
	testCaseWaitGroup sync.WaitGroup
)

func RunTestCases() {
	for items := range testCaseQueue {
		for _, v := range items {
			runCommand(v)
		}
		testCaseWaitGroup.Done()
	}
}

func writeIntoTaskCasesQueue(result []testItem) {
	if len(result) > 0 {
		testCaseQueue <- result
	}
}
