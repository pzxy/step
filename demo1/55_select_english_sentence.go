package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// sudo go build -o /usr/local/go/bin/ses 55_select_english_sentence.go
func main() {
	n := readArgs()
	data := readFile("/Users/pzxy/Workspace/english_sentence.txt")
	en, zh := selectSentence(data, n)
	for i, v := range en {
		fmt.Printf("%d. %s\n", i, v)
	}
	fmt.Printf("\n\n\n\n\n")
	for i, v := range zh {
		fmt.Printf("%d. %s\n", i, v)
	}
}

func readArgs() int {
	args := os.Args
	if len(args) == 1 {
		return 5
	}
	i, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("args has to a number")
		panic(err)
	}
	return i
}

func readFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return data
}

func selectSentence(data []byte, n int) ([]string, []string) {
	sentences := strings.Split(string(data), "\n")
	fmt.Println("the entry is ", len(sentences))
	if len(sentences) == 0 {
		panic("the content is empty")
	}
	// remove duplicate elements
	m := make(map[string]struct{})
	if len(sentences) > n {
		var limit int
		for len(m) < n && limit < 1000 {
			k := sentences[rand.Intn(len(sentences)-1)]
			if _, ok := m[k]; ok {
				limit++
				continue
			}
			m[k] = struct{}{}
		}
	} else {
		for _, k := range sentences {
			m[k] = struct{}{}
		}
	}
	// split en and zh
	en := make([]string, 0)
	zh := make([]string, 0)
	for k, _ := range m {
		s := strings.Split(k, ";")
		if len(s) == 0 {
			en = append(en, "")
			zh = append(zh, "")
		} else if len(s) == 1 {
			en = append(en, s[0])
			zh = append(zh, "")
		} else {
			en = append(en, s[0])
			zh = append(zh, s[1])
		}
	}
	return en, zh
}
