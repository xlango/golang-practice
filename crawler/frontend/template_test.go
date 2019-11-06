package frontend

import (
	"html/template"
	"os"
	"practice/crawler/model"
	"testing"
)

func TestTemplate(t *testing.T) {
	template := template.Must(
		template.ParseFiles("template.html"))

	page := model.Doctor{}
	err := template.Execute(os.Stdout, page)
	if err != nil {
		panic(err)
	}
}
