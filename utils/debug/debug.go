package debug

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"step/utils/log"
	"step/utils/myfile"
	"sync"
	"time"
)

var (
	open = true
)

type obj struct {
	io      *os.File
	lock    sync.Mutex
	fileIdx int
	tName   Type
}

func Open() bool {
	return open
}

func Init(o bool) {
	open = o
	if !o {
		return
	}
	InitDebugList()
}

func Close() {
	if !open {
		return
	}

	types.Range(
		func(key, value interface{}) bool {
			ins := value.(*obj)
			if ins.io != nil {
				ins.io.Close()
			}
			return true
		},
	)

}

func (o *obj) getDebugFileName() string {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	pid := os.Getpid()
	return fmt.Sprintf("%v[%v].%04d-%02d-%02d-%02d-%02d.%d.log" /*, proc*/, o.tName, o.fileIdx, year, month, day, hour, minute, pid)
}

func (o *obj) openFile() (err error) {
	fileName := o.getDebugFileName()
	filePath := filepath.Join("debug", fileName)
	if err = myfile.CreateOrAppendFile(filePath); err != nil {
		panic(err)
	}
	o.io, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	return err
}

func (o *obj) write(msg string) {
	o.lock.Lock()
	defer o.lock.Unlock()

	o.lazyInit()

	logInfo := fmt.Sprintf("%v - %v \n", time.Now().Format("15:04:05.000000"), msg)

	if _, err := o.io.Write([]byte(logInfo)); err != nil {
		log.ErrLog(errors.Errorf("write debug(%v) log failed(%v)", logInfo, err))
		return
	}

	file, _ := o.io.Stat()
	if file.Size() > 1024*1024*5 {
		o.fileIdx++
		o.io.Close()
		o.openFile()
	}
}

func (o *obj) lazyInit() {
	if o.io == nil {
		if err := o.openFile(); err != nil {
			panic(err)
		}
	}
}
