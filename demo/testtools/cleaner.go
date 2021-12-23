package main

import (
	"fmt"
	"os"
)

func clearBufferWorkspace() {
	if st, err := os.Stat(FAILED_TEST_CASES_FILE_NAME); err == nil {
		if err := os.Remove(st.Name()); err != nil {
			log(fmt.Sprintf("delete failed test case results failed for (%v) ", err))
		}
	}

	if st, err := os.Stat(COVERAGE_NAME); err == nil {
		if err := os.Remove(st.Name()); err != nil {
			log(fmt.Sprintf("delete coverage.out failed for (%v) ", err))
		}
	}

	if st, err := os.Stat(BUF_FILE_DIR); err == nil && st.IsDir() {
		if err := os.RemoveAll(st.Name()); err != nil {
			log(fmt.Sprintf("delete buffer directroy failed for (%v) ", err))
		}
	}

	if st, err := os.Stat(COVERAGE_DIR); err == nil && st.IsDir() {
		if err := os.RemoveAll(st.Name()); err != nil {
			log(fmt.Sprintf("delete coverage failed for (%v) ", err))
		}
	}

	if err := os.MkdirAll(BUF_FILE_DIR, os.ModePerm); err != nil {
		log(fmt.Sprintf("make buffer directory failed for (%v) ", err))
	}

	if err := os.MkdirAll(COVERAGE_DIR, os.ModePerm); err == nil {
		log(fmt.Sprintf("make coverage directory failed for (%v) ", err))
	}
}
