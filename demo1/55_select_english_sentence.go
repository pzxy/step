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
	fmt.Println("the total number is ", len(sentences))
	if len(sentences) == 0 {
		panic("the content is empty")
	}
	en := make([]string, n)
	zh := make([]string, n)
	var tmp string
	for i := 0; i < n; i++ {
		tmp = sentences[rand.Intn(len(sentences)-1)]
		s := strings.Split(tmp, ";")
		if len(s) == 0 {
			en[i] = ""
			zh[i] = ""
		} else if len(s) == 1 {
			en[i] = s[0]
			zh[i] = ""
		} else {
			en[i] = s[0]
			zh[i] = s[1]
		}
	}
	return en, zh
}
