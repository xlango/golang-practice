package parser

import (
	"fmt"
	"practice/crawler/engine"
	"practice/crawler/model"
	"regexp"
)

var scoreLineRe = regexp.MustCompile(`<tr><td>([^<]+)</td><td>([^<]+)</td><td>([^<]+)</td><td>([^<]+)</td><td>([^<]+)</td><td>([^<]+)</td><td>([^<]+)</td></tr>`)

func ParserSchoolScore(contents []byte, name string) engine.ParserResult {

	scores := model.Score{
		School:     name,
		ScoreLines: make([]model.ScoreLine, 0),
	}

	scores.ScoreLines = extract(contents, scoreLineRe)

	fmt.Printf("Score info : %v \n", scores)
	result := engine.ParserResult{
		Items: []interface{}{scores},
	}

	return result
}

func extract(contents []byte, re *regexp.Regexp) []model.ScoreLine {
	scoreLines := make([]model.ScoreLine, 0)

	matchs := re.FindAllSubmatch(contents, -1)

	for _, m := range matchs {
		scoreLine := model.ScoreLine{
			Year:     string(m[1]),
			Top:      string(m[2]),
			Lower:    string(m[3]),
			Averrage: string(m[4]),
			Archive:  string(m[5]),
			Person:   string(m[6]),
			Pici:     string(m[7]),
		}
		scoreLines = append(scoreLines, scoreLine)
	}

	return scoreLines
}
