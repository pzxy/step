package main

import "flag"

func getTestCasesSearchFromDirPath() (testDir string, ok bool) {
	pathFlag := flag.String(COMMAND_SEARCH_DIR, "", COMMAND_SEARCH_DIR_USAGE)
	flag.Parse()
	if *pathFlag != "" {
		return *pathFlag, true
	}

	flag.Usage()

	return "", false
}
