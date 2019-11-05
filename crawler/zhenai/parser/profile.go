package parser

import (
	"fmt"
	"practice/crawler/engine"
	"practice/crawler/model"
	"regexp"
)

var nikeName = regexp.MustCompile(`<h1 data-v-5b109fc3="" class="nickName">([^<]+)</h1>`)
var marriageRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`)

func ParserProfile(contents []byte, name string) engine.ParserResult {

	profile := model.Profile{}

	profile.Name = name
	profile.Marriage = extractString(contents, marriageRe)

	fmt.Printf("User info : %v", profile)
	result := engine.ParserResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	}

	return ""
}
