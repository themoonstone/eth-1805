package engine

// 请求
type Request struct {
	Url		string
	// 内容解析函数(解析器)
	ParseFunc func([]byte) ParseResult
}

// 解析结果
type ParseResult struct {
	Requests 	[]Request
	Items		[]interface{}
}

// 添加一个空的解析器
func NilParser([]byte) ParseResult  {
	return ParseResult{}
}