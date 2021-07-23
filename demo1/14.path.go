package main

import (
	"fmt"
	"path"
	"strings"
)

func main() {
	url := "https://asdadasdasd"
	cmd := "bash -c 'cd /tmp \\\n " +
		"&& wget " + url + " \\\n " +
		"&& tar -zxvf edgex_v1.6.0_Linux_x86_64.tar.gz  '"
	fmt.Println(cmd)

	fmt.Println(strings.Contains("armv7", ""))

}

func testPath() {
	a := "/tmp/"
	aa := "/tmp"
	aaa := "/tmp//"
	b := "//tare.tar"
	s := "https://github.com/pzxy/tuya-edge-driver-sdk-go/releases/download/"
	s2 := "v123"
	s3 := "nihaoya.tar"
	fmt.Println(path.Join(a, b))
	fmt.Println(path.Join(aa, b))
	fmt.Println(path.Join(aaa, b))
	fmt.Println(path.Join(s, s2, s3))
}
