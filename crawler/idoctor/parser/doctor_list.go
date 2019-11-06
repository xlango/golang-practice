package parser

import (
	"log"
	"practice/crawler/engine"
	"regexp"
)

const doctorUrl = `<span><a href="(/ys/([0-9]+).html)" target="_blank">([^<]+)</a></span>`

func ParserDoctor(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(doctorUrl)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matchs {
		name := string(m[3])
		//result.Items=append(result.Items,"Doctor "+name+"\n")
		log.Printf("Doctor : %v \n", name)
		result.Requests = append(
			result.Requests, engine.Request{
				Url:    "https://ysk.99.com.cn" + string(m[1]),
				Parser: NewDoctorParser(name, string(m[2]), "https://ysk.99.com.cn"+string(m[1])),
			})
	}

	return result
}
