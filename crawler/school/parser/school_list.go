package parser

import (
	"practice/crawler/engine"
	"regexp"
)

//const cityUrl  = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
const schoolUrl = `<a href="(/chaxun/[0-9]+.html)" target="_blank">([^<]+)</a>`

func ParserSchool(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(schoolUrl)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matchs {
		name := string(m[2])
		result.Items = append(result.Items, "School "+name+"\n")
		result.Requests = append(
			result.Requests, engine.Request{
				Url: "http://www.daxuecn.com" + string(m[1]),
				ParserFunc: func(c []byte) engine.ParserResult {
					return ParserSchoolScore(c, name)
				},
			})
	}

	return result
}
