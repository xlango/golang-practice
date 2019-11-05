package parser

import (
	"practice/crawler/engine"
	"regexp"
)

const areaListUrl = `<a class="" href="(https://cd.julive.com/project/s/[a-z]+)"`

func ParserAreaList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(areaListUrl)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}

	//limit:=10
	for _, m := range matchs {
		result.Items = append(result.Items, "Area "+string(m[1])+"\n")
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParserCommunity,
			})

		//limit--
		//if limit==0 {
		//	break
		//}
	}

	return result
}
