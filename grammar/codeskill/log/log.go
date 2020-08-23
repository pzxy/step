package log

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"os"
	"step/grammar/codeskill/helper/conv"
	"step/grammar/codeskill/mygo"

	"time"
)

const (
	ThLevel          = 1
	Nc               = 0
	Unc              = 1
	logDir           = "logs"
	MaxLogFileSize   = 1024 * 1024 * 5
	FileSaveTime     = 240 * time.Hour
	FileIntervalTime = 480 * time.Hour
)

func InitGLog() bool {
	if !CreateDir(logDir) {
		fmt.Println("CreateDir() failed, logDir = ", logDir)
		return false
	}
	flag.Set("alsologtostderr", "true") // 日志写入文件的同时，输出到stderr
	flag.Set("log_dir", logDir)         // 日志文件保存目录
	flag.Set("v", conv.FormatInt(Unc))  // 配置V输出的等级。
	flag.Parse()
	glog.MaxSize = MaxLogFileSize

	mygo.MyGo(findAndRemoveOutTimeLogFile)

	return true
}

func DestructGLog() {
	glog.Flush()
}

// 可以打印error的调用栈
func ErrLog(err error) error {
	glog.Errorf("%+v", err)
	return err
}

func TErrLog(taskId string, err error) error {
	glog.Errorf("%v, %+v", taskId, err)
	return err
}

func findAndRemoveOutTimeLogFile() {
	for {
		select {
		case <-time.After(time.Duration(FileIntervalTime)):
			if err := removeOutTimeLogFile(); err != nil {
				glog.Errorf("clean outTime LogFile Failed err(%v)", err)
			}
		}
	}
}

func removeOutTimeLogFile() error {
	files, err := ioutil.ReadDir(logDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if time.Since(file.ModTime()).Seconds() > time.Duration(FileSaveTime).Seconds() {
			if err := os.Remove(logDir + "\\" + file.Name()); err != nil {
				return err
			}
		}
	}
	return nil
}
