package parser

import (
	"fmt"
	"practice/crawler/engine"
	"practice/crawler/model"
	"regexp"
)

var zhichengRe = regexp.MustCompile(`<i>([^<]+)</i><a href="/ys/baocuo/[0-9]+.html" target="_blank">信息纠错</a>`)
var hospitalRe = regexp.MustCompile(`<p class="info-rtxt1">[^>]+>([^<]+)</a>--[^<]+</p>`)
var departmentRe = regexp.MustCompile(`<p class="info-rtxt1">[^>]+>[^<]+</a>--([^<]+)</p>`)
var diseaseRe = regexp.MustCompile(`<p class="info-rtxt2">[^>]+>([^<]+)</a></p>`)
var websiteRe = regexp.MustCompile(`<a href="http://www.[a-z0-9]+.com" target="_blank" rel="nofollow">(http://www.[a-z0-9]+.com)</a>`)
var telRe = regexp.MustCompile(`<li><span class="dot-ico2 dot-txt1">电&nbsp话：</span><em>([^<]*)</em></li>`)
var postRe = regexp.MustCompile(`<li><span class="dot-ico3 dot-txt1">传&nbsp真：</span><em>([^<]*)</em></li>`)
var emailRe = regexp.MustCompile(`<li><span class="dot-ico4">电子邮件：</span><em>([^<]*)</em></li>`)
var faxRe = regexp.MustCompile(`<li><span class="dot-ico5 dot-txt1">邮&nbsp编：</span><em>([^<]*)</em></li>`)
var addrRe = regexp.MustCompile(`<li><span class="dot-ico6 dot-txt1">地&nbsp址：</span><em>([^<]*)</em></li>`)

func ParserDoctorContent(contents []byte, name string, id string, url string) engine.ParserResult {

	doctor := model.Doctor{
		Name: name,
	}

	doctor.Zhicheng = extract(contents, zhichengRe)
	doctor.Hospital = extract(contents, hospitalRe)
	doctor.Department = extract(contents, departmentRe)
	doctor.Disease = extract(contents, diseaseRe)
	doctor.WebSite = extract(contents, websiteRe)
	doctor.Tel = extract(contents, telRe)
	doctor.Post = extract(contents, postRe)
	doctor.Email = extract(contents, emailRe)
	doctor.Fax = extract(contents, faxRe)
	doctor.Address = extract(contents, addrRe)

	item := engine.Item{
		Id:      id,
		Type:    "doctorInfo",
		Payload: doctor,
		Url:     url,
	}

	fmt.Printf("Doctor info : %v \n", doctor)
	result := engine.ParserResult{
		Items: []engine.Item{item},
	}

	return result
}

func extract(contents []byte, re *regexp.Regexp) string {

	matchs := re.FindSubmatch(contents)
	if len(matchs) >= 2 {
		return string(matchs[1])
	}

	return ""
}
