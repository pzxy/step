package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func readLine(file *os.File) (result []string) {
	br := bufio.NewReader(file)
	for {
		if a, _, c := br.ReadLine(); c != io.EOF {
			result = append(result, string(a))
			continue
		}
		break
	}

	return
}

func readIgnoreDirectoryList() (result []string, err error) {
	var file *os.File
	if file, err = os.Open(FILTER_FILE_NAME); err != nil {
		return nil, err
	}

	defer file.Close()
	result = append(result, readLine(file)...)

	return
}

func printIgnoreDirectoryList(filter []string) {
	log("tester will ignore test files in these directory :")
	for _, v := range filter {
		log(v)
	}
	log("-------------------------------------------------------------")
}

func filterContains(filter []string, path string) bool {
	for _, v := range filter {
		if strings.HasPrefix(path, v) {
			return true
		}
	}
	return false
}
