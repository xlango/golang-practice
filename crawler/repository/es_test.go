package repository

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"practice/crawler/engine"
	"practice/crawler/model"
	"testing"
)

func TestSave(t *testing.T) {
	doctor := model.Doctor{
		Name:       "test1",
		Zhicheng:   "test1",
		Hospital:   "test1",
		Department: "test1",
		Disease:    "test1",
		WebSite:    "test1",
		Tel:        "test1",
		Post:       "test1",
		Email:      "test1",
		Fax:        "test1",
		Address:    "test1",
	}

	item := engine.Item{
		Url:     "test",
		Type:    "test",
		Id:      "111111111",
		Payload: doctor,
	}

	err := save(item)

	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.10.222:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	result, err := client.Get().Index("idoctor").Type(item.Type).Id(item.Id).Do(context.Background())

	//var property  model.Property
	//err = json.Unmarshal(
	//	[]byte(result.Source), &property)
	if err != nil {
		panic(err)
	}

	t.Logf("%s", result.Source)
}
