package engine

import (
	"log"
	"practice/crawler/fetcher"
)

func Work(r Request) (ParserResult, error) {

	log.Printf("Fetching %s \n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+"fetching url %s: %v \n", r.Url, err)
		return ParserResult{}, err
	}

	parserResult := r.Parser.Parse(body)

	return parserResult, nil
}
