package engine

/**
用到的类型
*/
type Request struct { //请求
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct { //解析结果
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult { //空方法
	return ParseResult{}
}
