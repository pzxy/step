package main

import "fmt"

const (
	TEXT_EXT                    = ".txt"
	OUT                         = ".out"
	COVERAGE_DIR                = "coverage"
	COVERAGE_NAME               = "coverage" + OUT
	FILTER_FILE_NAME            = "filter" + TEXT_EXT
	FAILED_TEST_CASES_FILE_NAME = "failed_test_cases_outputs" + TEXT_EXT
	BUF_FILE_DIR                = "buf"
	COMMAND_SEARCH_DIR          = "path"
	COMMAND_SEARCH_DIR_USAGE    = "the directory to begin to test"
	SIZE_OF_TEST_CASES          = 1024 * 1024
)

func log(msg string) {
	fmt.Println(msg)
}
