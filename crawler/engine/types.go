package engine

type ParserFunc func(contents []byte, url string) ParseResult

/**
用到的类型
*/

type Request struct { //请求
	Url        string
	ParserFunc ParserFunc
}

type ParseResult struct { //解析结果
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

func NilParser([]byte) ParseResult { //空方法
	return ParseResult{}
}
