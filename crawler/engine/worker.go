package engine

import (
	"log"
	"practice/crawler/fetcher"
)

func work(r Request) (ParserResult, error) {

	log.Printf("Fetching %s \n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+"fetching url %s: %v \n", r.Url, err)
		return ParserResult{}, err
	}

	parserResult := r.ParserFunc(body)

	return parserResult, nil
}
