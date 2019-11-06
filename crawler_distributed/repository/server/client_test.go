package main

import (
	"practice/crawler/engine"
	"practice/crawler/model"
	"practice/crawler_distributed/config"
	"practice/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	go serveRpc(host, "test1")
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(host)

	if err != nil {
		panic(err)
	}

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

	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)

	if err != nil {
		t.Errorf("result: %s; err: %v", result, err)
	}
}
