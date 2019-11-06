package parser

import (
	"log"
	"practice/crawler/engine"
	"regexp"
)

const keshiListUrl = `<a href="(/department/[a-z]+/)" target="_self" class="">([^<]+)</a>`

func ParserKeshiList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(keshiListUrl)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}

	//limit:=10
	for _, m := range matchs {
		//result.Items=append(result.Items,"Keshi "+string(m[2])+"\n")
		log.Printf("Keshi : %v \n", string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url:    "https://ysk.99.com.cn" + string(m[1]),
				Parser: engine.NewFuncParser(ParserDoctor, "ParserDoctor"),
			})

		//limit--
		//if limit==0 {
		//	break
		//}
	}

	return result
}
