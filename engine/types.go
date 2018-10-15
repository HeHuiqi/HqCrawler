package engine

type ParserFunc func(contents []byte,url string) ParserResult

type Request struct {
	Url string
	ParserFunc ParserFunc
}

type ParserResult struct {
	Requests []Request
	//Items []interface{}
	Items []Item

}

type Item struct {

	Url string
	Id string
	Type string
	Payload interface{}
}

func NilParser([]byte) ParserResult  {

	return ParserResult{}
}