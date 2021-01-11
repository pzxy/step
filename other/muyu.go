package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	// 原生字符串
	buf := `
    jQuery1124016963078166077672_1576074523699({"rc":0,"rt":17,"svr":2887135946,"lt":1,"full":0,"data":{"code":"399001","market":0,"name":"深证成指","decimal":2,"dktotal":7296,"klines":["2020-12-31,14226.28,14470.68,14476.55,14226.28,372217265,511265652736.00,1.76"]}});
    `

	//解释正则表达式
	reg := regexp.MustCompile(`\((.*?)\)`)
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}

	//提取关键信息
	result := reg.FindStringSubmatch(buf)

	var e interface{}
	e, _ = os.Open("ssss")
	_, ok := e.(io.ReadWriter)
	fmt.Println(ok)
	fmt.Println("====>", result[1])
}
