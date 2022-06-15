package main

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"reflect"
	"time"
)

func main() {
	d := &Demo43{}
	Go(d.p, 10)
	time.Sleep(time.Second * 100)
}

func timer1() {

}

type Demo43 struct {
	A int
}

func (d *Demo43) p(c int) {
	for i := 0; i < c; i++ {
		fmt.Println(i)
	}
	time.Sleep(time.Second * 3)
	panic("aaaaaaa")
}
func Go(f interface{}, args ...interface{}) {
	fValue := reflect.ValueOf(f)
	if fValue.Kind() != reflect.Func {
		fmt.Println(fmt.Errorf("%v is not a func!", f))
		return
	}

	var argsValue []reflect.Value
	for i := range args {
		value := reflect.ValueOf(args[i])
		if value.Kind() == reflect.Invalid {
			fmt.Println(fmt.Errorf("%v Kind is Invalid", value))
			panic("Kind is Invalid")
			return
		}
		argsValue = append(argsValue, value)
	}

	go func() {
		// 另起一个协程，并调用该函数
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Recover Panic: ", err)
			}
		}()
		fValue.Call(argsValue)
	}()
}
func hash(seed []byte, size uint64) uint64 {
	b := sha1.Sum(seed)
	return binary.BigEndian.Uint64(b[:]) % size
}
