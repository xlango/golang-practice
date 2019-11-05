package engine

import (
	"log"
	"practice/crawler/fetcher"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, err := work(r)
		if err != nil {
			continue
		}

		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("Got item %v \n", item)
		}

	}
}

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
