package main

import (
	"context"
	"fmt"
	"github.com/lampnick/doctron-client-go"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

// start server:
// git clone git@github.com:lampnick/doctron.git
// go run main.go --config=./conf/default.yaml
const pathBase = "/Users/pzxy/Downloads/coolshell/"
const domain = "http://127.0.0.1:8080"
const defaultUsername = "doctron"
const defaultPassword = "lampnick"

func main() {
	client := doctron.NewClient(context.Background(), domain, defaultUsername, defaultPassword)
	urls := getList()

	for i, url := range urls {
		fmt.Println("total: ", len(urls), "current: ", i+1, "url: ", url)
		req := doctron.NewDefaultHTML2PdfRequestDTO()
		req.ConvertURL = url
		response, err := client.HTML2Pdf(req)
		if err != nil {
			panic(err)
		}
		data2pdf(url, response.Data)
		success(url)
		time.Sleep(time.Second * 10)
	}
}

func data2pdf(url string, data []byte) {
	fileName := getFileName(url)
	fmt.Println(fmt.Sprintf("fileName: %s ,data: %d", fileName, len(data)))
	if len(fileName) == 0 {
		panic("fileName is empty")
	}
	if err := os.WriteFile(pathBase+fileName, data, 0666); err != nil {
		panic(err)
	}
}

func getList() []string {
	task := getContext(pathBase + "task.txt")
	success := getContext(pathBase + "success.txt")
	var ret []string
	if len(success) == 0 {
		return task
	}
	m := make(map[string]struct{}, 0)
	for _, v := range success {
		m[v] = struct{}{}
	}
	for _, v := range task {
		if _, ok := m[v]; !ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getContext(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(data), "\n")
	var ret []string
	for _, v := range s {
		v = strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}
		ret = append(ret, v)
	}
	return ret
}

func getFileName(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile("<title>(.*?)</title>")
	match := re.FindStringSubmatch(string(data))
	var name string
	if len(match) > 1 {
		name = match[1]
	}
	if len(name) != 0 {
		name = strings.Replace(name, "/", "-", -1)
	}
	return fmt.Sprintf("%s-%s.pdf", getNum(url), name)
}

func getNum(url string) string {
	re := regexp.MustCompile(`\d+`)
	result := re.FindAllString(url, -1)
	return result[0]
}

func success(s string) {
	file, err := os.OpenFile(pathBase+"success.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err := file.WriteString(s + "\n"); err != nil {
		panic(err)
	}
}
