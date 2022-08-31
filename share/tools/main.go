package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/lunny/log"
)

var (
	failReg *regexp.Regexp
	passReg *regexp.Regexp
	logDir  string
)

type item struct {
	dir      string
	fileName string
	funcName string
}

func init() {
	failReg = regexp.MustCompile(`--- FAIL: [^(]+`)
	passReg = regexp.MustCompile(`--- PASS: [^(]+`)
	year, month, day := time.Now().Date()
	logDir = fmt.Sprintf("%v-%v-%v", year, month, day)
}

func main() {
	testDir := flag.String("path", "./", "the directory to begin to test")
	flag.Parse()
	items := make([]item, 0)

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Error("make buffer directory failed for (%s) ", err.Error())
		return
	}

	pass, err := os.OpenFile(path.Join(logDir, "pass.txt"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer pass.Close()
	fail, err := os.OpenFile(path.Join(logDir, "fail.txt"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer fail.Close()

	filepath.Walk(*testDir, func(fileName string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Error(err.Error())
			return nil
		}
		if !strings.HasSuffix(fileName, "_test.go") {
			return nil
		}
		astNode, err := parser.ParseFile(token.NewFileSet(), fileName, nil, parser.ParseComments)
		if err != nil {
			log.Error(err.Error())
			return nil
		}
		ast.Inspect(astNode,
			func(node ast.Node) bool {
				ret, ok := node.(*ast.FuncDecl)
				if !ok {
					return true
				}
				if !strings.HasPrefix(ret.Name.String(), "Test") {
					return true
				}
				items = append(items,
					item{
						dir:      filepath.Dir(fileName),
						fileName: fileName,
						funcName: ret.Name.String(),
					},
				)
				return true
			})
		return nil
	})
	for _, item := range items {
		testLogFile := path.Join(logDir, item.funcName+".txt")
		doItem(testLogFile, item)
		collectResult(testLogFile, pass, fail)
	}
}

func doItem(testLogFile string, item item) (err error) {
	f, err := os.OpenFile(testLogFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	defer f.Close()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return run(f, item)
}

func run(file *os.File, item item) error {
	switch runtime.GOOS {
	case "darwin":
		if err := execute(file, "bash", "-c",
			fmt.Sprintf("cd %v && go test -v  %v -run ^%v$ ", item.dir, item.dir, item.funcName)); err != nil {
			log.Error(err.Error())
		}
	case "windows":
		if err := execute(file, "cmd", "/c",
			fmt.Sprintf("%v && cd %v && go test -v -run ^%v$", filepath.VolumeName(item.dir), item.dir, item.funcName)); err != nil {
			log.Error(err.Error())
		}
	case "linux":
		if err := execute(file, "bash", "-c",
			fmt.Sprintf("cd %v && go test -v %v  -run ^%v$ ", item.dir, item.dir, item.funcName)); err != nil {
			log.Error(err.Error())
		}
	}
	return nil
}

func execute(outPut io.Writer, command string, params ...string) error {
	log.Infof("== %s %s", command, params)
	cmd := exec.Command(command, params...)
	cmd.Stdout = outPut
	cmd.Stderr = outPut
	return cmd.Run()
}

func collectResult(in string, pass *os.File, fail *os.File) {
	data, err := ioutil.ReadFile(in)
	if err != nil {
		log.Error(err.Error())
		return
	}

	for _, v := range passReg.FindSubmatch(data) {
		if len(v) <= 0 {
			continue
		}
		pass.Write(v)
		pass.Write([]byte("\n"))
		return
	}

	for _, v := range failReg.FindSubmatch(data) {
		if len(v) <= 0 {
			continue
		}
		fail.Write(v)
		fail.Write([]byte("\n"))
	}

}
