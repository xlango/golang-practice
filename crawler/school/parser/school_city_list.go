package parser

import (
	"practice/crawler/engine"
	"regexp"
)

const schoolCityListUrl = `<a href="(/chaxun/.+.html)" target="_blank">([^<]+)</a>`

func ParserSchoolCityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(schoolCityListUrl)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}

	//limit:=10
	for _, m := range matchs {
		result.Items = append(result.Items, "City "+string(m[2])+"\n")
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        "http://www.daxuecn.com" + string(m[1]),
				ParserFunc: ParserSchool,
			})

		//limit--
		//if limit==0 {
		//	break
		//}
	}

	return result
}
