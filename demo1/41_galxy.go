package main

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"strconv"
)

func main() {
	topArticles(3)
}

func topArticles(limit int32) []string {
	var totalData Articles
	for i := 0; i <= int(limit); i++ {
		resp, err := http.Get("https://jsonmock.hackerrank.com/api/articles?page=" + strconv.Itoa(i))
		if err != nil {
			break
		}
		var respData RespData
		buf, err := io.ReadAll(resp.Body)
		if err != nil {
			break
		}
		err = json.Unmarshal(buf, &respData)
		if err != nil {
			break
		}
		totalData = append(totalData, respData.Data...)
		if respData.TotalPages <= i {
			break
		}
	}
	var tmp Articles
	var ret []string
	for _, article := range totalData {
		if len(article.Title) != 0 {
			article.Name = article.Title
			tmp = append(tmp, article)
			continue
		}
		if len(article.StoryTitle) != 0 {
			article.Name = article.StoryTitle
			tmp = append(tmp, article)
			continue
		}
	}
	sort.Sort(tmp)
	for _, v := range tmp {
		ret = append(ret, v.Name)
	}
	//for _, v := range tmp {
	//	fmt.Println(fmt.Sprintf("%s----%d", v.Name, v.NumComments))
	//}
	return ret
}

type RespData struct {
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
	Total      int      `json:"total"`
	TotalPages int      `json:"total_pages"`
	Data       Articles `json:"data"`
}

type Article struct {
	Name        string
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

func (a Articles) Len() int      { return len(a) }
func (a Articles) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Articles) Less(i, j int) bool {
	if a[i].NumComments > a[j].NumComments {
		return true
	}
	if a[i].NumComments < a[j].NumComments {
		return false
	}
	return a[i].Name < a[j].Name
}
