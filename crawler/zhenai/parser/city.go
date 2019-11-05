package parser

import (
	"practice/crawler/engine"
	"regexp"
)

//const cityUrl  = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
const cityUrl = `(<a href="/chaxun/[0-9]+.html)" target="_blank">([^<]+)</a>`

func ParserCity(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityUrl)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matchs {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name+"\n")
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				ParserFunc: func(c []byte) engine.ParserResult {
					return ParserProfile(c, name)
				},
			})
	}

	return result
}
