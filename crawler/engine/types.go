package engine

type Parser interface {
	Parse(contents []byte) ParserResult
	Serialize() (name string, args interface{})
}

type ParserFunc func(contents []byte) ParserResult

type Request struct {
	Url    string
	Parser Parser
}

//type SerializeParser struct {
//	Name string
//	Args interface{}
//}

type ParserResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type NilParser struct {
}

func (NilParser) Parse(contents []byte) ParserResult {
	return ParserResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	paser ParserFunc
	name  string
}

func (f *FuncParser) Parse(contents []byte) ParserResult {
	return f.paser(contents)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		paser: p,
		name:  name,
	}
}
