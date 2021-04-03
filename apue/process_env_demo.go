package main

import (
	"fmt"
	"os"
	"step/grammar/codeskill/log"
)

func main() {
	s := os.Getenv("GOPATH")
	fmt.Println(s)
	changeEnvNameDemo()
}

/**

todo

关于syscall和os包的思考，是不是syscall包底层是内核，而os包是用户呀？这里不懂，疑惑。
*/

/**
env
*/
func envDemo() {
	//命令行设置 直接 WWW = www.socrates.com 就行
	fmt.Printf("get env:%v \n", os.Getenv("WWW")) //区分大小写
	err := os.Setenv("WWW", "www.meetsocrates.com")
	if err != nil {
		log.ErrLog(err)
		return
	}
	fmt.Printf("after set env  :%v \n", os.Getenv("WWW")) //区分大小写

	err = os.Unsetenv("WWW")
	if err != nil {
		log.ErrLog(err)
		return
	}
	fmt.Printf("after delete env  :%v \n", os.Getenv("WWW")) //区分大小写
}

func changeEnvNameDemo() {
	oldEnv := "WWW"
	newEnv := "W4"
	err := os.Setenv(oldEnv, "www.meetsocrates.com")
	if err != nil {
		log.ErrLog(err)
		return
	}
	envValue := os.Getenv(oldEnv)
	if len(envValue) == 0 {
		fmt.Printf("env `WWW` does not exist \n")
		return
	}
	fmt.Printf("before env WWW change WWW：%v \n", envValue) //区分大小写

	err = os.Unsetenv(oldEnv)
	if err != nil {
		log.ErrLog(err)
		return
	}
	fmt.Printf("after env WWW delete WWW:%v \n", os.Getenv(oldEnv)) //区分大小写

	err = os.Setenv(newEnv, envValue)
	if err != nil {
		log.ErrLog(err)
		return
	}
	fmt.Printf("after env WWW change W4:%v \n", os.Getenv(newEnv)) //区分大小写

	err = os.Unsetenv(newEnv)
	if err != nil {
		log.ErrLog(err)
		return
	}
	fmt.Printf("after env WWW delete W4:%v \n", os.Getenv(newEnv)) //区分大小写

}
