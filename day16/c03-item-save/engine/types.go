package engine

// 请求
type Request struct {
	Url			string
	ParseFunc	func([]byte) ParseResult
}

// 解析结果
type ParseResult struct {
	Requests 	[]Request
	Items		[]interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

// 有效数据结构封装
type Item struct {
	Type	string
	Index	string
	Data	interface{}
}