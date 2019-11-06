package worker

import (
	"errors"
	"fmt"
	"log"
	"practice/crawler/engine"
	"practice/crawler/idoctor/parser"
	"practice/crawler_distributed/config"
)

type SerializeParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializeParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializeParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParserResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParserResult {
	result := engine.ParserResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		request, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializeing request: %v \n", err)
			continue
		}
		result.Requests = append(result.Requests,
			request)
	}
	return result
}

func deserializeParser(p SerializeParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParserKeshi:
		return engine.NewFuncParser(
			parser.ParserKeshiList,
			config.ParserKeshi), nil
	case config.ParserDoctor:
		return engine.NewFuncParser(
			parser.ParserDoctor,
			config.ParserDoctor), nil
	case config.ParserDoctorProfile:
		if d, ok := p.Args.(parser.DoctorParser); !ok {
			return parser.NewDoctorParser(d.UserName, d.Id, d.Url), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", p.Args)
		}
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}

}

type DoctorParser struct {
	UserName string
	Id       string
	Url      string
}
