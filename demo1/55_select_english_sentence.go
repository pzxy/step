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
	var data []byte
	var n int
	args := readArgs()
	switch args[0] {
	case "s":
		data = readFile("/Users/pzxy/Workspace/english_sentence.txt")
	case "g":
		data = readFile("/Users/pzxy/Workspace/english_grammar.txt")
	default:
		help()
		return
	}
	if len(args) == 1 {
		n = 5
	} else {
		n = readInt(args[1])
	}
	en, zh := selectSentence(data, n)
	for i, v := range en {
		fmt.Printf("%d. %s\n", i, v)
	}
	fmt.Printf("\n\n\n\n\n")
	for i, v := range zh {
		fmt.Printf("%d. %s\n", i, v)
	}
}

func help() {
	fmt.Printf("s,sentence\ng,grammar\nh,help\nexample: ses s 10\n")
}
func readInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		help()
		panic(err)
	}
	return i
}

func readArgs() []string {
	args := os.Args
	if len(args) == 1 {
		return []string{"s", "5"}
	}
	if len(args) == 2 {
		_, err := strconv.Atoi(args[1])
		if err != nil {
			return []string{args[1]}
		}
		return []string{"s", args[1]}
	}
	args = args[1:]
	return args
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
			if len(strings.TrimSpace(k)) == 0 {
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
		s := strings.Split(k, "|")
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
