package parser

import (
	"fmt"
	"practice/crawler/engine"
	"practice/crawler/model"
	"regexp"
)

var unitPriceRe = regexp.MustCompile(`<span><em>(.*)元/m²起</em></span>`)
var totalPriceRe = regexp.MustCompile(`<span>(.*)万元</span>`)

//var typeRe  = regexp.MustCompile(`<a href="https://cd.julive.com/project/[0-9]+/details.html" data-sa-point="{&quot;event&quot;:&quot;e_click_project_type&quot;,&quot;properties&quot;:{&quot;fromModule&quot;:&quot;m_basic_info&quot;,&quot;fromItem&quot;:&quot;i_project_type&quot;,&quot;fromItemIndex&quot;:&quot;-1&quot;,&quot;toPage&quot;:&quot;p_project_details&quot;,&quot;toModule&quot;:&quot;&quot;}}">
//												([^<]+)</a>`)
//var addressRe  = regexp.MustCompile(`<a href="javascript:void[^;]+;" data-sa-point="{&quot;event&quot;:&quot;e_click_surrounding_analysis_entry&quot;,&quot;properties&quot;:{&quot;fromModule&quot;:&quot;m_basic_info&quot;,&quot;fromItem&quot;:&quot;i_surrounding_analysis_entry&quot;,&quot;fromItemIndex&quot;:&quot;-1&quot;,&quot;toPage&quot;:&quot;p_project_home&quot;,&quot;toModule&quot;:&quot;&quot;}}">
//												([^<]+)</a>`)
var timeRe = regexp.MustCompile(`<p class="txt-address"><span>([^<]+)</span></p>`)

func ParserProperty(contents []byte, name string) engine.ParserResult {

	property := model.Property{
		Name: name,
	}

	property.UnitPrice = extract(contents, unitPriceRe)
	//property.TotalPriceUp,_=strconv.Atoi(strings.Split(extract(contents,totalPriceRe), "-")[0])
	//property.TotalPriceDown,_=strconv.Atoi(strings.Split(extract(contents,totalPriceRe), "-")[1])
	//property.Address=extract(contents,addressRe)
	property.OpenTime = extract(contents, timeRe)

	fmt.Printf("Property info : %v \n", property)
	result := engine.ParserResult{
		Items: []interface{}{property},
	}

	return result
}

func extract(contents []byte, re *regexp.Regexp) string {

	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	}

	return ""

}
