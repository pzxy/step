package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	osEnv := os.Environ()
	var variables = make(map[string]string, len(osEnv))

	for _, env := range osEnv {
		// Can not use Split() on '=' since the value may have an '=' in it, so changed to use Index()
		index := strings.Index(env, "=")
		if index == -1 {
			continue
		}
		key := env[:index]
		value := env[index+1:]
		variables[key] = value
	}
	for k, v := range variables {
		fmt.Println(k, v)
	}
}
