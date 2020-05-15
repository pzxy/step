package mygo

import (
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"reflect"
	"step/utils/helper/stack"
)

var (
	RcvPanicChan = make(chan struct{}, 1)
	isPanic      bool
)

// 用于封装需要go的代码（即需要另起协程执行的代码），在执行他们具体的函数之前加上一个抓取异常的操作
func MyGo(f interface{}, args ...interface{}) {
	fValue := reflect.ValueOf(f)
	if fValue.Kind() != reflect.Func {
		glog.Error(errors.Errorf("%v is not a func!", f))
		return
	}

	var argsValue []reflect.Value
	for i := range args {
		value := reflect.ValueOf(args[i])
		if value.Kind() == reflect.Invalid {
			glog.Error(errors.Errorf("%v Kind is Invalid", value))
			panic("Kind is Invalid")
			return
		}
		argsValue = append(argsValue, value)
	}

	go func() {
		// 另起一个协程，并调用该函数
		defer func() {
			if isPanic {
				RcvPanicChan <- struct{}{}
			}
		}()
		defer stack.PrintRecoverFromPanic(&isPanic)
		fValue.Call(argsValue)
	}()
}
