package worker

import "practice/crawler/engine"

type CrawlService struct {
}

func (CrawlService) Process(
	request Request, result *ParseResult) error {

	engineReq, err := DeserializeRequest(request)
	if err != nil {
		return err
	}

	engineResult, err := engine.Work(engineReq)
	if err != nil {
		return err
	}

	*result = SerializeResult(engineResult)
	return nil
}
