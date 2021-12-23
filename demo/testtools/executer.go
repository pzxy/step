package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
)

func execute(outPut io.Writer, command string, params ...string) error {
	//fmt.Println(command, params)
	cmd := exec.Command(command, params...)
	cmd.Stdout = outPut
	cmd.Stderr = outPut

	return cmd.Run()
}

func getWindowsDiskOfPath(path string) string {
	return filepath.VolumeName(path)
}

func runByOsType(file *os.File, item testItem) {
	switch runtime.GOOS {
	case "darwin":
		if e := execute(file, "bash", "-c", fmt.Sprintf("cd %v && go test -v %v -test.run %v$", item.dirPath, item.fileName, item.testFuncNames)); e != nil {
			log(e.Error())
		}
	case "windows":
		disk := getWindowsDiskOfPath(item.dirPath)

		if e := execute(file, "cmd", "/c", fmt.Sprintf("%v && cd %v && go test -v -timeout=600s  -coverprofile=coverage.out -coverpkg=pss/task/...,pss/dijkstra/...,pss/pss/...,pss/minddleware/...  -test.run %v$",
			disk, item.dirPath, item.testFuncNames)); e != nil {
			log(e.Error())
		}
	case "linux":
		if e := execute(file, "bash", "-c", fmt.Sprintf("cd %v && go test -v %v -test.run %v$", item.dirPath, item.fileName, item.testFuncNames)); e != nil {
			log(e.Error())
		}
	}
}

func runCommand(item testItem) {
	singleTestCaseBufFile := path.Join(BUF_FILE_DIR, fmt.Sprintf("%v%v", item.testFuncNames, TEXT_EXT))

	singleCoverageFile := path.Join(COVERAGE_DIR, fmt.Sprintf("%v%v", item.testFuncNames, OUT))

	file, err := genTestResultFile(singleTestCaseBufFile)

	coverDstFile, err := genTestResultFile(singleCoverageFile)

	if err != nil {
		log(fmt.Sprintf("genTestResultFile faid for %v", err.Error()))
		return
	}

	runByOsType(file, item)
	//需要将覆盖率文件 item.pathDir 下的 cover.out 内容移动到 coverFile 中
	coverFile := path.Join(item.dirPath, fmt.Sprintf("%v", COVERAGE_NAME))

	//coverSrcFile 是当前方法所在包下的coverage.out的文件对象
	//coverDstFile 是coverage目录下的文件对象
	coverCopySrcFile, err := genORCoverageFile(coverFile)
	coverAppendSrcFile, err := genORCoverageFile(coverFile)
	if err != nil {
		log(fmt.Sprintf("open file faid err=%v\n", err.Error()))
	}
	copyFile(coverDstFile, coverCopySrcFile)

	//将覆盖率文件追加到coverage的同级目录下的 coverage.out 文件后面, 并且要去掉追加的第一行,这个文件对象是 testCoverageFile,现在已经有了
	saveCoverageResult(coverAppendSrcFile, item)

	file.Close()
	coverCopySrcFile.Close()
	coverAppendSrcFile.Close()
	coverDstFile.Close()
	os.Truncate(coverFile, 0)
	saveAndOutputFailedResult(singleTestCaseBufFile, item)
}

//func copyFile(dstFile *os.File, srcFile *os.File) {
//
//	reader := bufio.NewReader(srcFile)
//
//	writer := bufio.NewWriter(dstFile)
//	io.Copy(writer, reader)
//}
func copyFile(dstFile *os.File, srcFile *os.File) {
	reader := bufio.NewReader(srcFile)
	for {
		if a, _, c := reader.ReadLine(); c != io.EOF {
			writeStr := fmt.Sprintf("%v\n", string(a))
			dstFile.Write([]byte(writeStr))
			continue
		}
		break
	}
}
