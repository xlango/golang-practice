package parser

import (
	"practice/crawler/engine"
	"regexp"
)

const communityUrl = `<a href="(https://cd.julive.com/project/[0-9]+.html)"`

func ParserCommunity(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(communityUrl)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matchs {
		name := string(m[1])
		result.Items = append(result.Items, "Community "+name+"\n")
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				ParserFunc: func(c []byte) engine.ParserResult {
					return ParserProperty(c, name)
				},
			})
	}

	return result
}
