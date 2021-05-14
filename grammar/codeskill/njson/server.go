package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"net/http"
)

type Student struct {
	Name string
}

type httpServer struct{}

func (h *httpServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	if len(body) == 0 {
		return
	}
	str := string(body)
	var students []Student
	gjson.ForEachLine(str, func(line gjson.Result) bool {
		var a Student
		b := []byte(line.String())
		json.Unmarshal(b, &a)
		students = append(students, a)
		return true
	})

	for i, s := range students {
		fmt.Println("receive data ====>", i, s.Name)
	}
	fmt.Println()
	fmt.Fprintf(w, "ok ")
}

func main() {
	h := &httpServer{}
	fmt.Println("run server at port 9090")
	http.ListenAndServe("0.0.0.0:9090", h)
}

func ReadBody(reader io.Reader) ([]Student, error) {
	var abc []Student
	json.NewDecoder(reader).Buffered()
	err := json.NewDecoder(reader).Decode(&abc)
	if err != nil {
		return abc, err
	}
	return abc, nil
}
