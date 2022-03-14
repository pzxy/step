package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

var (
	reg4FailedTestCase  *regexp.Regexp
	testCasesFailedFile *os.File
	testCoverageFile    *os.File
)

func init() {
	reg4FailedTestCase = regexp.MustCompile(`--- FAIL: [^(]+`)
}

func collectFailedResults(fileName string) (results [][]byte) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log(fmt.Sprintf("error at collectAndOutPutResults for (%v)", err.Error()))
		return
	}
	fmt.Println(fmt.Sprintf("result data:%s", data))
	return reg4FailedTestCase.FindSubmatch(data)
}

func saveAndOutputFailedResult(fileName string, item testItem) {
	results := collectFailedResults(fileName)
	for _, v := range results {
		if testCasesFailedFile != nil && len(v) > 0 {
			withReturn := fmt.Sprintf("%v\n", string(v))
			testCasesFailedFile.Write([]byte(withReturn))
		}
		log(fmt.Sprintf("test case failed: %v", item.testFuncNames))
		return
	}

	log(fmt.Sprintf("test case passed: %v", item.testFuncNames))
}
func saveCoverageResult(fileName *os.File, item testItem) {

	result := readCoverFileLine(fileName)

	for i := 1; i < len(result); i++ {
		withReturn := fmt.Sprintf("%v\n", result[i])
		testCoverageFile.Write([]byte(withReturn))
	}

	log(fmt.Sprintf("append coverage success: %v", item.testFuncNames))
}

func readCoverFileLine(file *os.File) (result []string) {
	// 之所以这里入参是 io.Reader类型能传入 *os.File类型,
	// 是由于*File实现了 io.Reader 这个接口
	br := bufio.NewReader(file) // 返回这个文件足够大小的reader对象的缓冲区类型
	for {
		if a, _, c := br.ReadLine(); c != io.EOF { // 返回行内容,若是太长的话,返回头的前缀,然后下次继续读
			result = append(result, string(a))
			continue
		}
		break
	}
	return
}

func genTestResultFile(fileName string) (*os.File, error) {
	testResultFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return testResultFile, err
	}
	return testResultFile, nil
}
func genTestCoverageFile(fileName string) (*os.File, error) {
	testCoverageFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return testCoverageFile, err
	}
	return testCoverageFile, nil
}
func genORCoverageFile(fileName string) (*os.File, error) {
	testCoverageFile, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return testCoverageFile, err
	}
	return testCoverageFile, nil
}

func initCoverageFile() {
	withReturn := fmt.Sprintf("%v\n", "mode: set")
	testCoverageFile.Write([]byte(withReturn))
}
