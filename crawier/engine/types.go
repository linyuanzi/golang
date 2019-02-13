package engine

type ParseFunc func(contents []byte, url string) ParseResult

type Request struct {
	Url string
	ParserFunc ParseFunc
}

type ParseResult struct {
	Requests []Request
	Items []Item
}

type Item struct {
	Url 	string
	Id 		string
	Type 	string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}