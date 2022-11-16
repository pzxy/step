package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

func main() {
	resp, err := http.Get("https://jsonmock.hackerrank.com/api/articles")
	if err != nil {
		return
	}
	var respData RespData
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(buf, &respData)
	if err != nil {
		return
	}

	sort.Sort(respData.Data)
	for _, v := range respData.Data {
		fmt.Println(fmt.Sprintf("NumComments:%d:, %+v", v.NumComments, v))
	}
}

type RespData struct {
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
	Total      int      `json:"total"`
	TotalPages int      `json:"total_pages"`
	Data       Articles `json:"data"`
}

type Article struct {
	Title       string      `json:"title"`
	Url         string      `json:"url"`
	Author      string      `json:"author"`
	NumComments int         `json:"num_comments"`
	StoryId     interface{} `json:"story_id"`
	StoryTitle  string      `json:"story_title"`
	StoryUrl    string      `json:"story_url"`
	ParentId    int         `json:"parent_id"`
	CreatedAt   int         `json:"created_at"`
}

type Articles []Article

func (a Articles) Len() int           { return len(a) }
func (a Articles) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Articles) Less(i, j int) bool { return a[i].NumComments > a[j].NumComments }
